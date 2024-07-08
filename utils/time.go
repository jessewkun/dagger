package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

const TIME_FORMAT = "2006-01-02 15:04:05"
const DATE_FORMAT = "2006-01-02"
const TIMEZONE = "Asia/Shanghai"

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TIME_FORMAT)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TIME_FORMAT)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TIME_FORMAT+`"`, string(data), time.Local)
	*t = LocalTime(now)
	return
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TIME_FORMAT)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(TIMEZONE)
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
	return time.Time(t).Format(DATE_FORMAT)
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
	_, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return false
	}
	return true
}

func Today() string {
	return time.Now().Format(DATE_FORMAT)
}

func Now() string {
	return time.Now().Format(TIME_FORMAT)
}

func TodayTimeStamp() int64 {
	return time.Now().Unix()
}

func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DATE_FORMAT)
}
