package transmission

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setup(t *testing.T, opts ...Option) (client *Client, handle func(func(http.ResponseWriter, *http.Request)), teardown func()) { //nolint:lll
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client, err := New(server.URL, append(opts, WithHTTPClient(server.Client()))...)
	if err != nil {
		t.Fatalf("failed to initialize Client: %v", err)
	}

	return client, func(cb func(http.ResponseWriter, *http.Request)) {
		mux.HandleFunc(defaultRPCPath, cb)
	}, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if got := r.Method; want != got {
		t.Errorf("unexpected HTTP method, want = %q, got = %q", want, got)
	}
}

func testHeader(t *testing.T, r *http.Request, header, want string) {
	t.Helper()

	if got := r.Header.Get(header); want != got {
		t.Errorf("unexpected HTTP headed %q, want = %q, got = %q", header, want, got)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()

	gotBody := new(bytes.Buffer)
	wantBody := new(bytes.Buffer)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("failed to read request body: %v", err)
	}
	if err := json.Indent(gotBody, body, "", "  "); err != nil {
		t.Fatalf("failed to indent JSON body: %v", err)
	}
	if err := json.Indent(wantBody, []byte(want), "", "  "); err != nil {
		t.Fatalf("failed to indent JSON test body: %v", err)
	}
	if want, got := strings.TrimSpace(wantBody.String()), strings.TrimSpace(gotBody.String()); !cmp.Equal(want, got) {
		t.Errorf("unexpected request body, diff = \n%s", cmp.Diff(want, got))
	}
}

func parseTestURL(t *testing.T, str string) *url.URL {
	t.Helper()

	u, err := url.Parse(str)
	if err != nil {
		t.Fatalf("failed to parse URL %q: %v", str, err)
	}
	return u
}

func TestCallRPC(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"method":"test","arguments":{"arg":"testarg"}}`)

		fmt.Fprintf(w, `{"result":"success","arguments":{"resp":"testresp"}}`)
	})

	type testRequest struct {
		Arg string `json:"arg"`
	}
	type testResponse struct {
		Resp string `json:"resp"`
	}

	var gotResponse testResponse
	err := client.callRPC(context.Background(), "test", &testRequest{Arg: "testarg"}, &gotResponse)
	if err != nil {
		t.Fatalf("failed to execute RPC call: %v", err)
	}

	if want, got := (testResponse{Resp: "testresp"}), gotResponse; !cmp.Equal(want, got) {
		t.Errorf("unexpected response, diff = \n%s", cmp.Diff(want, got))
	}
}

func TestCallRPC_noArgs(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, `{"method":"test"}`)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	if err := client.callRPC(context.Background(), "test", nil, nil); err != nil {
		t.Errorf("failed to execute RPC call: %v", err)
	}
}

func TestCallRPC_auth(t *testing.T) {
	client, handle, teardown := setup(t, WithAuth("admin", "qwerty"))
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			t.Errorf("no auth header in the request")
		}
		if want, got := "admin", user; want != got {
			t.Errorf("unexpected username, want = %q, got = %q", want, got)
		}
		if want, got := "qwerty", pass; want != got {
			t.Errorf("unexpected password, want = %q, got = %q", want, got)
		}

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	if err := client.callRPC(context.Background(), "test", nil, nil); err != nil {
		t.Errorf("failed to execute RPC call: %v", err)
	}
}

func TestCallRPC_userAgent(t *testing.T) {
	client, handle, teardown := setup(t, WithUserAgent("go-rtorrent"))
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, "User-Agent", "go-rtorrent")

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	if err := client.callRPC(context.Background(), "test", nil, nil); err != nil {
		t.Errorf("failed to execute RPC call: %v", err)
	}
}

func TestCallRPC_badHTTPCode(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.callRPC(context.Background(), "test", nil, nil); err == nil {
		t.Errorf("expected RPC call to fail")
	}
}

func TestCallRPC_csrf(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	token := "some-magic-token"
	var reqNum int
	handle(func(w http.ResponseWriter, r *http.Request) {
		reqNum++

		testMethod(t, r, "POST")
		testBody(t, r, `{"method":"test","arguments":1}`)
		if reqNum == 1 {
			w.Header().Add(headerCSRF, token)
			w.WriteHeader(http.StatusConflict)
			return
		}
		testHeader(t, r, headerCSRF, token)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	if err := client.callRPC(context.Background(), "test", 1, nil); err != nil {
		t.Errorf("failed to execute RPC call: %v", err)
	}
}

func TestCallRPC_failed(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, `{"result":"some serius failure"}`)
	})

	if err := client.callRPC(context.Background(), "test", nil, nil); err == nil {
		t.Errorf("expected RPC call to fail")
	}
}
