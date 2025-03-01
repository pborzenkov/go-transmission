package transmission

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Weekday specifies a day of the week.
type Weekday int

// Constants for each day of week
const (
	Sunday Weekday = (1 << iota)
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	weekday  = Monday | Tuesday | Wednesday | Thursday | Friday
	weekend  = Saturday | Sunday
	everyday = weekday | weekend
)

var (
	weekdayNames = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
)

// String returns an English name of the day taking into account special day
// combination, like weekday or weekend.
func (d Weekday) String() string {
	switch d {
	case everyday:
		return "Every Day"
	case weekend:
		return "Weekends"
	case weekday:
		return "Weekdays"
	default:
		days := make([]string, 0)
		for i := 0; i < len(weekdayNames); i++ {
			if d&(1<<i) != 0 {
				days = append(days, weekdayNames[i])
			}
		}
		return strings.Join(days, ", ")
	}
}

// Encryption specifies encryption configuration
type Encryption int

const (
	// EncryptionRequired means Transmission requires encrypted connections for this torrent
	EncryptionRequired Encryption = iota
	// EncryptionPreferred means Transmission prefers encrypted connections for this torrent
	EncryptionPreferred
	// EncryptionTolerated means Transmission can use encrypted connections for this torrent
	EncryptionTolerated
)

var (
	encryptionNames = [...]string{"required", "preferred", "tolerated"}
)

// String returns string representation of the encryption configuration.
func (e Encryption) String() string {
	if int(e) >= len(encryptionNames) {
		return fmt.Sprintf("Encryption(%d)", e)
	}
	return encryptionNames[e]
}

// MarshalJSON implements custom JSON marshaler for encryption configuration.
func (e Encryption) MarshalJSON() ([]byte, error) {
	if int(e) >= len(encryptionNames) {
		return nil, fmt.Errorf("unsupported Encryption value %d", e)
	}
	return []byte(`"` + encryptionNames[e] + `"`), nil
}

// UnmarshalJSON implements custom JSON unmarshaler for encryption
// configuration.
func (e *Encryption) UnmarshalJSON(data []byte) error {
	var enc string
	if err := json.Unmarshal(data, &enc); err != nil {
		return err
	}
	for i := 0; i < len(encryptionNames); i++ {
		if encryptionNames[i] == enc {
			*e = Encryption(i)
			return nil
		}
	}

	return fmt.Errorf("unsupported Encryption value %q", enc)
}

// Priority indicates torrent or file priority.
type Priority int

const (
	// PriorityLow indicates low priority
	PriorityLow Priority = -1
	// PriorityNormal indicates normal priority
	PriorityNormal Priority = 0
	// PriorityHigh indicates high priority
	PriorityHigh Priority = 1
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityNormal:
		return "normal"
	case PriorityHigh:
		return "high"
	default:
		return fmt.Sprintf("Priority(%d)", p)
	}
}

// Limit controls whether a particular torrent follows global limits or not.
type Limit int

const (
	// LimitGlobal configures torrent to honor global limit
	LimitGlobal Limit = 0
	// LimitLocal configures torent to honor local torrent limit
	LimitLocal Limit = 1
	// LimitUnlimited configures torrent to not honor any limit
	LimitUnlimited Limit = 2
)

func (l Limit) String() string {
	switch l {
	case LimitGlobal:
		return "global"
	case LimitLocal:
		return "normal"
	case LimitUnlimited:
		return "unlimited"
	default:
		return fmt.Sprintf("Limit(%d)", l)
	}
}

// Status indicates torrent status
type Status int

const (
	// StatusStopped indicates that torrent is stopped
	StatusStopped Status = 0
	// StatusCheckWait indicates that torrent is queued for checking
	StatusCheckWait Status = 1
	// StatusCheck indicates that torrent is being checked
	StatusCheck Status = 2
	// StatusDownloadWait indicates that torrent is queued for downloading
	StatusDownloadWait Status = 3
	// StatusDownload indicates that torrent is being downloaded
	StatusDownload Status = 4
	// StatusSeedWait indicates that torrent is queued for seeding
	StatusSeedWait Status = 5
	// StatusSeed indicates that torrent is being seeded
	StatusSeed Status = 6
)

func (s Status) String() string {
	switch s {
	case StatusStopped:
		return "stopped"
	case StatusCheckWait:
		return "queued for checking"
	case StatusCheck:
		return "checking"
	case StatusDownloadWait:
		return "queued for downloading"
	case StatusDownload:
		return "downloading"
	case StatusSeedWait:
		return "queued for seeding"
	case StatusSeed:
		return "seeding"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}

// ErrorType defines a category of torrent error.
type ErrorType int

const (
	// ErrorTypeOK means no error
	ErrorTypeOK ErrorType = 0
	// ErrorTypeTrackerWarning indicates a warning from tracker
	ErrorTypeTrackerWarning ErrorType = 1
	// ErrorTypeTrackerError indicates an error from tracker
	ErrorTypeTrackerError ErrorType = 2
	// ErrorTypeLocalError indicates a local problem
	ErrorTypeLocalError ErrorType = 3
)

func (e ErrorType) String() string {
	switch e {
	case ErrorTypeOK:
		return "OK"
	case ErrorTypeTrackerWarning:
		return "tracker warning"
	case ErrorTypeTrackerError:
		return "tracker error"
	case ErrorTypeLocalError:
		return "local error"
	default:
		return fmt.Sprintf("ErrorType(%d)", e)
	}
}

// TrackerState defines a state of a tracker.
type TrackerState int

const (
	// TrackerStateInactive indicates that Transmission is not announcing to tracker
	TrackerStateInactive TrackerState = 0
	// TrackerStateWaiting indicates that Transmission is waiting to announce to tracker
	TrackerStateWaiting TrackerState = 1
	// TrackerStateQueued indicates that Transmission has queued the announce to tracker
	TrackerStateQueued TrackerState = 2
	// TrackerStateActive indicates that Transmission has announced to tracker
	TrackerStateActive TrackerState = 3
)

func (t TrackerState) String() string {
	switch t {
	case TrackerStateInactive:
		return "inactive"
	case TrackerStateWaiting:
		return "waiting"
	case TrackerStateQueued:
		return "queued"
	case TrackerStateActive:
		return "active"
	default:
		return fmt.Sprintf("TrackerState(%d)", t)
	}
}

// boolint handles both JSON bools and ints
type boolint bool

func (b *boolint) UnmarshalJSON(data []byte) error {
	var tb bool
	if err := json.Unmarshal(data, &tb); err == nil {
		*b = boolint(tb)
		return nil
	}
	var ti int
	if err := json.Unmarshal(data, &ti); err != nil {
		return err
	}
	*b = boolint(ti != 0)
	return nil
}

// Cookie is an HTTP cookie
type Cookie struct {
	Name  string
	Value string
}

func (c *Cookie) String() string {
	return c.Name + "=" + c.Value
}
