package lla

import (
	"database/sql/driver"
	"strings"
	"time"
)

const (
	textDate     = "2006-01-02"
	textTime     = "15:04:05"
	textDateTime = "2006-01-02 15:04:05"
	textFullTime = "2006-01-02 15:04:05.999999 -0700"
	jsonDate     = "\"2006-01-02\""
	jsonTime     = "\"15:04:05\""
	jsonDateTime = "\"2006-01-02 15:04:05\""
	jsonFullTime = "\"2006-01-02 15:04:05.999999 -0700\""
	jsonNull     = "null"
)

///////////////////////////////////////////////////////////////////////
//
type DateOnly struct {
	time.Time
}

func ParseDateOnly(str string) (DateOnly, error) {
	tt, err := time.Parse(textDate, str)
	out := DateOnly{tt}
	return out, err
}

func (t DateOnly) String() string {
	return t.Format(textDate)
}

func (t *DateOnly) Today() {
	tm := time.Now()
	*t = DateOnly{time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())}
}

// MarshalJSON implements the json.Marshaler interface.
func (t DateOnly) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(jsonNull), nil
	}
	return []byte(t.Format(jsonDate)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *DateOnly) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		*t = DateOnly{}
		return nil
	}
	if strings.Compare(str, jsonNull) == 0 {
		*t = DateOnly{}
		return nil
	}
	tt, err := time.Parse(jsonDate, str)
	*t = DateOnly{tt}
	return err
}

// Value - Implementation of valuer for database/sql
func (t DateOnly) Value() (driver.Value, error) {
	return time.Time(t.Time), nil
}

// Scan - Implement the database/sql scanner interface
func (t *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*t = DateOnly{}
		return nil
	}
	switch vt := value.(type) {
	case time.Time:
		*t = DateOnly{vt}
	case []byte:
		tt, err := time.Parse(textDate, string(vt))
		if err != nil {
			return err
		}
		*t = DateOnly{tt}
	case string:
		tt, err := time.Parse(textDate, vt)
		if err != nil {
			return err
		}
		*t = DateOnly{tt}
	default:
		*t = DateOnly{}
	}
	return nil
}

///////////////////////////////////////////////////////////////////////
//
type TimeOnly struct {
	time.Time
}

func ParseTimeOnly(str string) (TimeOnly, error) {
	tt, err := time.Parse(textTime, str)
	out := TimeOnly{tt}
	return out, err
}

func (t TimeOnly) String() string {
	return t.Format(textTime)
}

func (t *TimeOnly) Now() {
	tm := time.Now()
	*t = TimeOnly{time.Date(0, 0, 0, tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond(), tm.Location())}
}

// MarshalJSON implements the json.Marshaler interface.
func (t TimeOnly) MarshalJSON() ([]byte, error) {
	if t.Nanosecond() == -1 {
		return []byte(jsonNull), nil
	}
	return []byte(t.Format(jsonTime)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		*t = TimeOnly{}
		return nil
	}
	if strings.Compare(str, jsonNull) == 0 {
		*t = TimeOnly{time.Unix(0, -1)}
		return nil
	}
	tt, err := time.Parse(jsonTime, str)
	*t = TimeOnly{tt}
	return err
}

// Value - Implementation of valuer for database/sql
func (t TimeOnly) Value() (driver.Value, error) {
	return time.Time(t.Time), nil
}

// Scan - Implement the database/sql scanner interface
func (t *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		*t = TimeOnly{time.Unix(0, -1)}
		return nil
	}
	switch vt := value.(type) {
	case time.Time:
		*t = TimeOnly{vt}
	case []byte:
		tt, err := time.Parse(textTime, string(vt))
		if err != nil {
			return err
		}
		*t = TimeOnly{tt}
	case string:
		tt, err := time.Parse(textTime, vt)
		if err != nil {
			return err
		}
		*t = TimeOnly{tt}
	default:
		*t = TimeOnly{time.Unix(0, -1)}
	}
	return nil
}

///////////////////////////////////////////////////////////////////////
//
type DateTime struct {
	time.Time
}

func ParseDateTime(str string) (DateTime, error) {
	tt, err := time.Parse(textDateTime, str)
	out := DateTime{tt}
	return out, err
}

func (t DateTime) String() string {
	return t.Format(textDateTime)
}

func (t *DateTime) Now() {
	*t = DateTime{time.Now()}
}

// MarshalJSON implements the json.Marshaler interface.
func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(jsonNull), nil
	}
	return []byte(t.Format(jsonDateTime)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *DateTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		*t = DateTime{}
		return nil
	}
	if strings.Compare(str, jsonNull) == 0 {
		*t = DateTime{}
		return nil
	}
	tt, err := time.Parse(jsonDateTime, str)
	*t = DateTime{tt}
	return err
}

// Value - Implementation of valuer for database/sql
func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t.Time), nil
}

// Scan - Implement the database/sql scanner interface
func (t *DateTime) Scan(value interface{}) error {
	if value == nil {
		*t = DateTime{}
		return nil
	}
	switch vt := value.(type) {
	case time.Time:
		*t = DateTime{vt}
	case []byte:
		tt, err := time.Parse(textDateTime, string(vt))
		if err != nil {
			return err
		}
		*t = DateTime{tt}
	case string:
		tt, err := time.Parse(textDateTime, vt)
		if err != nil {
			return err
		}
		*t = DateTime{tt}
	default:
		*t = DateTime{}
	}
	return nil
}

///////////////////////////////////////////////////////////////////////
//
type FullTime struct {
	time.Time
}

func ParseFullTime(str string) (FullTime, error) {
	tt, err := time.Parse(textFullTime, str)
	out := FullTime{tt}
	return out, err
}

func (t FullTime) String() string {
	return t.Format(textFullTime)
}

func (t *FullTime) Now() {
	*t = FullTime{time.Now()}
}

// MarshalJSON implements the json.Marshaler interface.
func (t FullTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(jsonNull), nil
	}
	return []byte(t.Format(jsonFullTime)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *FullTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		*t = FullTime{}
		return nil
	}
	if strings.Compare(str, jsonNull) == 0 {
		*t = FullTime{}
		return nil
	}
	tt, err := time.Parse(jsonFullTime, str)
	*t = FullTime{tt}
	return err
}

// Value - Implementation of valuer for database/sql
func (t FullTime) Value() (driver.Value, error) {
	return time.Time(t.Time), nil
}

// Scan - Implement the database/sql scanner interface
func (t *FullTime) Scan(value interface{}) error {
	if value == nil {
		*t = FullTime{}
		return nil
	}
	switch vt := value.(type) {
	case time.Time:
		*t = FullTime{vt}
	case []byte:
		tt, err := time.Parse(textFullTime, string(vt))
		if err != nil {
			return err
		}
		*t = FullTime{tt}
	case string:
		tt, err := time.Parse(textFullTime, vt)
		if err != nil {
			return err
		}
		*t = FullTime{tt}
	default:
		*t = FullTime{}
	}
	return nil
}
