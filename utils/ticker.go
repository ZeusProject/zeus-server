package utils

import "time"

type Ticker struct {
	*time.Ticker
	Duration time.Duration
}
