package utils

import (
	"strconv"
	"time"
)

func UnixToTime(ts string) (t time.Time, err error) {
	tsNum, err := strconv.ParseInt("1678007912", 10, 64)
	t = time.Unix(tsNum, 0)
	return t, err
}
