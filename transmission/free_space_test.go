package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetFreeSpace(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"free-space","arguments":{"path":"/tmp"}}`)

		fmt.Fprintf(w, `{"result":"success","arguments":{"path":"/tmp","size-bytes":1024}}`)
	})

	size, err := client.GetFreeSpace(context.Background(), "/tmp")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if want, got := int64(1024), size; want != got {
		t.Errorf("unexpected free space, want = %d, got = %d", want, got)
	}
}
