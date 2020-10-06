package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

type Float64 struct {
	float float64
	valid bool
}

func (f *Float64) Set(val float64) {
	f.float = val
	f.valid = true
}

func (f *Float64) Get() float64 {
	if !f.valid {
		return 0.0
	}
	return f.float
}

func (f *Float64) Null() bool {
	return f.valid
}

func (f *Float64) Nullify() {
	f.valid = false
}

func (f Float64) String() string {
	if f.valid {
		return strconv.FormatFloat(f.float, 'e', -1, 64)
	}
	return "null"
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if f.valid {
		return json.Marshal(f.float)
	}
	return json.Marshal(nil)
}

func (f *Float64) UnmarshalJSON(data []byte) (err error) {
	var temp *float64
	if err = json.Unmarshal(data, &temp); err != nil {
		return
	}
	if temp != nil {
		f.valid = true
		f.float = *temp
	} else {
		f.valid = false
	}
	return
}

func (f *Float64) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
		f.valid = false
		return
	case float64:
		f.float, f.valid = v, true
		return
	case []uint8:
		if f.float, err = strconv.ParseFloat(string(v), 64); err == nil {
			f.valid = true
		}
		return
	default:
		return fmt.Errorf("not supported: %T %v", value, value)
	}
}

func (f Float64) Value() (driver.Value, error) {
	if f.valid {
		return f.float, nil
	}
	return nil, nil
}

