//go:generate go run ../tools/gen-fields.go -type Torrent
package transmission

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"net/url"
	"time"
)

// Identifier can identify one or multiple torrents.
type Identifier interface {
	canID()
}

// SingularIdentifier can identify exactly one torrent.
type SingularIdentifier interface {
	canIDOne()
}

// ID identifies torrents. It implements both Identifier and
// SingularIdentifier.
type ID int

var _ Identifier = ID(0)
var _ SingularIdentifier = ID(0)

func (ID) canID()    {}
func (ID) canIDOne() {}

// Hash identifier torrents by hash. It implements both Identifier and
// SingularIdentifier.
type Hash string

var _ Identifier = Hash("")
var _ SingularIdentifier = Hash("")

func (Hash) canID()    {}
func (Hash) canIDOne() {}

type idstring string

var _ Identifier = idstring("")

func (idstring) canID() {}

// RecentlyActive identifies torrents that have been active in the last hour.
func RecentlyActive() Identifier {
	return idstring("recently-active")
}

// All identifies all torrents.
func All() Identifier {
	return nil
}

// IDList is a list of torrent IDs
type IDList []SingularIdentifier

var _ Identifier = IDList([]SingularIdentifier{})

func (IDList) canID() {}

// IDs returns an identifier that identifies a list of provided torrents.
func IDs(ids ...SingularIdentifier) Identifier {
	return IDList(ids)
}

