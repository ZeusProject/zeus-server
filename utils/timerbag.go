package utils

import (
	"sync"
	"time"
)

type ScheduledFn func()

type TimerBag struct {
	mutex  sync.Mutex
	timers *Timer
}

func NewTimerBag() *TimerBag {
	return &TimerBag{}
}

func (b *TimerBag) New(d time.Duration) *Timer {
	t := &Timer{
		Timer:    time.NewTimer(d),
		Duration: d,
		bag:      b,
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	t.next = b.timers
	b.timers = t

	return t
}

func (b *TimerBag) remove(t *Timer) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if t.prev != nil {
		t.prev.next = t.next
	} else {
		b.timers = t.next
	}

	if t.next != nil {
		t.next.prev = t.prev
	}
}

func (b *TimerBag) Schedule(d time.Duration, fn ScheduledFn) *Timer {
	t := b.New(d)

	go func() {
		_, ok := <-t.C

		if ok {
			fn()
		}
	}()

	return t
}

func (b *TimerBag) ScheduleRecurrent(d time.Duration, fn ScheduledFn) *Timer {
	t := b.New(d)

	go func() {
		for true {
			_, ok := <-t.C

			if ok {
				fn()
				t.Reset()
			}
		}
	}()

	return t
}

func (b *TimerBag) Close() {
	t := b.timers

	for t != nil {
		t.Stop()
		t = t.next
	}

	b.timers = nil
}
