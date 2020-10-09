package utilities

import "time"

// Create a Unix timestamp with milliseconds
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
