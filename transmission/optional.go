package transmission

import (
	"time"
)

// OptInt is a helper routine that allocates new int to store v and returns a
// pointer to it.
func OptInt(v int) *int {
	return &v
}

// OptInt64 is a helper routine that allocates new int64 to store v and returns
// a pointer to it.
func OptInt64(v int64) *int64 {
	return &v
}

// OptFloat64 is a helper routine that allocates new float64 to store v and
// returns a pointer to it.
func OptFloat64(v float64) *float64 {
	return &v
}

// OptBool is a helper routine that allocates new bool to store v and return a
// pointer to it.
func OptBool(v bool) *bool {
	return &v
}

// OptString is a helper routine that allocates new string to store v and
// returns a pointer to it.
func OptString(v string) *string {
	return &v
}

// OptDuration is a helper routine that allocates new time.Duration to store v
// and return a pointer to it.
func OptDuration(v time.Duration) *time.Duration {
	return &v
}

// OptWeekday is a helper routine that allocates new Weekday to store v and
// returns a pointer to it.
func OptWeekday(v Weekday) *Weekday {
	return &v
}

// OptEncryption is a helper routine that allocates new Encryption to store v
// and returns a pointer to it.
func OptEncryption(v Encryption) *Encryption {
	return &v
}

// OptPriority is a helper routine that allocates new Priority to store v and
// returns a pointer to it.
func OptPriority(v Priority) *Priority {
	return &v
}

// OptLimit is a helper routine that allocates new Limit to store v and returns
// a pointer to it.
func OptLimit(v Limit) *Limit {
	return &v
}
