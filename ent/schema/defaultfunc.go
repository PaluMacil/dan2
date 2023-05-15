package schema

import "time"

func NowYear() int {
	return time.Now().Year()
}

func NowMonth() int {
	return int(time.Now().Month())
}

func NowDay() int {
	return time.Now().Day()
}
