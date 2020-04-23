package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetSessionStats(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"method":"session-stats"}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "activeTorrentCount":3,
			  "torrentCount":10,
			  "pausedTorrentCount":7,
			  "downloadSpeed":11534336,
			  "uploadSpeed":7340032,
			  "current-stats": {
			    "downloadedBytes": 26264875081,
			    "uploadedBytes": 1914479012,
			    "filesAdded": 13,
			    "secondsActive": 60076,
			    "sessionCount": 1
		          },
			  "cumulative-stats": {
			    "downloadedBytes": 2431199423108,
			    "uploadedBytes": 3588185894743,
			    "filesAdded": 14964,
			    "secondsActive": 8392124,
			    "sessionCount":409
			  }
			}
		  }`)
	})

	got, err := client.GetSessionStats(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &SessionStats{
		Torrents:       10,
		ActiveTorrents: 3,
		PausedTorrents: 7,
		DownloadRate:   11534336,
		UploadRate:     7340032,
		CurrentSession: Stats{
			Downloaded: 26264875081,
			Uploaded:   1914479012,
			Files:      13,
			Sessions:   1,
			ActiveFor:  60076 * time.Second,
		},
		AllSessions: Stats{
			Downloaded: 2431199423108,
			Uploaded:   3588185894743,
			Files:      14964,
			Sessions:   409,
			ActiveFor:  8392124 * time.Second,
		},
	}

	if !cmp.Equal(want, got) {
		t.Errorf("unexpected session stats, diff = \n%s", cmp.Diff(want, got))
	}
}
