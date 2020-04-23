package transmission

import (
	"context"
)

func (c *Client) queueMove(ctx context.Context, cmd string, ids Identifier) error {
	type queueMoveReq struct {
		IDs Identifier `json:"ids,omitempty"`
	}

	return c.callRPC(ctx, cmd, &queueMoveReq{IDs: ids}, nil)
}

// QueueMoveToTop tells Transmission to move torrents identified by ids to the
// top of the queue.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L601
func (c *Client) QueueMoveToTop(ctx context.Context, ids Identifier) error {
	return c.queueMove(ctx, "queue-move-top", ids)
}

// QueueMoveToBottom tells Transmission to move torrents identified by ids to
// the bottom of the queue.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L601
func (c *Client) QueueMoveToBottom(ctx context.Context, ids Identifier) error {
	return c.queueMove(ctx, "queue-move-bottom", ids)
}

// QueueMoveUp tells transmission to move torrents identified by ids up in the
// queue.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L601
func (c *Client) QueueMoveUp(ctx context.Context, ids Identifier) error {
	return c.queueMove(ctx, "queue-move-up", ids)
}

// QueueMoveDown tells transmission to move torrents identified by ids down in
// the queue.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L601
func (c *Client) QueueMoveDown(ctx context.Context, ids Identifier) error {
	return c.queueMove(ctx, "queue-move-down", ids)
}
