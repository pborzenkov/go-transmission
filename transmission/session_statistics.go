package transmission

import (
	"context"
	"time"
)

// SessionStats holds session statistics
type SessionStats struct {
	// Total number of torrents
	Torrents int `json:"torrentCount"`
	// Number of active torrents
	ActiveTorrents int `json:"activeTorrentCount"`
	// Number of paused torrents
	PausedTorrents int `json:"pausedTorrentCount"`

	// Cumulative download rate (bytes/s)
	DownloadRate int64 `json:"downloadSpeed"`
	// Cumulative upload rate (bytes/s)
	UploadRate int64 `json:"uploadSpeed"`

	// Statistics about current session
	CurrentSession Stats `json:"current-stats"`
	// Statistics about all sessesions (including current)
	AllSessions Stats `json:"cumulative-stats"`
}

// Stats contains statistics about current sessions or cumulative statistics
// about all sessions.
type Stats struct {
	// Total amount of downloaded data (bytes)
	Downloaded int64 `json:"downloadedBytes"`
	// Total amount of uploaded data (bytes)
	Uploaded int64 `json:"uploadedBytes"`
	// Total number of added files
	Files int `json:"filesAdded"`
	// Number of sessions (always 1 for current session).
	Sessions int `json:"sessionCount"`
	// The amount of time the session has been active.
	ActiveFor time.Duration `json:"secondsActive"`
}

// GetSessionStats returns statistics about current session and cumulative
// statistics about all sessions.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L546
func (c *Client) GetSessionStats(ctx context.Context) (*SessionStats, error) {
	resp := new(SessionStats)

	if err := c.callRPC(ctx, "session-stats", nil, resp); err != nil {
		return nil, err
	}
	resp.CurrentSession.ActiveFor *= time.Second
	resp.AllSessions.ActiveFor *= time.Second

	return resp, nil
}
