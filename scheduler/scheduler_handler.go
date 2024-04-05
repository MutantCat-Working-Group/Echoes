package scheduler

import (
	"github.com/antlabs/timer"
	"github.com/jasonlvhit/gocron"
	"time"
)

func IntervalDo(interval int, f func()) {
	tm := timer.NewTimer()

	tm.ScheduleFunc(time.Duration(interval)*time.Second, f)

	go tm.Run()
}

func DayDo(daily_time string, f func()) {
	gocron.Every(1).Day().At(daily_time).Do(f)
	<-gocron.Start()
}
