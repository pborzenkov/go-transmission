package transmission

import (
	"context"
)

func (c *Client) torrentAction(ctx context.Context, cmd string, ids Identifier) error {
	type torrentActionReq struct {
		IDs Identifier `json:"ids"`
	}

	return c.callRPC(ctx, cmd, &torrentActionReq{IDs: ids}, nil)
}

// StartTorrents starts torrents identified by ids. If Transmission already has
// enough torrents in its download or upload queues, the torrents are added to
// the end of the queue instead.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L86
func (c *Client) StartTorrents(ctx context.Context, ids Identifier) error {
	return c.torrentAction(ctx, "torrent-start", ids)
}

// StartTorrentsNow forcibly starts torrents identified by ids even of download
// or upload queues are full.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L86
func (c *Client) StartTorrentsNow(ctx context.Context, ids Identifier) error {
	return c.torrentAction(ctx, "torrent-start-now", ids)
}

// StopTorrents stops torrents identified by ids.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L86
func (c *Client) StopTorrents(ctx context.Context, ids Identifier) error {
	return c.torrentAction(ctx, "torrent-stop", ids)
}

// VerifyTorrents instructs Transmission to verify torrents identified by ids.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L86
func (c *Client) VerifyTorrents(ctx context.Context, ids Identifier) error {
	return c.torrentAction(ctx, "torrent-verify", ids)
}

// ReannounceTorrents tells Transmission to reannounce (ask tracker for more
// peers) torrents identified by ids.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L86
func (c *Client) ReannounceTorrents(ctx context.Context, ids Identifier) error {
	return c.torrentAction(ctx, "torrent-reannounce", ids)
}

// RenameTorrentPath renames a file or directory in a torrent.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L438
func (c *Client) RenameTorrentPath(ctx context.Context, id SingularIdentifier, path, name string) error {
	var renameTorrentPathReq = struct {
		IDs  SingularIdentifier `json:"ids"`
		Path string             `json:"path"`
		Name string             `json:"name"`
	}{id, path, name}

	return c.callRPC(ctx, "torrent-rename-path", &renameTorrentPathReq, nil)
}
