package transmission

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestTorrentActions(t *testing.T) {
	var tests = []struct {
		name string
		fn   func(ctx context.Context, c *Client, ids Identifier) error
		ids  Identifier
		body string
	}{
		{
			name: "start",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.StartTorrents(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"torrent-start","arguments":{"ids":1}}`,
		},
		{
			name: "start_now",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.StartTorrentsNow(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"torrent-start-now","arguments":{"ids":1}}`,
		},
		{
			name: "stop",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.StopTorrents(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"torrent-stop","arguments":{"ids":1}}`,
		},
		{
			name: "verify",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.VerifyTorrents(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"torrent-verify","arguments":{"ids":1}}`,
		},
		{
			name: "reannounce",
			fn:   func(ctx context.Context, c *Client, ids Identifier) error { return c.ReannounceTorrents(ctx, ids) },
			ids:  ID(1),
			body: `{"method":"torrent-reannounce","arguments":{"ids":1}}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			client, handle, teardown := setup(t)
			defer teardown()

			handle(func(w http.ResponseWriter, r *http.Request) {
				testBody(t, r, tc.body)

				fmt.Fprintf(w, `{"result":"success"}`)
			})

			if err := tc.fn(context.Background(), client, tc.ids); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestRenameTorrentPath(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-rename-path",
			"arguments": {
				"ids": "abcde",
				"path": "oldtorrentpath/file.iso",
				"name": "newfile.iso"
			}
		}`)

		fmt.Fprintf(w, `{
			"result":"success",
			"arguments": {
				"ids": 1,
				"path": "oldtorrentpath/file.iso",
				"name": "newfile.iso"
			}
		}`)
	})

	err := client.RenameTorrentPath(context.Background(), Hash("abcde"), "oldtorrentpath/file.iso", "newfile.iso")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSetTorrentsLocation(t *testing.T) {
	client, handle, teardown := setup(t)
	defer teardown()

	handle(func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{
			"method": "torrent-set-location",
			"arguments": {
				"ids": [1, "abcde"],
				"location": "/new/path",
				"move": true
			}
		}`)

		fmt.Fprintf(w, `{"result":"success"}`)
	})

	err := client.SetTorrentsLocation(context.Background(), IDs(ID(1), Hash("abcde")), "/new/path", true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
