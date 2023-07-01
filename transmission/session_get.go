package transmission

import (
	"context"
	"time"
)

// Session holds Transmission session information.
type Session struct {
	// ID of the current session
	ID string `json:"session-id"`

	// Maximum allowed download rate in "turtle" mode (bytes/s)
	TurtleDownloadRateLimit int64 `json:"alt-speed-down"`
	// Maximum allowed upload rate in "turtle" mode (bytes/s)
	TurtleUploadRateLimit int64 `json:"alt-speed-up"`
	// Indicates whether "turtle" mode is enabled right now
	TurtleEnabled bool `json:"alt-speed-enabled"`
	// Indicates whether "turtle" mode is controlled by the configured
	// schedule
	TurtleScheduleEnabled bool `json:"alt-speed-time-enabled"`
	// Indicates on what days of the week to turn on "turtle" mode
	TurtleScheduleOnDays Weekday `json:"alt-speed-time-day"`
	// When to turn on "turtle" mode (in minutes after midnight)
	TurtleScheduleStartsAt int `json:"alt-speed-time-begin"`
	// When to turn off "turtle" mode (in minutes after midnight)
	TurtleScheduleStopsAt int `json:"alt-speed-time-end"`

	// Maximum allowed download rate (bytes/s)
	DownloadRateLimit int64 `json:"speed-limit-down"`
	// Indicates whether download rate limit is enabled
	DownloadRateLimitEnabled bool `json:"speed-limit-down-enabled"`
	// Maximum allowed upload rate (bytes/s)
	UploadRateLimit int64 `json:"speed-limit-up"`
	// Indicates whether upload rate limit is enabled
	UploadRateLimitEnabled bool `json:"speed-limit-up-enabled"`

	// Location of the peer blocklist
	BlocklistURL string `json:"blocklist-url"`
	// Indicates whether or not peer blocklist is enabled
	BlocklistEnabled bool `json:"blocklist-enabled"`
	// Number of entries in the peer blocklist
	BlocklistSize int `json:"blocklist-size"`

	// Maximum size of disk cache in bytes
	CacheSize int64 `json:"cache-size-mb"`

	// Location of Transmission config directory
	ConfigDirectory string `json:"config-dir"`
	// Default path to download torrents
	DownloadDirectory string `json:"download-dir"`
	// Path for incomplete torrents (if enabled)
	IncompleteDirectory string `json:"incomplete-dir"`
	// Indicates whether to keep torrents in incomplete directory until done
	IncompleteDirectoryEnabled bool `json:"incomplete-dir-enabled"`
	// Indicates whether Transmission will append '.part' suffix to
	// incomplete files
	RenameIncompleteFiles bool `json:"rename-partial-files"`

	// Max number of torrents to download at once
	DownloadQueueLimit int `json:"download-queue-size"`
	// Indicates whether or not download queue limit is enabled
	DownloadQueueLimitEnabled bool `json:"download-queue-enabled"`
	// Max number of torrents to seed at once
	UploadQueueLimit int `json:"seed-queue-size"`
	// Indicates whether or not upload queue limit is enabled
	UploadQueueLimitEnabled bool `json:"seed-queue-enabled"`
	// Torrents that are idle for more than specified time aren't counted
	// toward download and upload queue limits
	QueueStalled time.Duration `json:"queue-stalled-minutes"`
	// Indicates whether or not to consider idle torrents as stalled
	QueueStalledEnabled bool `json:"queue-stalled-enabled"`
	// The default upload ratio for torrents
	UploadRatio float64 `json:"seedRatioLimit"`
	// Indicates whether or not to consider upload ration
	UploadRatioEnabled bool `json:"seedRatioLimited"`

	// Indicates whether DHT is allowed for public torrents
	DHTEnabled bool `json:"dht-enabled"`
	// Indicates whether local peer discovery is allowed for public torrents
	LPDEnabled bool `json:"lpd-enabled"`
	// Indicates whether peer exchange is allowed for public torrents
	PEXEnabled bool `json:"pex-enabled"`
	// Indicates whether µTP is allowed
	UTPEnabled bool `json:"utp-enabled"`

	// Peer encryption configuration
	Encryption Encryption `json:"encryption"`

	// Inactive seeding torrents will be stopped after this time
	IdleSeedingLimit time.Duration `json:"idle-seeding-limit"`
	// Indicates whether or not inactive seeding limit is enabled
	IdleSeedingLimitEnabled bool `json:"idle-seeding-limit-enabled"`

	// Maximum number of peers across all torrents
	GlobalPeerLimit int `json:"peer-limit-global"`
	// Maximum number of peers for a single torrent
	TorrentPeerLimit int `json:"peer-limit-per-torrent"`

	// Incoming peer port
	PeerPort int `json:"peer-port"`
	// Indicates whether Transmission randomizes peer port on start
	RandomizePeerPort bool `json:"peer-port-random-on-start"`
	// Indicates whether Transmission will try to request port forwading
	// using NAT-PMP or UPnP
	PortForwardingEnabled bool `json:"port-forwarding-enabled"`

	// Path to the script to run when torrent is done downloading
	ScriptPath string `json:"script-torrent-done-filename"`
	// Indicates whether to run script when torrent is done downloading or
	// not
	ScriptEnabled bool `json:"script-torrent-done-enabled"`

	// Indicates whether newly added torrents are started automatically or
	// not
	AutostartTorrents bool `json:"start-added-torrents"`
	// Indicates whether original torrent files are automatically deleted
	// or not
	RemoveTorrentFiles bool `json:"trash-original-torrent-files"`

	// Current RPC API version
	RPCVersion int `json:"rpc-version"`
	// Minimum supported RPC version
	RPCVersionMinimum int `json:"rpc-version-minimum"`
	// Transmission version
	Version string `json:"version"`

	// Value conversion units
	Units SessionUnits `json:"units"`
}

