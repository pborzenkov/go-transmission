package transmission

import (
	"context"
)

type blocklistResponse struct {
	Size int `json:"blocklist-size"`
}

// UpdateBlocklist updates peer blocklist and returns the size of newly obtained list.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L578
func (c *Client) UpdateBlocklist(ctx context.Context) (int, error) {
	var resp blocklistResponse

	if err := c.callRPC(ctx, "blocklist-update", nil, &resp); err != nil {
		return 0, err
	}

	return resp.Size, nil
}
