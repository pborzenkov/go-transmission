package transmission

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetTorrents(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-get",
			"arguments": {
			  "ids": 1,
			  "fields": [
			    "id",
			    "hashString",
			    "name",
			    "status",
			    "creator",
			    "comment",
			    "labels",
			    "eta",
			    "etaIdle",
			    "error",
			    "errorString",
			    "torrentFile",
			    "magnetLink",
			    "downloadDir",
			    "dateCreated",
			    "addedDate",
			    "startDate",
			    "activityDate",
			    "doneDate",
			    "manualAnnounceTime",
			    "rateDownload",
			    "rateUpload",
			    "downloadLimit",
			    "downloadLimited",
			    "uploadLimit",
			    "uploadLimited",
			    "honorsSessionLimits",
			    "downloadedEver",
			    "uploadedEver",
			    "corruptEver",
			    "bandwidthPriority",
			    "queuePosition",
			    "seedIdleLimit",
			    "seedIdleMode",
			    "seedRatioLimit",
			    "seedRatioMode",
			    "uploadRatio",
			    "secondsDownloading",
			    "secondsSeeding",
			    "totalSize",
			    "sizeWhenDone",
			    "desiredAvailable",
			    "leftUntilDone",
			    "haveUnchecked",
			    "haveValid",
			    "percentDone",
			    "recheckProgress",
			    "metadataPercentComplete",
			    "isFinished",
			    "isPrivate",
			    "isStalled",
			    "peer-limit",
			    "peersConnected",
			    "peersGettingFromUs",
			    "peersSendingToUs",
			    "peers",
			    "peersFrom",
			    "webseedsSendingToUs",
			    "webseeds",
			    "wanted",
			    "files",
			    "fileStats",
			    "priorities",
			    "pieceCount",
			    "pieceSize",
			    "pieces",
			    "trackers",
			    "trackerStats"
			  ]
		        }
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "torrents": [{
			    "id": 1,
                            "hashString": "b24c456970f6013e14d01ab2defe2ceb67bb6e07",
			    "name": "A-test-torrent",
			    "status": 4,
			    "creator": "go-transmission",
			    "comment": "Just a test torrent",
			    "labels": ["iso", "linux"],
			    "eta": 503,
			    "etaIdle": -1,
			    "error": 1,
			    "errorString": "A warning from tracker",
			    "torrentFile": "/home/transmission/torrents/example.torrent",
			    "magnetLink":"magnet:?xt=urn:btih:b24c456970f6013e14d01ab2defe2ceb67bb6e07&dn=A-test-torrent",
			    "downloadDir": "/home/transmission/download",
			    "dateCreated": 1505188945,
			    "addedDate": 1576856565,
			    "startDate": 1587996143,
			    "activityDate": 1587996507,
			    "doneDate":0,
			    "manualAnnounceTime": 1587996707,
			    "rateDownload": 3804000,
			    "rateUpload": 339000,
			    "downloadLimit": 10240,
			    "downloadLimited": true,
			    "uploadLimit": 10240,
			    "uploadLimited": true,
			    "honorsSessionLimits": true,
			    "downloadedEver": 1439775828,
			    "uploadedEver": 345775123,
			    "corruptEver": 12345,
			    "bandwidthPriority": 1,
			    "queuePosition": 1,
			    "seedIdleLimit": 30,
			    "seedIdleMode": 2,
			    "seedRatioLimit": 2.42,
			    "seedRatioMode": 1,
			    "uploadRatio": 1.23,
			    "secondsDownloading": 591,
			    "secondsSeeding": 10,
			    "totalSize": 31066499565,
			    "sizeWhenDone": 30066499565,
			    "desiredAvailable": 29627170816,
			    "leftUntilDone": 29627170816,
			    "haveUnchecked": 76890112,
			    "haveValid": 1362438637,
			    "percentDone": 0.23,
			    "recheckProgress": 0.1,
			    "metadataPercentComplete": 1,
			    "isFinished": true,
			    "isPrivate": true,
			    "isStalled": true,
			    "peer-limit": 60,
			    "peersConnected":30,
			    "peersGettingFromUs": 7,
			    "peersSendingToUs": 17,
			    "peers": [
			      {
			        "address": "31.9.72.249",
				"clientIsChoked": true,
				"clientIsInterested": true,
				"clientName": "\u00b5Torrent 3.5.5",
				"flagStr": "TdUEH",
				"isDownloadingFrom": false,
				"isEncrypted": true,
				"isIncoming": false,
				"isUTP": true,
				"isUploadingTo": true,
				"peerIsChoked": false,
				"peerIsInterested": true,
				"port": 40224,
				"progress": 0.2653,
				"rateToClient": 0,
				"rateToPeer": 11000
			      },
			      {
				"address": "37.28.155.78",
				"clientIsChoked": true,
				"clientIsInterested": true,
				"clientName": "libTorrent (Rakshasa) 0.13.7",
				"flagStr": "DXI",
				"isDownloadingFrom": true,
				"isEncrypted": false,
				"isIncoming": true,
				"isUTP": false,
				"isUploadingTo": false,
				"peerIsChoked": true,
				"peerIsInterested":false,
				"port":6991,
				"progress": 1,
				"rateToClient": 434000,
				"rateToPeer": 0
			      }
		            ],
			    "peersFrom": {
			      "fromCache": 3,
			      "fromDht": 4,
			      "fromIncoming": 1,
			      "fromLpd": 5,
			      "fromLtep": 2,
			      "fromPex": 10,
			      "fromTracker": 7
			    },
			    "webseedsSendingToUs": 2,
			    "webseeds":["http://seed1", "http://seed2"],
			    "wanted": [1, 0, 1],
			    "files": [
			      {
			        "bytesCompleted": 191450496,
				"length": 3093908864,
				"name": "A-test-torrent/file1"
			      },
			      {
				"bytesCompleted": 73684668,
				"length": 3395573436,
				"name": "A-test-torrent/file2"
			      },
			      {
				"bytesCompleted": 124850352,
				"length": 2683113648,
				"name": "A-test-torrent/file3"
			      }
			    ],
			    "fileStats": [
			      {
			        "bytesCompleted": 191450496,
				"priority": -1,
				"wanted": true
			      },
			      {
				"bytesCompleted": 73684668,
				"priority": 0,
				"wanted": false
			      },
			      {
				"bytesCompleted": 124850352,
				"priority": 1,
				"wanted": true
			      }
			    ],
			    "priorities": [-1, 0, 1],
			    "pieceCount":3704,
			    "pieceSize":8388608,
			    "pieces": "gAIwQgg=",
			    "trackers": [
			      {
			        "id": 0,
			        "tier": 0,
			        "announce": "http://tracker.trackerfix.com:80/announce",
			        "scrape": "http://tracker.trackerfix.com:80/scrape"
		              },
			      {
			        "id": 1,
			        "tier": 1,
			        "announce": "udp://9.rarbg.to:2740",
			        "scrape": "udp://9.rarbg.to:2740"
			      }
			    ],
			    "trackerStats": [
			      {
                                "announce": "http://tracker.trackerfix.com:80/announce",
				"announceState": 3,
				"downloadCount": -1,
				"hasAnnounced": true,
				"hasScraped": true,
				"host": "http://tracker.trackerfix.com:80",
				"id": 0,
				"isBackup": false,
				"lastAnnouncePeerCount": 7,
				"lastAnnounceResult": "timed out",
				"lastAnnounceStartTime": 1588064117,
				"lastAnnounceSucceeded": true,
				"lastAnnounceTime": 1588064119,
				"lastAnnounceTimedOut": true,
				"lastScrapeResult": "some result",
				"lastScrapeStartTime": 0,
				"lastScrapeSucceeded": true,
				"lastScrapeTime": 0,
				"lastScrapeTimedOut": 1,
				"leecherCount": 150,
				"nextAnnounceTime": 0,
				"nextScrapeTime": 1588064150,
				"scrape": "http://tracker.trackerfix.com:80/scrape",
				"scrapeState": 1,
				"seederCount": 30,
				"tier": 0
			      },
			      {
				"announce": "udp://9.rarbg.me:2770",
				"announceState": 1,
				"downloadCount": 30,
				"hasAnnounced": true,
				"hasScraped": true,
				"host": "udp://9.rarbg.me:2770",
				"id": 1,
				"isBackup": true,
				"lastAnnouncePeerCount": 0,
				"lastAnnounceResult": "Could not connect to tracker",
				"lastAnnounceStartTime": 0,
				"lastAnnounceSucceeded": true,
				"lastAnnounceTime": 1588064137,
				"lastAnnounceTimedOut": false,
				"lastScrapeResult": "Could not connect to tracker",
				"lastScrapeStartTime": 1588064150,
				"lastScrapeSucceeded": false,
				"lastScrapeTime": 1588064160,
				"lastScrapeTimedOut": true,
				"leecherCount": -1,
				"nextAnnounceTime": 1588064463,
				"nextScrapeTime": 1588065060,
				"scrape": "udp://9.rarbg.me:2770",
				"scrapeState": 1,
				"seederCount": -1,
				"tier": 1
			      }
			    ]
		          }]
		        }
		  }`)
	})

	got, err := client.GetTorrents(context.Background(), ID(1))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := []*Torrent{
		{
			ID:      ID(1),
			Hash:    "b24c456970f6013e14d01ab2defe2ceb67bb6e07",
			Name:    "A-test-torrent",
			Status:  StatusDownload,
			Creator: "go-transmission",
			Comment: "Just a test torrent",
			Labels:  []string{"iso", "linux"},

			ETA:     503 * time.Second,
			IdleETA: -1,

			ErrorType: ErrorTypeTrackerWarning,
			Error:     "A warning from tracker",

			File:              "/home/transmission/torrents/example.torrent",
			MagnetLink:        "magnet:?xt=urn:btih:b24c456970f6013e14d01ab2defe2ceb67bb6e07&dn=A-test-torrent",
			DownloadDirectory: "/home/transmission/download",

			CreatedAt:             time.Date(2017, 9, 12, 4, 2, 25, 0, time.UTC),
			AddedAt:               time.Date(2019, 12, 20, 15, 42, 45, 0, time.UTC),
			StartedAt:             time.Date(2020, 4, 27, 14, 2, 23, 0, time.UTC),
			LastActiveAt:          time.Date(2020, 4, 27, 14, 8, 27, 0, time.UTC),
			DoneAt:                time.Time{},
			CanManuallyAnnounceAt: time.Date(2020, 4, 27, 14, 11, 47, 0, time.UTC),

			DownloadRate:             3804000,
			UploadRate:               339000,
			DownloadRateLimit:        10240000,
			DownloadRateLimitEnabled: true,
			UploadRateLimit:          10240000,
			UploadRateLimited:        true,
			HonorSessionLimits:       true,

			DownloadedTotal: 1439775828,
			UploadedTotal:   345775123,
			CorruptedTotal:  12345,

			Priority:        PriorityHigh,
			PositionInQueue: 1,

			IdleSeedingLimit:     30 * time.Minute,
			IdleSeedingLimitMode: LimitUnlimited,
			UploadRatioLimit:     2.42,
			UploadRatioLimitMode: LimitLocal,
			UploadRatio:          1.23,

			DownloadingFor: 591 * time.Second,
			SeedingFor:     10 * time.Second,

			TotalSize:       31066499565,
			WantedSize:      30066499565,
			WantedAvailable: 29627170816,
			WantedLeft:      29627170816,
			UncheckedSize:   76890112,
			ValidSize:       1362438637,

			DataDone:     0.23,
			DataChecked:  0.1,
			MetadataDone: 1,

			IsFinished: true,
			IsPrivate:  true,
			IsStalled:  true,

			PeerLimit:          60,
			ConnectedPeers:     30,
			PeersGettingFromUs: 7,
			PeersSendingToUs:   17,
			Peers: []Peer{
				{
					Address:          net.IP{31, 9, 72, 249},
					Port:             40224,
					ClientName:       "\u00b5Torrent 3.5.5",
					DownloadRate:     0,
					UploadRate:       11000,
					Progress:         0.2653,
					IsDownloading:    false,
					IsUploading:      true,
					IsUTP:            true,
					IsIncoming:       false,
					IsEncrypted:      true,
					AreWeChoked:      true,
					IsPeerChoked:     false,
					AreWeInterested:  true,
					IsPeerInterested: true,
				},
				{
					Address:          net.IP{37, 28, 155, 78},
					Port:             6991,
					ClientName:       "libTorrent (Rakshasa) 0.13.7",
					DownloadRate:     434000,
					UploadRate:       0,
					Progress:         1,
					IsDownloading:    true,
					IsUploading:      false,
					IsUTP:            false,
					IsIncoming:       true,
					IsEncrypted:      false,
					AreWeChoked:      true,
					IsPeerChoked:     true,
					AreWeInterested:  true,
					IsPeerInterested: false,
				},
			},
			PeersFrom: PeersOrigin{
				Tracker:  7,
				Incoming: 1,
				Cache:    3,
				DHT:      4,
				LPD:      5,
				PEX:      10,
				LTEP:     2,
			},
			WebSeedsSendingToUs: 2,
			WebSeeds: []string{
				"http://seed1",
				"http://seed2",
			},

			Wanted: []bool{true, false, true},
			Files: []File{
				{
					Name:       "A-test-torrent/file1",
					Size:       3093908864,
					Downloaded: 191450496,
				},
				{
					Name:       "A-test-torrent/file2",
					Size:       3395573436,
					Downloaded: 73684668,
				},
				{
					Name:       "A-test-torrent/file3",
					Size:       2683113648,
					Downloaded: 124850352,
				},
			},
			FileStats: []FileStat{
				{
					Downloaded: 191450496,
					Priority:   PriorityLow,
					Wanted:     true,
				},
				{
					Downloaded: 73684668,
					Priority:   PriorityNormal,
					Wanted:     false,
				},
				{
					Downloaded: 124850352,
					Priority:   PriorityHigh,
					Wanted:     true,
				},
			},
			Priorities: []Priority{
				PriorityLow,
				PriorityNormal,
				PriorityHigh,
			},

			PieceCount: 3704,
			PieceSize:  8388608,
			Pieces:     []byte{0x80, 0x02, 0x30, 0x42, 0x08},

			Trackers: []Tracker{
				{
					ID:          0,
					Tier:        0,
					AnnounceURL: parseTestURL(t, "http://tracker.trackerfix.com:80/announce"),
					ScrapeURL:   parseTestURL(t, "http://tracker.trackerfix.com:80/scrape"),
				},
				{
					ID:          1,
					Tier:        1,
					AnnounceURL: parseTestURL(t, "udp://9.rarbg.to:2740"),
					ScrapeURL:   parseTestURL(t, "udp://9.rarbg.to:2740"),
				},
			},
			TrackerStats: []TrackerStat{
				{
					ID:          0,
					Tier:        0,
					IsBackup:    false,
					Host:        parseTestURL(t, "http://tracker.trackerfix.com:80"),
					AnnounceURL: parseTestURL(t, "http://tracker.trackerfix.com:80/announce"),
					ScrapeURL:   parseTestURL(t, "http://tracker.trackerfix.com:80/scrape"),

					Leechers:  150,
					Seeders:   30,
					Downloads: -1,

					HasAnnounced:            true,
					AnnounceState:           TrackerStateActive,
					LastAnnounceStartTime:   time.Date(2020, 04, 28, 8, 55, 17, 0, time.UTC),
					LastAnnounceTime:        time.Date(2020, 04, 28, 8, 55, 19, 0, time.UTC),
					IsLastAnnounceTimedOut:  true,
					IsLastAnnounceSucceeded: true,
					LastAnnounceResult:      "timed out",
					LastAnnouncePeerCount:   7,
					NextAnnounceTime:        time.Time{},

					HasScraped:            true,
					ScrapeState:           TrackerStateWaiting,
					LastScrapeStartTime:   time.Time{},
					LastScrapeTime:        time.Time{},
					IsLastScrapeTimedOut:  true,
					IsLastScrapeSucceeded: true,
					LastScrapeResult:      "some result",
					NextScrapeTime:        time.Date(2020, 04, 28, 8, 55, 50, 0, time.UTC),
				},
				{
					ID:          1,
					Tier:        1,
					IsBackup:    true,
					Host:        parseTestURL(t, "udp://9.rarbg.me:2770"),
					AnnounceURL: parseTestURL(t, "udp://9.rarbg.me:2770"),
					ScrapeURL:   parseTestURL(t, "udp://9.rarbg.me:2770"),

					Leechers:  -1,
					Seeders:   -1,
					Downloads: 30,

					HasAnnounced:            true,
					AnnounceState:           TrackerStateWaiting,
					LastAnnounceStartTime:   time.Time{},
					LastAnnounceTime:        time.Date(2020, 04, 28, 8, 55, 37, 0, time.UTC),
					IsLastAnnounceTimedOut:  false,
					IsLastAnnounceSucceeded: true,
					LastAnnounceResult:      "Could not connect to tracker",
					LastAnnouncePeerCount:   0,
					NextAnnounceTime:        time.Date(2020, 04, 28, 9, 01, 03, 0, time.UTC),

					HasScraped:            true,
					ScrapeState:           TrackerStateWaiting,
					LastScrapeStartTime:   time.Date(2020, 04, 28, 8, 55, 50, 0, time.UTC),
					LastScrapeTime:        time.Date(2020, 04, 28, 8, 56, 0, 0, time.UTC),
					IsLastScrapeTimedOut:  true,
					IsLastScrapeSucceeded: false,
					LastScrapeResult:      "Could not connect to tracker",
					NextScrapeTime:        time.Date(2020, 04, 28, 9, 11, 0, 0, time.UTC),
				},
			},
		},
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected torrent data, diff = \n%s", cmp.Diff(want, got))
	}
	if !got[0].Pieces.IsDownloaded(0) {
		t.Errorf("expected first piece to be downloaded")
	}
	if got[0].Pieces.IsDownloaded(1) {
		t.Errorf("expected second piece to not be downloaded")
	}
	if got[0].Pieces.IsDownloaded(100) {
		t.Errorf("expected non-existing piece to be reported as not downloaded")
	}
}

func TestGetTorrents_all(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-get",
			"arguments": {
			  "fields": [
			    "id"
			  ]
			}
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "torrents": [
			    {"id": 1},
			    {"id": 2},
			    {"id": 3}
			  ]
		        }
		  }`)
	})

	got, err := client.GetTorrents(context.Background(), All(), TorrentFieldID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := []*Torrent{
		{ID: ID(1)},
		{ID: ID(2)},
		{ID: ID(3)},
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected torrent data, diff = \n%s", cmp.Diff(want, got))
	}
}

func TestGetRecentlyRemovedTorrnetIDs(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-get",
			"arguments": {
			  "ids": "recently-active",
			  "fields": [
			    "id"
			  ]
		        }
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "removed": [3, 7, 10]
		        }
		  }`)
	})

	got, err := client.GetRecentlyRemovedTorrentIDs(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := []ID{3, 7, 10}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected list of removed torrents, diff = \n%s", cmp.Diff(want, got))
	}
}
