package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestUpdateBlocklist(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"blocklist-update"}`)

		fmt.Fprintf(w, `{"result":"success","arguments":{"blocklist-size":42}}`)
	})

	size, err := client.UpdateBlocklist(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if want, got := 42, size; want != got {
		t.Errorf("unexpected blocklist size, want = %d, got = %d", want, got)
	}
}
