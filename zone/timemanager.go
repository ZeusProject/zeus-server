package zone

import (
	"time"
)

type TimeManager struct {
	start time.Time
}

func NewTimeManager() *TimeManager {
	return &TimeManager{
		start: time.Now(),
	}
}

func (t *TimeManager) GetTick() uint32 {
	return uint32(time.Now().Sub(t.start).Seconds())
}
