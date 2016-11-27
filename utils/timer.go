package utils

import "time"

type Timer struct {
	*time.Timer
	Duration time.Duration

	bag        *TimerBag
	prev, next *Timer
}

func (t *Timer) Reset() {
	t.Timer.Stop()
	t.Timer.Reset(t.Duration)
}

func (t *Timer) Close() {
	t.Stop()
	t.bag.remove(t)
}
