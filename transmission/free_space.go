package transmission

import (
	"context"
)

type freeSpaceRequest struct {
	Path string `json:"path"`
}

type freeSpaceResponse struct {
	Path      string `json:"path"`
	SizeBytes int64  `json:"size-bytes"`
}

// GetFreeSpace returns how much space in bytes is available in the specified folder.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L618
func (c *Client) GetFreeSpace(ctx context.Context, path string) (int64, error) {
	var resp freeSpaceResponse

	if err := c.callRPC(ctx, "free-space", &freeSpaceRequest{Path: path}, &resp); err != nil {
		return 0, err
	}

	return resp.SizeBytes, nil
}