// Torrent describes a torrent file.
type Torrent struct {
	// ID of the torrent
	ID ID `json:"id"`
	// Hash of the torrent
	Hash Hash `json:"hashString"`
	// Name of the torrent
	Name string `json:"name"`
	// Current torrent status
	Status Status `json:"status"`
	// Torrent creator
	Creator string `json:"creator"`
	// An optional commen
	Comment string `json:"comment"`
	// Torrent labels
	Labels []string `json:"labels"`

	// ETA until the torrent is done. This is 0 if ETA is unknown or not
	// available
	ETA time.Duration `json:"eta"`
	// ETA until the idle time limit is reached if seeding
	IdleETA time.Duration `json:"etaIdle"`

	// Type of error
	ErrorType ErrorType `json:"error"`
	// Error message
	Error string `json:"errorString"`

	// Path to torrent file
	File string `json:"torrentFile"`
	// Torrent magnet link
	MagnetLink string `json:"magnetLink"`
	// Torrent download directory
	DownloadDirectory string `json:"downloadDir"`

	// Torrent creation date
	CreatedAt time.Time `json:"-" field:"dateCreated"`
	// Date when torrent was added
	AddedAt time.Time `json:"-" field:"addedDate"`
	// Date when torrent was created
	StartedAt time.Time `json:"-" field:"startDate"`
	// Date when torrent was last active
	LastActiveAt time.Time `json:"-" field:"activityDate"`
	// Date when torrent was completed
	DoneAt time.Time `json:"-" field:"doneDate"`
	// Time when one of the trackers will allow to manually ask for more
	// peers
	CanManuallyAnnounceAt time.Time `json:"-" field:"manualAnnounceTime"`

	// Current download rate
	DownloadRate int64 `json:"rateDownload"`
	// Current upload rate
	UploadRate int64 `json:"rateUpload"`
	// Download rate limit
	DownloadRateLimit int64 `json:"downloadLimit"`
	// Indicates if download rate is limited
	DownloadRateLimitEnabled bool `json:"downloadLimited"`
	// Upload rate limit
	UploadRateLimit int64 `json:"uploadLimit"`
	// Indicates if upload rate is limited
	UploadRateLimited bool `json:"uploadLimited"`
	// Idicates if session limits are honored for this torrent
	HonorSessionLimits bool `json:"honorsSessionLimits"`

	// Total amount of data downloaded for this torrent
	DownloadedTotal int64 `json:"downloadedEver"`
	// Total amount of data uploaded for this torrent
	UploadedTotal int64 `json:"uploadedEver"`
	// Total amount of corrupted data downloaded for this torrent
	CorruptedTotal int64 `json:"corruptEver"`

	// Bandwidth priority
	Priority Priority `json:"bandwidthPriority"`
	// Position in queue
	PositionInQueue int `json:"queuePosition"`

	// Stop torrent after given time of inactivity
	IdleSeedingLimit time.Duration `json:"seedIdleLimit"`
	// Which IdleSeedingLimit value to use
	IdleSeedingLimitMode Limit `json:"seedIdleMode"`
	// Stop seeding after reaching the given ratio
	UploadRatioLimit float64 `json:"seedRatioLimit"`
	// Which UploadRatioLimit value to use
	UploadRatioLimitMode Limit `json:"seedRatioMode"`
	// Current upload ration
	UploadRatio float64 `json:"uploadRatio"`

	// Time spent downloading this torrent
	DownloadingFor time.Duration `json:"secondsDownloading"`
	// Time spent uploading this torrent
	SeedingFor time.Duration `json:"secondsSeeding"`

	// Total size of all files in the torrent
	TotalSize int64 `json:"totalSize"`
	// Size of wanted files in the torrent
	WantedSize int64 `json:"sizeWhenDone"`
	// Size of wanted files that is available for download from peers
	WantedAvailable int64 `json:"desiredAvailable"`
	// Size of yet to download wanted files
	WantedLeft int64 `json:"leftUntilDone"`
	// Total amount of downloaded but not yet checked data
	UncheckedSize int64 `json:"haveUnchecked"`
	// Total amount of downloaded and checked data
	ValidSize int64 `json:"haveValid"`

	// Percentage of data completed
	DataDone float64 `json:"percentDone"`
	// Percentage of data checked
	DataChecked float64 `json:"recheckProgress"`
	// Percentage of metadata completed
	MetadataDone float64 `json:"metadataPercentComplete"`

	// Is torrent finished (met its seeding ration)
	IsFinished bool `json:"isFinished"`
	// Is torrent private
	IsPrivate bool `json:"isPrivate"`
	// Is torrent stalled
	IsStalled bool `json:"isStalled"`

	// Maximum allowed number of peers
	PeerLimit int `json:"peer-limit"`
	// Current number of peers
	ConnectedPeers int `json:"peersConnected"`
	// Number of peers getting data from us
	PeersGettingFromUs int `json:"peersGettingFromUs"`
	// Number of peers sending data to us
	PeersSendingToUs int `json:"peersSendingToUs"`
	// Array of peer descriptors
	Peers []Peer `json:"peers"`
	// Origin of the peers
	PeersFrom PeersOrigin `json:"peersFrom"`
	// Number of web seeds seeding to us
	WebSeedsSendingToUs int `json:"webseedsSendingToUs"`
	// Array of web seeds
	WebSeeds []string `json:"webseeds"`

	// An array the size of number of files with boolean flag indicating
	// whether we want the file or not
	Wanted []bool `json:"-" field:"wanted"`
	// Array of files in the torrent
	Files []File `json:"files"`
	// File statistics
	FileStats []FileStat `json:"fileStats"`
	// An array of file priorities
	Priorities []Priority `json:"priorities"`

	// Number of pieces
	PieceCount int64 `json:"pieceCount"`
	// Size of each piece
	PieceSize int64 `json:"pieceSize"`
	// Pieces holds info about downloaded torren pieces
	Pieces Pieces `json:"pieces"`

	// Trackers holds the list of torrent trackers
	Trackers []Tracker `json:"-" field:"trackers"`
	// TrackerStats holds statistics about trackers
	TrackerStats []TrackerStat `json:"-" field:"trackerStats"`
}

// Peer identifies a single peer
type Peer struct {
	// Address of the peer
	Address net.IP `json:"address"`
	// Connection port
	Port int `json:"port"`
	// Name of the torrent client
	ClientName string `json:"clientName"`

	// Rate of downloading data from the peer
	DownloadRate int64 `json:"rateToClient"`
	// Rate of uploading data to the peer
	UploadRate int64 `json:"rateToPeer"`

	// Percentage of data peer has available
	Progress float64 `json:"progress"`

	// Indicates whether we are downloading from the peer
	IsDownloading bool `json:"isDownloadingFrom"`
	// Indicates whether we are uploading to the peer
	IsUploading bool `json:"isUploadingTo"`
	// Indicates whether we are connected via µTP
	IsUTP bool `json:"isUTP"`
	// Indicates whether peer is incoming or not
	IsIncoming bool `json:"isIncoming"`
	// Indicates if connection with the peer is encrypted
	IsEncrypted bool `json:"isEncrypted"`
	// Indicates whether Transmission is choked received from the client
	AreWeChoked bool `json:"clientIsChoked"`
	// Indicates whether peer is choked receiving from us
	IsPeerChoked bool `json:"peerIsChoked"`
	// Indicates whether we are interested in receiving data from the peer
	AreWeInterested bool `json:"clientIsInterested"`
	// Indicates whether peer is interested in receiveing data form us
	IsPeerInterested bool `json:"peerIsInterested"`
}

