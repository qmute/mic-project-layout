package types

import (
	"strings"
	"time"

	"github.com/quexer/utee"
)

const (
	DateLayout = "20060102"
)

// Date 日期类型， 格式形如"20060102"
type Date string

func (p Date) Tick() utee.Tick {
	// 空日期返回0
	if p == "" {
		return 0
	}
	t, _ := p.Parse()
	return utee.NewTick(t)
}

func (p Date) Parse() (time.Time, error) {
	return time.ParseInLocation(DateLayout, string(p), time.Local)
}

func (p Date) ToTime() time.Time {
	t, err := p.Parse()
	utee.Chk(err)
	return t
}

func (p Date) Month() Month {
	t, err := p.Parse()
	utee.Chk(err)

	return MonthFromTime(t)
}

func (p Date) Valid() bool {
	_, err := time.Parse(DateLayout, string(p))
	return err == nil
}

func (p Date) AddDay(day int) Date {
	t, err := p.Parse()
	utee.Chk(err)

	return DateFromTime(t.AddDate(0, 0, day))
}

func (p Date) Weekday() time.Weekday {
	t, err := p.Parse()
	utee.Chk(err)

	return t.Weekday()
}

func (p Date) Format() string {
	if p == "" {
		return ""
	}

	t, _ := p.Parse()
	return t.Format("2006-01-02")
}

func (p Date) Equal(x Date) bool {
	t, _ := p.Parse()
	xt, _ := x.Parse()
	return t.Equal(xt)
}

func (p Date) Value() string {
	return string(p)
}

func DateFromTime(t time.Time) Date {
	return Date(t.Format(DateLayout))
}

func DateFromTick(tick utee.Tick) Date {
	if tick == 0 {
		return ""
	}
	return Date(tick.ToTime().Format(DateLayout))
}

// DateLong 日期长类型， 格式形如"2006-01-02"
type DateLong string

func (p DateLong) ToDate() Date {
	return Date(strings.Replace(string(p), "-", "", -1))
}
