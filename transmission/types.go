package transmission

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Weeekday specifies a day of the week.
type Weekday int

const (
	Sunday Weekday = (1 << iota)
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

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
	EncryptionRequired Encryption = iota
	EncryptionPreferred
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
	PriorityLow    Priority = -1
	PriorityNormal Priority = 0
	PriorityHigh   Priority = 1
)

// Limit controls whether a particular torrent follows global limits or not.
type Limit int

const (
	LimitGlobal    Limit = 0 // Honor global limit
	LimitLocal     Limit = 1 // Honor local torrent limit
	LimitUnlimited Limit = 2 // Don't honor any limit
)
