package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

const timeFormat = "2006-01-02 15:04:05"
const dateFormat = "2006-01-02"
const timezone = "Asia/Shanghai"

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t LocalTime) Date() string {
	return time.Time(t).Format(dateFormat)
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func IsDate(date string) bool {
	_, err := time.Parse(dateFormat, date)
	if err != nil {
		return false
	}
	return true
}

func Today() string {
	return time.Now().Format(dateFormat)
}

func Now() string {
	return time.Now().Format(timeFormat)
}

func TodayTimeStamp() int64 {
	return time.Now().Unix()
}

func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(dateFormat)
}
