package transmission

import (
	"context"
	"time"
)

// SetSessionReq holds modifications to be applied to current session. Only
// non-nil values are taken into account.
type SetSessionReq struct {
	// Maximum allowed download rate in "turtle" mode (bytes/s)
	TurtleDownloadRateLimit *int64 `json:"-"`
	// Maximum allowed upload rate in "turtle" mode (bytes/s)
	TurtleUploadRateLimit *int64 `json:"-"`
	// Indicates whether "turtle" mode is enabled right now
	TurtleEnabled *bool `json:"alt-speed-enabled,omitempty"`
	// Indicates whether "turtle" mode is controlled by the configured
	// schedule
	TurtleScheduleEnabled *bool `json:"alt-speed-time-enabled,omitempty"`
	// Indicates on what days of the week to turn on "turtle" mode
	TurtleScheduleOnDays *Weekday `json:"alt-speed-time-day,omitempty"`
	// When to turn on "turtle" mode (in minutes after midnight)
	TurtleScheduleStartsAt *int `json:"alt-speed-time-begin,omitempty"`
	// When to turn off "turtle" mode (in minutes after midnight)
	TurtleScheduleStopsAt *int `json:"alt-speed-time-end,omitempty"`

	// Maximum allowed download rate (bytes/s)
	DownloadRateLimit *int64 `json:"-"`
	// Indicates whether download rate limit is enabled
	DownloadRateLimitEnabled *bool `json:"speed-limit-down-enabled,omitempty"`
	// Maximum allowed upload rate (bytes/s)
	UploadRateLimit *int64 `json:"-"`
	// Indicates whether upload rate limit is enabled
	UploadRateLimitEnabled *bool `json:"speed-limit-up-enabled,omitempty"`

	// Location of the peer blocklist
	BlocklistURL *string `json:"blocklist-url,omitempty"`
	// Indicates whether or not peer blocklist is enabled
	BlocklistEnabled *bool `json:"blocklist-enabled,omitempty"`

	// Maximum size of disk cache in bytes
	CacheSize *int64 `json:"-"`

	// Default path to download torrents
	DownloadDirectory *string `json:"download-dir,omitempty"`
	// Path for incomplete torrents (if enabled)
	IncompleteDirectory *string `json:"incomplete-dir,omitempty"`
	// Indicates whether to keep torrents in incomplete directory until done
	IncompleteDirectoryEnabled *bool `json:"incomplete-dir-enabled,omitempty"`
	// Indicates whether Transmission will append '.part' suffix to
	// incomplete files
	RenameIncompleteFiles *bool `json:"rename-partial-files,omitempty"`

	// Max number of torrents to download at once
	DownloadQueueLimit *int `json:"download-queue-size,omitempty"`
	// Indicates whether or not download queue limit is enabled
	DownloadQueueLimitEnabled *bool `json:"download-queue-enabled,omitempty"`
	// Max number of torrents to seed at once
	UploadQueueLimit *int `json:"seed-queue-size,omitempty"`
	// Indicates whether or not upload queue limit is enabled
	UploadQueueLimitEnabled *bool `json:"seed-queue-enabled,omitempty"`
	// Torrents that are idle for more than specified time aren't counted
	// toward download and upload queue limits
	QueueStalled *time.Duration `json:"-"`
	// Indicates whether or not to consider idle torrents as stalled
	QueueStalledEnabled *bool `json:"queue-stalled-enabled,omitempty"`
	// The default upload ratio limit for torrents
	UploadRatioLimit *float64 `json:"seedRatioLimit,omitempty"`
	// Indicates whether or not to consider upload ration
	UploadRatioLimitEnabled *bool `json:"seedRatioLimited,omitempty"`

	// Indicates whether DHT is allowed for public torrents
	DHTEnabled *bool `json:"dht-enabled,omitempty"`
	// Indicates whether local peer discovery is allowed for public torrents
	LPDEnabled *bool `json:"lpd-enabled,omitempty"`
	// Indicates whether peer exchange is allowed for public torrents
	PEXEnabled *bool `json:"pex-enabled,omitempty"`
	// Indicates whether µTP is allowed
	UTPEnabled *bool `json:"utp-enabled,omitempty"`

	// Peer encryption configuration
	Encryption *Encryption `json:"encryption,omitempty"`

	// Inactive seeding torrents will be stopped after this time
	IdleSeedingLimit *time.Duration `json:"-"`
	// Indicates whether or not inactive seeding limit is enabled
	IdleSeedingLimitEnabled *bool `json:"idle-seeding-limit-enabled,omitempty"`

	// Maximum number of peers across all torrents
	GlobalPeerLimit *int `json:"peer-limit-global,omitempty"`
	// Maximum number of peers for a single torrent
	TorrentPeerLimit *int `json:"peer-limit-per-torrent,omitempty"`

	// Incoming peer port
	PeerPort *int `json:"peer-port,omitempty"`
	// Indicates whether Transmission randomizes peer port on start
	RandomizePeerPort *bool `json:"peer-port-random-on-start,omitempty"`
	// Indicates whether Transmission will try to request port forwading
	// using NAT-PMP or UPnP
	PortForwardingEnabled *bool `json:"port-forwarding-enabled,omitempty"`

	// Path to the script to run when torrent is done downloading
	ScriptPath *string `json:"script-torrent-done-filename,omitempty"`
	// Indicates whether to run script when torrent is done downloading or
	// not
	ScriptEnabled *bool `json:"script-torrent-done-enabled,omitempty"`

	// Indicates whether newly added torrents are started automatically or
	// not
	AutostartTorrents *bool `json:"start-added-torrents,omitempty"`
	// Indicates whether original torrent files are automatically deleted
	// or not
	RemoveTorrentFiles *bool `json:"trash-original-torrent-files,omitempty"`
}

