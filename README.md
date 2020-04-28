# go-transmission

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]
[![Test Status](https://github.com/pborzenkov/go-transmission/workflows/CI/badge.svg)][ci]

go-transmission is a Go client library for talking to Transmission torrent client via JSON RPC.

The library is written for RPC version 15 (Transmission >= 2.80), but should work with future Transmission versions without problems.

The API is not yet considered stable and might break without prior notice.

[godev]: https://pkg.go.dev/github.com/pborzenkov/go-transmission/transmission
[ci]: https://github.com/pborzenkov/go-transmission/actions?query=workflow%3ACI

## Usage

```go
import "github.com/pborzenkov/go-transmission/transmission"
```

Construct a new Transmission client and use it to access Transmission functions. For example:

```go
client, err := transmission.New("http://localhost:9091")

// Add new torrent from local file contents
file, err := os.Open("/local/file.torrent")
torrent, err := client.AddTorrent(context.Background(), &transmission.AddTorrentReq{
    Meta: file,
})

// Add new torrent using magnet link
torrent, err := client.AddTorrent(context.Background(), &transmission.AddTorrentReq{
    URL: <magnet link>,
})

// Get IDs and names of recently active torrents
torrents, err := client.GetTorrents(context.Background(), transmission.RecentlyActive(),
        transmission.TorrentFieldID, transmission.TorrentFieldName)
```

## License

MIT - See [LICENSE][license] file

[license]: https://github.com/pborzenkov/go-transmission/blob/master/LICENSE

