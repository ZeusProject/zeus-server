package utils

import "time"

type ScheduledFn func()

type TimerBag struct {
	close chan bool
}

func NewTimerBag() *TimerBag {
	return &TimerBag{
		close: make(chan bool),
	}
}

func (t *TimerBag) Schedule(d time.Duration, fn ScheduledFn) *Timer {
	timer := &Timer{
		Timer:    time.NewTimer(d),
		Duration: d,
	}

	go func() {
		select {
		case _, ok := <-timer.C:
			if ok {
				fn()
			}
		case _, ok := <-t.close:
			if !ok {
				timer.Stop()
			}
		}
	}()

	return timer
}

func (t *TimerBag) ScheduleRecurrent(d time.Duration, fn ScheduledFn) *Ticker {
	ticker := &Ticker{
		Ticker:   time.NewTicker(d),
		Duration: d,
	}

	go func() {
		for true {
			select {
			case _, ok := <-ticker.C:
				if !ok {
					break
				}

				fn()
			case _, ok := <-t.close:
				if !ok {
					ticker.Stop()
					break
				}
			}
		}
	}()

	return ticker
}

func (t *TimerBag) Close() {
	close(t.close)
}
