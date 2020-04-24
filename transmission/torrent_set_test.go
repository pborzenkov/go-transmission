package transmission

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestSetTorrent(t *testing.T) {
	parseURL := func(t *testing.T, str string) *url.URL {
		t.Helper()

		u, err := url.Parse(str)
		if err != nil {
			t.Fatalf("failed to parse URL %q: %v", str, err)
		}
		return u
	}

	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-set",
			"arguments": {
			  "downloadLimited": true,
			  "uploadLimited": true,
			  "honorsSessionLimits": true,
                          "bandwidthPriority": -1,
			  "priority-high": [1, 4, 7],
			  "priority-normal": [2, 5, 8],
			  "priority-low": [3, 6, 9],
			  "queuePosition": 3,
			  "files-wanted": [1, 2, 3, 4, 5, 6, 7, 8, 9],
			  "files-unwanted": [10, 11],
			  "peer-limit": 100,
			  "location": "/home/transmission/new",
			  "seedIdleMode": 0,
			  "seedRatioLimit": 2.45,
			  "seedRatioMode": 1,
			  "trackerRemove": [1, 2, 3],
			  "ids": [2, "abcde"],
                          "downloadLimit": 10240,
			  "uploadLimit": 10240,
			  "seedIdleLimit": 60,
			  "trackerAdd": ["http://retracker.local:80", "http://bt2.t-ru.org:80"],
			  "trackerReplace": [1, "http://retracker.local:80", 2, "http://bt2.t-ru.org:80"]
			}
		  }`)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	err := client.SetTorrents(context.Background(), IDs(ID(2), Hash("abcde")), &SetTorrentReq{
		DownloadRateLimit:        OptInt64(10240000),
		DownloadRateLimitEnabled: OptBool(true),
		UploadRateLimit:          OptInt64(10240000),
		UploadRateLimitEnabled:   OptBool(true),
		HonorSessionLimits:       OptBool(true),

		Priority:            OptPriority(PriorityLow),
		HighPriorityFiles:   []int{1, 4, 7},
		NormalPriorityFiles: []int{2, 5, 8},
		LowPriorityFiles:    []int{3, 6, 9},
		PositionInQueue:     OptInt(3),

		WantedFiles:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		UnwantedFiles: []int{10, 11},

		PeerLimit: OptInt(100),

		Location: OptString("/home/transmission/new"),

		IdleSeedingLimit:     OptDuration(60 * time.Minute),
		IdleSeedingLimitMode: OptLimit(LimitGlobal),
		UploadRatioLimit:     OptFloat64(2.45),
		UploadRatioLimitMode: OptLimit(LimitLocal),

		TrackersToAdd: []*url.URL{
			parseURL(t, "http://retracker.local:80"),
			parseURL(t, "http://bt2.t-ru.org:80"),
		},
		TrackerToRemove: []int{1, 2, 3},
		TrackersToReplace: []TrackerReplacement{
			{ID: 1, AnnounceURL: parseURL(t, "http://retracker.local:80")},
			{ID: 2, AnnounceURL: parseURL(t, "http://bt2.t-ru.org:80")},
		},
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
