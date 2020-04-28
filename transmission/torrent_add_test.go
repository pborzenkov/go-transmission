package transmission

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddTorrent_link(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-add",
			"arguments": {
			  "filename": "magnet:?xt=urn:btih:somelink",
			  "download-dir": "/home/transmission/download",
			  "paused": true,
			  "bandwidthPriority": 1,
			  "priority-high": [1, 4, 7],
			  "priority-normal": [2, 5, 8],
			  "priority-low": [3, 6, 9],
			  "peer-limit": 10,
			  "files-wanted": [1, 2, 3],
			  "files-unwanted": [4, 5, 6],
			  "cookies": "a=b; c=d"
		        }
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "torrent-added": {
			    "id": 1,
			    "hashString": "12345",
			    "name": "some-torrent"
			  }
		        }
		  }`)
	})

	got, err := client.AddTorrent(context.Background(), &AddTorrentReq{
		URL:                 OptString("magnet:?xt=urn:btih:somelink"),
		DownloadDirectory:   OptString("/home/transmission/download"),
		Cookies:             []Cookie{{Name: "a", Value: "b"}, {Name: "c", Value: "d"}},
		Paused:              OptBool(true),
		Priority:            OptPriority(PriorityHigh),
		HighPriorityFiles:   []int{1, 4, 7},
		NormalPriorityFiles: []int{2, 5, 8},
		LowPriorityFiles:    []int{3, 6, 9},
		PeerLimit:           OptInt(10),
		WantedFiles:         []int{1, 2, 3},
		UnwatedFiles:        []int{4, 5, 6},
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &NewTorrent{
		ID:   1,
		Hash: "12345",
		Name: "some-torrent",
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected response, diff = \n%s", cmp.Diff(want, got))
	}
}

func TestAddTorrent_contents(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-add",
			"arguments": {
			  "metainfo": "dG9ycmVudC1jb250ZW50cw=="
		        }
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "torrent-added": {
			    "id": 1,
			    "hashString": "12345",
			    "name": "some-torrent"
			  }
		        }
		  }`)
	})

	got, err := client.AddTorrent(context.Background(), &AddTorrentReq{
		Meta: strings.NewReader("torrent-contents"),
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &NewTorrent{
		ID:   1,
		Hash: "12345",
		Name: "some-torrent",
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected response, diff = \n%s", cmp.Diff(want, got))
	}
}

func TestAddTorrent_duplicate(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-add",
			"arguments": {
			  "metainfo": "dG9ycmVudC1jb250ZW50cw=="
		        }
		}`)

		fmt.Fprintf(w, `{
			"result": "success",
			"arguments": {
			  "torrent-duplicate": {
			    "id": 1,
			    "hashString": "12345",
			    "name": "some-torrent"
			  }
		        }
		  }`)
	})

	got, err := client.AddTorrent(context.Background(), &AddTorrentReq{
		Meta: strings.NewReader("torrent-contents"),
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := &NewTorrent{
		ID:   1,
		Hash: "12345",
		Name: "some-torrent",
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("unexpected response, diff = \n%s", cmp.Diff(want, got))
	}
}
