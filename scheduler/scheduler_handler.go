package scheduler

import (
	"github.com/jasonlvhit/gocron"
)

func IntervalDo(interval int, f func()) {
	gocron.Every(uint64(interval)).Seconds().Do(f)
}

func DayDo(daily_time string, f func()) {
	gocron.Every(1).Day().At(daily_time).Do(f)
	<-gocron.Start()
}
