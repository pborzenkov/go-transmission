package transmission

import (
	"context"
)

// GetTorrents returns information requested by fields for the torrents identified by ids.
func (c *Client) GetTorrents(ctx context.Context, ids Identifier, fields ...TorrentField) ([]*Torrent, error) {
	if len(fields) == 0 {
		fields = allTorrentFields
	}

	var getTorrentsReq = struct {
		IDs    Identifier     `json:"ids,omitempty"`
		Fields []TorrentField `json:"fields"`
	}{ids, fields}

	var resp = struct {
		Torrents []*torrentJSON `json:"torrents"`
	}{}
	if err := c.callRPC(ctx, "torrent-get", getTorrentsReq, &resp); err != nil {
		return nil, err
	}

	uc := c.getUnitConversion()
	torrents := make([]*Torrent, 0, len(resp.Torrents))
	for _, tj := range resp.Torrents {
		t, err := tj.torrent(uc)
		if err != nil {
			return nil, err
		}
		torrents = append(torrents, t)
	}

	return torrents, nil
}

// GetRecentlyRemovedTorrentIDs returns a slice of torrent IDs that's been
// removed in the past hour.
func (c *Client) GetRecentlyRemovedTorrentIDs(ctx context.Context) ([]ID, error) {
	var getTorrentsReq = struct {
		IDs    Identifier     `json:"ids"`
		Fields []TorrentField `json:"fields"`
	}{RecentlyActive(), []TorrentField{TorrentFieldID}}

	var resp = struct {
		Removed []ID `json:"removed"`
	}{}
	if err := c.callRPC(ctx, "torrent-get", getTorrentsReq, &resp); err != nil {
		return nil, err
	}

	return resp.Removed, nil
}
