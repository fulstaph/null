package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

type Bool struct {
	bool  bool
	valid bool
}

func (b *Bool) Set(val bool) {
	b.bool = val
	b.valid = true
}

func (b *Bool) Get() bool {
	if !b.valid { return false }
	return b.bool
}

func (b *Bool) Null() bool {
	return b.valid
}

func (b *Bool) Nullify() {
	b.valid = false
}

func (b Bool) String() string {
	if b.valid {
		return strconv.FormatBool(b.bool)
	}
	return "null"
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.valid {
		return json.Marshal(b.bool)
	}
	return json.Marshal(nil)
}

func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	var temp *bool
	if err = json.Unmarshal(data, &temp); err != nil {
		return
	}
	if temp != nil {
		b.valid = true
		b.bool = *temp
	} else {
		b.valid = false
	}
	return
}

func (b *Bool) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
		b.valid = false
		return
	case bool:
		b.bool, b.valid = v, true
		return
	case int64:
		if v == 0 {
			b.bool, b.valid = false, true
		} else {
			b.bool, b.valid = true, true
		}
		return
	default:
		return fmt.Errorf("not supported: %T %v", value, value)
	}
}

func (b Bool) Value() (driver.Value, error) {
	if b.valid {
		return b.bool, nil
	}
	return nil, nil
}

