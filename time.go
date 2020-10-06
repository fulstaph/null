package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool
}

func (t *Time) Null() bool {
	return t.Valid
}

func (t *Time) Nullify() {
	t.Valid = false
}

func (t Time) String() string {
	if t.Valid {
		return t.Time.Format("2006-01-02T15:04:05Z07:00")
	}
	return "null"
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time.Format("2006-01-02 15:04:05"))
	}
	return json.Marshal(nil)
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 || string(data) == "null" {
		t.Valid = false
		return
	}
	if t.Time, err = time.Parse(`"2006-01-02T15:04:05Z07:00"`, string(data)); err != nil {
		if t.Time, err = time.Parse(`"2006-01-02T15:04:05Z0700"`, string(data)); err != nil {
			if t.Time, err = time.Parse(`"2006-01-02T15:04:05"`, string(data)); err != nil {
				if t.Time, err = time.Parse(`"2006-01-02T15:04"`, string(data)); err != nil {
					if t.Time, err = time.Parse(`"2006-01-02 15:04:05 -07:00"`, string(data)); err != nil {
						if t.Time, err = time.Parse(`"2006-01-02 15:04:05 -0700"`, string(data)); err != nil {
							if t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data)); err != nil {
								if t.Time, err = time.Parse(`"2006-01-02 15:04"`, string(data)); err != nil {
									if t.Time, err = time.Parse(`"2006-01-02"`, string(data)); err != nil {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
	t.Valid = true
	return
}

func (t *Time) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
		t.Valid = false
		return
	case time.Time:
		t.Time, t.Valid = v, true
		return
	default:
		return fmt.Errorf("not supported: %T %v", value, value)
	}
}

func (t Time) Value() (driver.Value, error) {
	if t.Valid {
		return t.Time, nil
	}
	return nil, nil
}

type TimeT struct {
	Time
}

func (t TimeT) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time.Time.Format("2006-01-02T15:04:05"))
	}
	return json.Marshal(nil)
}

type TimeTz struct {
	Time
}

func (t TimeTz) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time.Time.Format("2006-01-02T15:04:05-07:00"))
	}
	return json.Marshal(nil)
}

type TimeDate struct {
	Time
}

func (t TimeDate) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time.Time.Format("2006-01-02"))
	}
	return json.Marshal(nil)
}
