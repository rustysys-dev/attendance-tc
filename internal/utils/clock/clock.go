package clock

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rustysys-dev/attendance-tc/internal/utils/config"
)

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomTime(from, until int) *time.Time {
	now := time.Now()
	// skip saturday and sunday
	if now.Weekday() == 0 || now.Weekday() == 6 {
		return nil
	}
	if now.Hour() >= until-1 {
		return nil
	}

	hour := randomRange(from, until)
	minute := randomRange(0, 59)

	rtn := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
	return &rtn
}

func InTime() *time.Time {
	t := randomTime(config.StartMin(), config.StartMax())
	if t != nil {
		_, _ = fmt.Println("scheduling clockin at:", t.String())
	}
	return t
}

func OutTime() *time.Time {
	t := randomTime(config.EndMin(), config.EndMax())
	if t != nil {
		_, _ = fmt.Println("scheduling clockout at:", t.String())
	}
	return t
}
