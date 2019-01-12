package kval

import (
	"time"
)
func (k Kval) Time() (string) {
	return time.Now().Format(time.UnixDate)
}
