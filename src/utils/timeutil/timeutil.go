package timeutil

import (
	"github.com/spf13/cast"
	"time"
)

// Datetime string
const Datetime = "2006-01-02 15:04:05"

// ToDateTime parse string to datetime format
func ToDateTime(str string) time.Time {
	t, err := time.Parse(Datetime+" -0700", str+" +0800")
	if err != nil {
		panic(err)
	}
	return t
}

func ToUTCTime(t time.Time, zone int) time.Time {
	t = t.UTC()
	dur, _ := time.ParseDuration("+" + cast.ToString(zone) + "h")
	return t.Add(dur)
}

func GetMonthInt(month string) int {
	switch month {
	case "January":
		{
			return 1
		}
	case "February":
		{
			return 2
		}
	case "March":
		{
			return 3
		}
	case "April":
		{
			return 4
		}
	case "May":
		{
			return 5
		}
	case "June":
		{
			return 6
		}
	case "July":
		{
			return 7
		}
	case "August":
		{
			return 8
		}
	case "September":
		{
			return 9
		}
	case "October":
		{
			return 10
		}
	case "November":
		{
			return 11
		}
	case "December":
		{
			return 12
		}
	}
	return 0
}