//go:generate go run ../tools/gen-fields.go -type Session

// SessionUnits holds value conversion units.
type SessionUnits struct {
	// KB/s, MB/s, GB/s, TB/s
	Speed []string `json:"speed-units"`
	// Number of bytes per KB
	SpeedBytesPerKB int `json:"speed-bytes"`
	// KB, MB, GB, TB
	Size []string `json:"size-units"`
	// Number of bytes per KB
	SizeBytesPerKB int `json:"size-bytes"`
	// KB, MB, GB, TB
	Memory []string `json:"memory-units"`
	// Number of bytes per KB
	MemoryBytesPerKB int `json:"memory-bytes"`
}

// GetSession returns detailed information about current Transmission session.
//
// https://github.com/transmission/transmission/blob/46b3e6c8dae02531b1eb8907b51611fb9229b54a/extras/rpc-spec.txt#L540
func (c *Client) GetSession(ctx context.Context, fields ...SessionField) (*Session, error) {
	var getSessionReq = struct {
		Fields []SessionField `json:"fields,omitempty"`
	}{fields}

	resp := new(Session)
	if err := c.callRPC(ctx, "session-get", getSessionReq, resp); err != nil {
		return nil, err
	}

	uc := unitConversion{
		speed:  int64(resp.Units.SpeedBytesPerKB),
		size:   int64(resp.Units.SizeBytesPerKB),
		memory: int64(resp.Units.MemoryBytesPerKB),
	}

	resp.TurtleDownloadRateLimit *= uc.speed
	resp.TurtleUploadRateLimit *= uc.speed
	resp.DownloadRateLimit *= uc.speed
	resp.UploadRateLimit *= uc.speed
	resp.CacheSize *= uc.size * uc.size
	resp.QueueStalled *= time.Minute
	resp.IdleSeedingLimit *= time.Minute

	c.setUnitConversion(uc)

	return resp, nil
}
