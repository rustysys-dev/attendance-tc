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
	// schedule messages for right away
	newSchedule()
	// find beginning of tomorrow
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	untilTomorrow := time.Until(tomorrow)
	_, _ = fmt.Println("Rescheduling new messages in", untilTomorrow)
	// schedule main sub-routine to run tomorrow at 0000
	_ = time.AfterFunc(untilTomorrow, func() {
		// schedule new messages
		newSchedule()
		_, _ = fmt.Println("Rescheduling new messages in", mainCycleRefreshTime)
		// make ticket to schedule messages every set period of time (in this case 24h)
		tick := time.NewTicker(mainCycleRefreshTime)
		quit := make(chan struct{})
		defer close(quit)
		// start main sub-routine main loop
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