// PeersOrigin holds information about origin of the peers.
type PeersOrigin struct {
	// Number of peers from tracker
	Tracker int `json:"fromTracker"`
	// Number of incoming peers
	Incoming int `json:"fromIncoming"`
	// Number of peers from cache
	Cache int `json:"fromCache"`
	// Number of peers from DHT
	DHT int `json:"fromDht"`
	// Number of peers from local peer discovery
	LPD int `json:"fromLpd"`
	// Number of peers from peer exchange
	PEX int `json:"fromPex"`
	// Number of peers from LTEP handshake
	LTEP int `json:"fromLtep"`
}

// File describes a single file within a torrent.
type File struct {
	// Name of the file
	Name string `json:"name"`
	// Size of the file
	Size int64 `json:"length"`
	// The amount of downloaded data
	Downloaded int64 `json:"bytesCompleted"`
}

// FileStat holds statistics about single file within a torrent.
type FileStat struct {
	// The amount of downloaded data
	Downloaded int64 `json:"bytesCompleted"`
	// File priority
	Priority Priority `json:"priority"`
	// Indicates whether we want the file or not
	Wanted bool `json:"wanted"`
}

// Pieces holds info about downloaded pices
type Pieces []byte

// IsDownloaded returns true if the given piece is downloaded. It doesn't
// return an error if there is no such piece, but rather just returns false.
func (p Pieces) IsDownloaded(piece int) bool {
	if piece < 0 || piece/8 >= len(p) {
		return false
	}

	return ((p[piece>>3] << (piece & 0x7)) & 0x80) != 0
}

// UnmarshalJSON unmarshals pieces data from JSON
func (p *Pieces) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*p, err = base64.StdEncoding.DecodeString(str)
	return err
}

// Tracker describes a single tracker
type Tracker struct {
	ID          int      `json:"id"`
	Tier        int      `json:"tier"`
	AnnounceURL *url.URL `json:"-"`
	ScrapeURL   *url.URL `json:"-"`
}

// TrackerStat holds stats for a single tracker.
type TrackerStat struct {
	ID          int      `json:"id"`
	Tier        int      `json:"tier"`
	IsBackup    bool     `json:"isBackup"`
	Host        *url.URL `json:"-"`
	AnnounceURL *url.URL `json:"-"`
	ScrapeURL   *url.URL `json:"-"`

	Leechers  int `json:"leecherCount"`
	Seeders   int `json:"seederCount"`
	Downloads int `json:"downloadCount"`

	HasAnnounced            bool         `json:"hasAnnounced"`
	AnnounceState           TrackerState `json:"announceState"`
	LastAnnounceStartTime   time.Time    `json:"-"`
	LastAnnounceTime        time.Time    `json:"-"`
	IsLastAnnounceTimedOut  bool         `json:"lastAnnounceTimedOut"`
	IsLastAnnounceSucceeded bool         `json:"lastAnnounceSucceeded"`
	LastAnnounceResult      string       `json:"lastAnnounceResult"`
	LastAnnouncePeerCount   int          `json:"lastAnnouncePeerCount"`
	NextAnnounceTime        time.Time    `json:"-"`

	HasScraped            bool         `json:"hasScraped"`
	ScrapeState           TrackerState `json:"scrapeState"`
	LastScrapeStartTime   time.Time    `json:"-"`
	LastScrapeTime        time.Time    `json:"-"`
	IsLastScrapeTimedOut  bool         `json:"-"`
	IsLastScrapeSucceeded bool         `json:"lastScrapeSucceeded"`
	LastScrapeResult      string       `json:"lastScrapeResult"`
	NextScrapeTime        time.Time    `json:"-"`
}

func duration(a, b time.Duration) time.Duration {
	if a > 0 {
		return a * b
	}
	return a
}

func unixtime(t int64) time.Time {
	if t > 0 {
		return time.Unix(t, 0)
	}
	return time.Time{}
}

type trackerJSON struct {
	Tracker
	AnnounceURL string `json:"announce"`
	ScrapeURL   string `json:"scrape"`
}

func (tj *trackerJSON) tracker() (Tracker, error) {
	var err error

	if tj.Tracker.AnnounceURL, err = url.Parse(tj.AnnounceURL); err != nil {
		return Tracker{}, err
	}
	if tj.Tracker.ScrapeURL, err = url.Parse(tj.ScrapeURL); err != nil {
		return Tracker{}, err
	}

	return tj.Tracker, nil
}

