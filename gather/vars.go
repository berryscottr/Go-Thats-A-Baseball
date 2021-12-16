package gather

import "time"

type Today struct {
	day   int
	month int
	year  int
}

var (
	scheduleUrlTemplate = "https://www.baseball-reference.com/leagues/majors/{year}-schedule.shtml"
	today               = Today{
		day:   time.Now().Day(),
		month: int(time.Now().Month()),
		year:  time.Now().Year(),
	}
)
