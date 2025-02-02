package types

import (
	"time"

	"github.com/quexer/utee"
)

const (
	MonthLayout  = "200601"
	patternMonth = `^\d{6}$`
)

func MonthFromTime(t time.Time) Month {
	return Month(t.Format(MonthLayout))
}

func MonthFromTick(tick utee.Tick) Month {
	if tick == 0 {
		return ""
	}
	return Month(tick.ToTime().Format(MonthLayout))
}

type Month string

func (p Month) ToTime() time.Time {
	t, err := time.ParseInLocation("200601", string(p), time.Local)
	utee.Chk(err)
	return t
}

func (p Month) Tick() utee.Tick {
	return utee.NewTick(p.ToTime())
}

// Days 当月总天数
func (p Month) Days() int {
	current := p.ToTime()
	next := current.AddDate(0, 1, 0)
	return next.AddDate(0, 0, -1).Day()
}