// SetSession applies given configuration to the current Transmission session.
// Only non-nil fields of the request have any effect.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L532
func (c *Client) SetSession(ctx context.Context, req *SetSessionReq) error {
	uc := c.getUnitConversion()

	var setSessionJSON = struct {
		*SetSessionReq

		TurtleDownloadRateLimit *int64         `json:"alt-speed-down,omitempty"`
		TurtleUploadRateLimit   *int64         `json:"alt-speed-up,omitempty"`
		DownloadRateLimit       *int64         `json:"speed-limit-down,omitempty"`
		UploadRateLimit         *int64         `json:"speed-limit-up,omitempty"`
		CacheSize               *int64         `json:"cache-size-mb,omitempty"`
		QueueStalled            *time.Duration `json:"queue-stalled-minutes,omitempty"`
		IdleSeedingLimit        *time.Duration `json:"idle-seeding-limit,omitempty"`
	}{
		SetSessionReq: req,
	}
	if req.TurtleDownloadRateLimit != nil {
		setSessionJSON.TurtleDownloadRateLimit = OptInt64(*req.TurtleDownloadRateLimit / uc.speed)
	}
	if req.TurtleUploadRateLimit != nil {
		setSessionJSON.TurtleUploadRateLimit = OptInt64(*req.TurtleUploadRateLimit / uc.speed)
	}
	if req.DownloadRateLimit != nil {
		setSessionJSON.DownloadRateLimit = OptInt64(*req.DownloadRateLimit / uc.speed)
	}
	if req.UploadRateLimit != nil {
		setSessionJSON.UploadRateLimit = OptInt64(*req.UploadRateLimit / uc.speed)
	}
	if req.CacheSize != nil {
		setSessionJSON.CacheSize = OptInt64(*req.CacheSize / uc.size / uc.size)
	}
	if req.QueueStalled != nil {
		setSessionJSON.QueueStalled = OptDuration(*req.QueueStalled / time.Minute)
	}
	if req.IdleSeedingLimit != nil {
		setSessionJSON.IdleSeedingLimit = OptDuration(*req.IdleSeedingLimit / time.Minute)
	}

	return c.callRPC(ctx, "session-set", &setSessionJSON, nil)
}
