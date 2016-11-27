package utils

import "time"

type Timer struct {
	*time.Timer
	Duration time.Duration
}

func (t *Timer) Reset() {
	t.Timer.Stop()
	t.Timer.Reset(t.Duration)
}
