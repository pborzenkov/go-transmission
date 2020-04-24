package transmission

import (
	"context"
	"net/url"
	"time"
)

// TrackerReplacement contains a replacement announce URL for a tracker.
type TrackerReplacement struct {
	ID          int
	AnnounceURL *url.URL
}

// SetTorrentReq holds input for SetTorrent request.
type SetTorrentReq struct {
	// Torrent download rate limit (bytes/s)
	DownloadRateLimit *int64 `json:"-"`
	// Honor torrent download rate limit
	DownloadRateLimitEnabled *bool `json:"downloadLimited,omitempty"`
	// Torrent upload rate limit (bytes/s)
	UploadRateLimit *int64 `json:"-"`
	// Honor torrent upload rate limit
	UploadRateLimitEnabled *bool `json:"uploadLimited,omitempty"`
	// Whether to honor session download/upload limits or not
	HonorSessionLimits *bool `json:"honorsSessionLimits,omitempty"`

	// Torrent priority
	Priority *Priority `json:"bandwidthPriority,omitempty"`
	// An array of high priority file indicies (empty array means all
	// files)
	HighPriorityFiles []int `json:"priority-high"`
	// An array of normal priority file indicies (empty array means all
	// files)
	NormalPriorityFiles []int `json:"priority-normal"`
	// An array of low priority file indicies (empty array means all files)
	LowPriorityFiles []int `json:"priority-low"`
	// Position of this torrent in the queue
	PositionInQueue *int `json:"queuePosition,omitempty"`

	// Indicies of files to download (empty array means all files)
	WantedFiles []int `json:"files-wanted"`
	// Indicies of files to not download (empty array means all files)
	UnwantedFiles []int `json:"files-unwanted"`

	// Maximum number of peers
	PeerLimit *int `json:"peer-limit,omitempty"`

	// New location of the torrent contents
	Location *string `json:"location,omitempty"`

	// Stop torrent after given time of inactivity
	IdleSeedingLimit *time.Duration `json:"-"`
	// Which IdleSeedingLimit value to use
	IdleSeedingLimitMode *Limit `json:"seedIdleMode,omitempty"`
	// Stop seeding after reaching the given ratio
	UploadRatioLimit *float64 `json:"seedRatioLimit,omitempty"`
	// Which UploadRatioLimit value to use
	UploadRatioLimitMode *Limit `json:"seedRatioMode,omitempty"`

	// Add trackers with the given URIs to the torrent
	TrackersToAdd []*url.URL `json:"-"`
	// Remove trackers with the given ids from the torrent
	TrackerToRemove []int `json:"trackerRemove,omitempty"`
	// List of trackers with updated announcement URIs
	TrackersToReplace []TrackerReplacement `json:"-"`
}

// SetTorrents modifies parameters for the torrents identified by ids.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L105
func (c *Client) SetTorrents(ctx context.Context, ids Identifier, req *SetTorrentReq) error {
	uc := c.getUnitConversion()

	var setTorrentsJSON = struct {
		*SetTorrentReq
		IDs               Identifier     `json:"ids,omitempty"`
		DownloadRateLimit *int64         `json:"downloadLimit,omitempty"`
		UploadRateLimit   *int64         `json:"uploadLimit,omitempty"`
		IdleSeedingLimit  *time.Duration `json:"seedIdleLimit,omitempty"`
		TrackersToAdd     []string       `json:"trackerAdd,omitempty"`
		TrackersToReplace []interface{}  `json:"trackerReplace,omitempty"`
	}{
		SetTorrentReq: req,
		IDs:           ids,
	}
	if req.DownloadRateLimit != nil {
		setTorrentsJSON.DownloadRateLimit = OptInt64(*req.DownloadRateLimit / uc.speed)
	}
	if req.UploadRateLimit != nil {
		setTorrentsJSON.UploadRateLimit = OptInt64(*req.UploadRateLimit / uc.speed)
	}
	if req.IdleSeedingLimit != nil {
		setTorrentsJSON.IdleSeedingLimit = OptDuration(*req.IdleSeedingLimit / time.Minute)
	}
	if len(req.TrackersToAdd) > 0 {
		setTorrentsJSON.TrackersToAdd = make([]string, len(req.TrackersToAdd))
		for i, t := range req.TrackersToAdd {
			setTorrentsJSON.TrackersToAdd[i] = t.String()
		}
	}
	if len(req.TrackersToReplace) > 0 {
		setTorrentsJSON.TrackersToReplace = make([]interface{}, len(req.TrackersToReplace)*2)
		for i, t := range req.TrackersToReplace {
			setTorrentsJSON.TrackersToReplace[i*2] = t.ID
			setTorrentsJSON.TrackersToReplace[i*2+1] = t.AnnounceURL.String()
		}
	}

	return c.callRPC(ctx, "torrent-set", &setTorrentsJSON, nil)
}
