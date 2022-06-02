package main

import (
	"fmt"
	"time"

	"github.com/rustysys-dev/attendance-tc/internal/usecase"
	"github.com/rustysys-dev/attendance-tc/internal/utils/clock"
)

const (
	mainCycleRefreshTime = 24 * time.Hour
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
	until := time.Until(today)
	_, _ = fmt.Println("Rescheduling new messages in", until)
	_ = time.AfterFunc(time.Until(today), func() {
		newSchedule()
		_, _ = fmt.Println("Rescheduling new messages in", mainCycleRefreshTime)
		tick := time.NewTicker(mainCycleRefreshTime)
		quit := make(chan struct{})
		defer close(quit)

		for {
			select {
			case <-tick.C:
				newSchedule()
				_, _ = fmt.Println("Rescheduling new messages in", mainCycleRefreshTime)
			case <-quit:
				tick.Stop()
				return
			}
		}
	})
	select {}
}
