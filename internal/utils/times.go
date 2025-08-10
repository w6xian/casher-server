package utils

import "time"

func UnixTime() int64 {
	t := time.Now()
	return t.Unix()
}

func Day(t time.Time) (int, int) {
	year, month, day := t.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	end := time.Date(year, month, day, 23, 59, 59, 999, t.Location())
	return int(start.Unix()), int(end.Unix())
}

func Month(t time.Time) (int, int) {
	year, month, _ := t.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	if month == 12 {
		month = 1
		year = year + 1
	}
	end := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	return int(start.Unix()), int(end.Unix()) - 1
}

func UnixYear(ts int64) (int, int) {
	t := time.Unix(int64(ts), 0)
	year, _, _ := t.Date()
	start := time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
	end := time.Date(year+1, 1, 1, 0, 0, 0, 0, t.Location())
	return int(start.Unix()), int(end.Unix()) - 1
}

func UnixYearStart(t int64) int {
	ti := time.Unix(t, 0)
	year, _, _ := ti.Date()
	start := time.Date(year, 1, 1, 0, 0, 0, 0, ti.Location())
	return int(start.Unix())
}

func UnixYearEnd(t int64) int {
	ti := time.Unix(t, 0)
	year, _, _ := ti.Date()
	end := time.Date(year, 12, 31, 23, 59, 59, 0, ti.Location())
	return int(end.Unix())
}
