package transmission

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

// AddTorrentReq holds information needed to add new torrent to transission.
// Either URL or Meta must be set.
type AddTorrentReq struct {
	// Either a path/URL to torrent or magnet link
	URL *string `json:"filename,omitempty"`
	// Contents of the torrent file
	Meta io.Reader `json:"-"`
	// Custom download directory for the torrent
	DownloadDirectory *string `json:"download-dir,omitempty"`
	// Cookies to attach to HTTP request when downloading torrent file over
	// the network
	Cookies []Cookie `json:"-"`

	// Don't automatically start torrent
	Paused *bool `json:"paused,omitempty"`

	// Bandwidth priority
	Priority *Priority `json:"bandwidthPriority,omitempty"`
	// List of high priority file indicies
	HighPriorityFiles []int `json:"priority-high,omitempty"`
	// List of normal priority file indicies
	NormalPriorityFiles []int `json:"priority-normal,omitempty"`
	// List of low priority file indicies
	LowPriorityFiles []int `json:"priority-low,omitempty"`

	// Custom peer limit
	PeerLimit *int `json:"peer-limit,omitempty"`

	// List of file indicies to download
	WantedFiles []int `json:"files-wanted,omitempty"`
	// List of files indicies to not download
	UnwatedFiles []int `json:"files-unwanted,omitempty"`
}

// NewTorrent describes newly added torrent.
type NewTorrent struct {
	ID   ID     `json:"id"`
	Hash Hash   `json:"hashString"`
	Name string `json:"name"`
}

// AddTorrent adds new torrent to Transmission.
func (c *Client) AddTorrent(ctx context.Context, req *AddTorrentReq) (*NewTorrent, error) {
	if req.URL != nil && req.Meta != nil {
		return nil, errors.New("transmission: can't have both URL and Meta set")
	}

	var addTorrentJSON = struct {
		*AddTorrentReq
		Meta    *string `json:"metainfo,omitempty"`
		Cookies *string `json:"cookies,omitempty"`
	}{
		AddTorrentReq: req,
	}
	if req.Meta != nil {
		meta := new(strings.Builder)
		w := base64.NewEncoder(base64.StdEncoding, meta)
		if _, err := io.Copy(w, req.Meta); err != nil {
			return nil, err
		}
		if err := w.Close(); err != nil {
			return nil, err
		}
		addTorrentJSON.Meta = OptString(meta.String())
	}
	if len(req.Cookies) > 0 {
		c := make([]string, len(req.Cookies))
		for i := range req.Cookies {
			c[i] = req.Cookies[i].String()
		}
		addTorrentJSON.Cookies = OptString(strings.Join(c, "; "))
	}

	var addTorrentResp = struct {
		Added     *NewTorrent `json:"torrent-added"`
		Duplicate *NewTorrent `json:"torrent-duplicate"`
	}{}
	if err := c.callRPC(ctx, "torrent-add", addTorrentJSON, &addTorrentResp); err != nil {
		return nil, err
	}

	t := addTorrentResp.Added
	if t == nil {
		t = addTorrentResp.Duplicate
	}

	return t, nil
}
