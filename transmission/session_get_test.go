package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetSession(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"session-get"}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
                          "alt-speed-down": 5120,
                          "alt-speed-enabled": true,
                          "alt-speed-time-begin": 360,
                          "alt-speed-time-day": 65,
                          "alt-speed-time-enabled": true,
                          "alt-speed-time-end": 960,
                          "alt-speed-up": 5120,
                          "blocklist-enabled": true,
                          "blocklist-size": 1,
			  "blocklist-url": "http://torrents.com/peers.blocklist",
                          "cache-size-mb": 4,
                          "config-dir": "/home/transmission/config",
                          "dht-enabled": true,
                          "download-dir": "/home/transmission/downloads",
                          "download-queue-enabled": true,
                          "download-queue-size": 3,
                          "encryption": "preferred",
                          "idle-seeding-limit": 30,
                          "idle-seeding-limit-enabled": true,
                          "incomplete-dir": "/home/transmission/incomplete",
                          "incomplete-dir-enabled": true,
                          "lpd-enabled": true,
                          "peer-limit-global": 200,
                          "peer-limit-per-torrent": 60,
                          "peer-port": 51970,
                          "peer-port-random-on-start": true,
                          "pex-enabled": true,
                          "port-forwarding-enabled": true,
                          "queue-stalled-enabled": true,
                          "queue-stalled-minutes": 30,
                          "rename-partial-files": true,
                          "rpc-version": 15,
                          "rpc-version-minimum": 1,
                          "script-torrent-done-enabled": true,
                          "script-torrent-done-filename": "/home/transmission/done.script",
                          "seed-queue-enabled": true,
                          "seed-queue-size": 3,
                          "seedRatioLimit": 2,
                          "seedRatioLimited": true,
                          "speed-limit-down": 10240,
                          "speed-limit-down-enabled": true,
                          "speed-limit-up": 10240,
                          "speed-limit-up-enabled": true,
                          "start-added-torrents": true,
                          "trash-original-torrent-files": true,
                          "units": {
                            "memory-bytes": 1000,
                            "memory-units": [
                              "KB",
                              "MB",
                              "GB",
                              "TB"
                            ],
                            "size-bytes": 1000,
                            "size-units": [
                              "KB",
                              "MB",
                              "GB",
                              "TB"
                            ],
                            "speed-bytes": 1000,
                            "speed-units": [
                              "KB/s",
                              "MB/s",
                              "GB/s",
                              "TB/s"
                            ]
                          },
                          "utp-enabled": true,
                          "version": "2.94 (d8e60ee44f)"
			}
		  }`)
	})

	got, err := client.GetSession(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &Session{
		TurtleDownloadRateLimit: 5120000,
		TurtleUploadRateLimit:   5120000,
		TurtleEnabled:           true,
		TurtleScheduleEnabled:   true,
		TurtleScheduleOnDays:    Saturday | Sunday,
		TurtleScheduleStartsAt:  360,
		TurtleScheduleStopsAt:   960,

		DownloadRateLimit:        10240000,
		DownloadRateLimitEnabled: true,
		UploadRateLimit:          10240000,
		UploadRateLimitEnabled:   true,

		BlocklistURL:     "http://torrents.com/peers.blocklist",
		BlocklistEnabled: true,
		BlocklistSize:    1,

		CacheSize: 4000000,

		ConfigDirectory:            "/home/transmission/config",
		DownloadDirectory:          "/home/transmission/downloads",
		IncompleteDirectory:        "/home/transmission/incomplete",
		IncompleteDirectoryEnabled: true,
		RenameIncompleteFiles:      true,

		DownloadQueueLimit:        3,
		DownloadQueueLimitEnabled: true,
		UploadQueueLimit:          3,
		UploadQueueLimitEnabled:   true,
		QueueStalled:              30 * time.Minute,
		QueueStalledEnabled:       true,
		UploadRatio:               2,
		UploadRatioEnabled:        true,

		DHTEnabled: true,
		LPDEnabled: true,
		PEXEnabled: true,
		UTPEnabled: true,

		Encryption: EncryptionPreferred,

		IdleSeedingLimit:        30 * time.Minute,
		IdleSeedingLimitEnabled: true,

		GlobalPeerLimit:  200,
		TorrentPeerLimit: 60,

		PeerPort:              51970,
		RandomizePeerPort:     true,
		PortForwardingEnabled: true,

		ScriptPath:    "/home/transmission/done.script",
		ScriptEnabled: true,

		AutostartTorrents:  true,
		RemoveTorrentFiles: true,

		RPCVersion:        15,
		RPCVersionMinimum: 1,
		Version:           "2.94 (d8e60ee44f)",

		Units: SessionUnits{
			Speed:            []string{"KB/s", "MB/s", "GB/s", "TB/s"},
			SpeedBytesPerKB:  1000,
			Size:             []string{"KB", "MB", "GB", "TB"},
			SizeBytesPerKB:   1000,
			Memory:           []string{"KB", "MB", "GB", "TB"},
			MemoryBytesPerKB: 1000,
		},
	}

	if !cmp.Equal(want, got) {
		t.Errorf("unexpected session data, diff = \n%s", cmp.Diff(want, got))
	}
}
