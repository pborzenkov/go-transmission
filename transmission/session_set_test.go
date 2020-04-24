package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSetSession(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "session-set",
			"arguments": {
                          "alt-speed-enabled": true,
                          "alt-speed-time-enabled": true,
                          "alt-speed-time-day": 65,
                          "alt-speed-time-begin": 360,
                          "alt-speed-time-end": 960,
                          "speed-limit-down-enabled": true,
                          "speed-limit-up-enabled": true,
			  "blocklist-url": "http://torrents.com/peers.blocklist",
                          "blocklist-enabled": true,
                          "download-dir": "/home/transmission/downloads",
                          "incomplete-dir": "/home/transmission/incomplete",
                          "incomplete-dir-enabled": true,
                          "rename-partial-files": true,
                          "download-queue-size": 3,
                          "download-queue-enabled": true,
                          "seed-queue-size": 3,
                          "seed-queue-enabled": true,
                          "queue-stalled-enabled": true,
                          "seedRatioLimit": 2,
                          "seedRatioLimited": true,
                          "dht-enabled": true,
                          "lpd-enabled": true,
                          "pex-enabled": true,
                          "utp-enabled": true,
                          "encryption": "preferred",
                          "idle-seeding-limit-enabled": true,
                          "peer-limit-global": 200,
                          "peer-limit-per-torrent": 60,
                          "peer-port": 51970,
                          "peer-port-random-on-start": true,
                          "port-forwarding-enabled": true,
                          "script-torrent-done-filename": "/home/transmission/done.script",
                          "script-torrent-done-enabled": true,
                          "start-added-torrents": true,
                          "trash-original-torrent-files": true,
                          "alt-speed-down": 5120,
                          "alt-speed-up": 5120,
                          "speed-limit-down": 10240,
                          "speed-limit-up": 10240,
                          "cache-size-mb": 4,
                          "queue-stalled-minutes": 30,
                          "idle-seeding-limit": 30
			}
		  }`)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	err := client.SetSession(context.Background(), &SetSessionReq{
		TurtleDownloadRateLimit: OptInt64(5120000),
		TurtleUploadRateLimit:   OptInt64(5120000),
		TurtleEnabled:           OptBool(true),
		TurtleScheduleEnabled:   OptBool(true),
		TurtleScheduleOnDays:    OptWeekday(Saturday | Sunday),
		TurtleScheduleStartsAt:  OptInt(360),
		TurtleScheduleStopsAt:   OptInt(960),

		DownloadRateLimit:        OptInt64(10240000),
		DownloadRateLimitEnabled: OptBool(true),
		UploadRateLimit:          OptInt64(10240000),
		UploadRateLimitEnabled:   OptBool(true),

		BlocklistURL:     OptString("http://torrents.com/peers.blocklist"),
		BlocklistEnabled: OptBool(true),

		CacheSize: OptInt64(4000000),

		DownloadDirectory:          OptString("/home/transmission/downloads"),
		IncompleteDirectory:        OptString("/home/transmission/incomplete"),
		IncompleteDirectoryEnabled: OptBool(true),
		RenameIncompleteFiles:      OptBool(true),

		DownloadQueueLimit:        OptInt(3),
		DownloadQueueLimitEnabled: OptBool(true),
		UploadQueueLimit:          OptInt(3),
		UploadQueueLimitEnabled:   OptBool(true),
		QueueStalled:              OptDuration(30 * time.Minute),
		QueueStalledEnabled:       OptBool(true),
		UploadRatioLimit:          OptFloat64(2),
		UploadRatioLimitEnabled:   OptBool(true),

		DHTEnabled: OptBool(true),
		LPDEnabled: OptBool(true),
		PEXEnabled: OptBool(true),
		UTPEnabled: OptBool(true),

		Encryption: OptEncryption(EncryptionPreferred),

		IdleSeedingLimit:        OptDuration(30 * time.Minute),
		IdleSeedingLimitEnabled: OptBool(true),

		GlobalPeerLimit:  OptInt(200),
		TorrentPeerLimit: OptInt(60),

		PeerPort:              OptInt(51970),
		RandomizePeerPort:     OptBool(true),
		PortForwardingEnabled: OptBool(true),

		ScriptPath:    OptString("/home/transmission/done.script"),
		ScriptEnabled: OptBool(true),

		AutostartTorrents:  OptBool(true),
		RemoveTorrentFiles: OptBool(true),
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
