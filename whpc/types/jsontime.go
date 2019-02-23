package types

import (
	"time"
	"fmt"
)

const timeDateFormat = "2006-01-02T15:04:05Z"

type JsonTime time.Time

func (t JsonTime) Unix() int64 {
	return time.Time(t).Unix()
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeDateFormat))
	return []byte(stamp), nil
}

func (t JsonTime) String() string {
	return fmt.Sprintf("%v", t.Unix())
}

