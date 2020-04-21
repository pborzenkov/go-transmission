package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestIsPortOpen(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"port-test"}`)

		fmt.Fprintf(w, `{"result":"success","arguments":{"port-is-open":true}}`)
	})

	open, err := client.IsPortOpen(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if want, got := true, open; want != got {
		t.Errorf("got unexpected result, want = %v, got = %v", want, got)
	}
}
