package main

import (
	"time"

	"github.com/rustysys-dev/attendance-tc/internal/usecase"
	"github.com/rustysys-dev/attendance-tc/internal/utils/clock"
)

func scheduleMessage(t *time.Time, msg string) {
	if t == nil {
		return
	}
	tick := time.NewTicker(time.Until(*t))
	quit := make(chan struct{})
	defer close(quit)
	for {
		select {
		case <-tick.C:
			usecase.SendMessage(msg)
			quit <- struct{}{}
		case <-quit:
			tick.Stop()
			return
		}
	}
}

func newSchedule() {
	go scheduleMessage(clock.InTime(), "started work")
	go scheduleMessage(clock.OutTime(), "finishing work")
}

func main() {
	newSchedule()
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	_ = time.AfterFunc(time.Until(today), func() {
		tick := time.NewTicker(24 * time.Hour)
		quit := make(chan struct{})
		defer close(quit)

		for {
			select {
			case <-tick.C:
				newSchedule()
				quit <- struct{}{}
			case <-quit:
				tick.Stop()
				return
			}
		}
	})
	select {}
}
