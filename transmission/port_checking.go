package transmission

import (
	"context"
)

type portTestResponse struct {
	Open bool `json:"port-is-open"`
}

// IsPortOpen reports if Transmission incoming port is accessible from the
// outside world.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L584
func (c *Client) IsPortOpen(ctx context.Context) (bool, error) {
	var resp portTestResponse

	if err := c.callRPC(ctx, "port-test", nil, &resp); err != nil {
		return false, err
	}

	return resp.Open, nil
}
