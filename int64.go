package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

type Int64 struct {
	int64 int64
	valid bool
}

func (f *Int64) Set(val int64) {
	f.int64 = val
	f.valid = true
}

func (f *Int64) Get() int64 {
	if !f.valid {
		return 0
	}
	return f.int64
}

func (f *Int64) Null() bool {
	return f.valid
}

func (f *Int64) Nullify() {
	f.valid = false
}

func (f Int64) String() string {
	if f.valid {
		return strconv.FormatInt(f.int64, 10)
	}
	return "null"
}

func (f Int64) MarshalJSON() ([]byte, error) {
	if f.valid {
		return json.Marshal(f.int64)
	}
	return json.Marshal(nil)
}

func (f *Int64) UnmarshalJSON(data []byte) (err error) {
	var temp *int64
	if err = json.Unmarshal(data, &temp); err != nil {
		return
	}
	if temp != nil {
		f.valid = true
		f.int64 = *temp
	} else {
		f.valid = false
	}
	return
}

func (f *Int64) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
		f.valid = false
		return
	case int64:
		f.int64, f.valid = v, true
		return
	case bool:
		if v {
			f.int64, f.valid = 1, true
		} else {
			f.int64, f.valid = 0, true
		}
		return
	default:
		return fmt.Errorf("not supported: %T %v", value, value)
	}
}

func (f Int64) Value() (driver.Value, error) {
	if f.valid {
		return f.int64, nil
	}
	return nil, nil
}

