package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestQueueMove(t *testing.T) {
	var tests = []struct {
		name string
		fn   func(ctx context.Context, c *Client, ids Identifier) error
		ids  Identifier
		body string
	}{
		{
			name: "to_top",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.QueueMoveToTop(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"queue-move-top","arguments":{"ids":1}}`,
		},
		{
			name: "to_bottom",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.QueueMoveToBottom(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"queue-move-bottom","arguments":{"ids":1}}`,
		},
		{
			name: "up",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.QueueMoveUp(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"queue-move-up","arguments":{"ids":1}}`,
		},
		{
			name: "down",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.QueueMoveDown(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"queue-move-down","arguments":{"ids":1}}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			client, handle, teardown := setup(t)
			defer teardown()

			handle(func(w http.ResponseWriter, r *http.Request) {
				testBody(t, r, tc.body)

				fmt.Fprintf(w, `{"result":"success"}`)
			})

			if err := tc.fn(context.Background(), client, tc.ids); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
