package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCloseSession(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"session-close"}`)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	if err := client.CloseSession(context.Background()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