type trackerStatJSON struct {
	TrackerStat
	Host                  string  `json:"host"`
	AnnounceURL           string  `json:"announce"`
	ScrapeURL             string  `json:"scrape"`
	LastAnnounceStartTime int64   `json:"lastAnnounceStartTime"`
	LastAnnounceTime      int64   `json:"lastAnnounceTime"`
	NextAnnounceTime      int64   `json:"nextAnnounceTime"`
	LastScrapeStartTime   int64   `json:"lastScrapeStartTime"`
	LastScrapeTime        int64   `json:"lastScrapeTime"`
	IsLastScrapeTimedOut  boolint `json:"lastScrapeTimedOut"`
	NextScrapeTime        int64   `json:"nextScrapeTime"`
}

func (tj *trackerStatJSON) trackerStat() (TrackerStat, error) {
	var t = &tj.TrackerStat
	var err error

	if t.Host, err = url.Parse(tj.Host); err != nil {
		return tj.TrackerStat, err
	}
	if t.AnnounceURL, err = url.Parse(tj.AnnounceURL); err != nil {
		return tj.TrackerStat, err
	}
	if t.ScrapeURL, err = url.Parse(tj.ScrapeURL); err != nil {
		return tj.TrackerStat, err
	}
	t.LastAnnounceStartTime = unixtime(tj.LastAnnounceStartTime)
	t.LastAnnounceTime = unixtime(tj.LastAnnounceTime)
	t.NextAnnounceTime = unixtime(tj.NextAnnounceTime)
	t.LastScrapeStartTime = unixtime(tj.LastScrapeStartTime)
	t.LastScrapeTime = unixtime(tj.LastScrapeTime)
	t.IsLastScrapeTimedOut = bool(tj.IsLastScrapeTimedOut)
	t.NextScrapeTime = unixtime(tj.NextScrapeTime)

	return tj.TrackerStat, err
}

type torrentJSON struct {
	*Torrent
	CreatedAt             int64             `json:"dateCreated"`
	AddedAt               int64             `json:"addedDate"`
	StartedAt             int64             `json:"startDate"`
	LastActiveAt          int64             `json:"activityDate"`
	DoneAt                int64             `json:"doneDate"`
	CanManuallyAnnounceAt int64             `json:"manualAnnounceTime"`
	Wanted                []int             `json:"wanted"`
	Trackers              []trackerJSON     `json:"trackers"`
	TrackerStats          []trackerStatJSON `json:"trackerStats"`
}

func (tj *torrentJSON) torrent(uc unitConversion) (*Torrent, error) {
	t := tj.Torrent

	t.ETA = duration(t.ETA, time.Second)
	t.IdleETA = duration(t.IdleETA, time.Second)
	t.CreatedAt = unixtime(tj.CreatedAt)
	t.AddedAt = unixtime(tj.AddedAt)
	t.StartedAt = unixtime(tj.StartedAt)
	t.LastActiveAt = unixtime(tj.LastActiveAt)
	t.DoneAt = unixtime(tj.DoneAt)
	t.CanManuallyAnnounceAt = unixtime(tj.CanManuallyAnnounceAt)
	t.DownloadRateLimit *= uc.speed
	t.UploadRateLimit *= uc.speed
	t.IdleSeedingLimit = duration(t.IdleSeedingLimit, time.Minute)
	t.DownloadingFor = duration(t.DownloadingFor, time.Second)
	t.SeedingFor = duration(t.SeedingFor, time.Second)
	if len(tj.Wanted) > 0 {
		t.Wanted = make([]bool, len(tj.Wanted))
		for i, v := range tj.Wanted {
			t.Wanted[i] = v > 0
		}
	}
	var err error
	if len(tj.Trackers) > 0 {
		t.Trackers = make([]Tracker, len(tj.Trackers))
		for i := range tj.Trackers {
			if t.Trackers[i], err = tj.Trackers[i].tracker(); err != nil {
				return nil, err
			}
		}
	}
	if len(tj.TrackerStats) > 0 {
		t.TrackerStats = make([]TrackerStat, len(tj.TrackerStats))
		for i := range tj.TrackerStats {
			if t.TrackerStats[i], err = tj.TrackerStats[i].trackerStat(); err != nil {
				return nil, err
			}
		}
	}

	return tj.Torrent, nil
}
