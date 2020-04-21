package transmission

import (
	"context"
)

// CloseSession terminates Transmission session.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L593
func (c *Client) CloseSession(ctx context.Context) error {
	return c.callRPC(ctx, "session-close", nil, nil)
}
