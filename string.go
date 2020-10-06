package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type String struct {
	str   string
	valid bool
}

func (s *String) Get() string {
	if !s.valid { return "" }
	return s.str
}

func (s *String) Set(str string) {
	s.str = str
	s.valid = true
}

func (s *String) Null() bool {
	return s.valid
}

func (s *String) Nullify() {
	s.valid = true
}

func (s String) String() string {
	if s.valid {
		return s.str
	}
	return "null"
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.valid {
		return json.Marshal(s.str)
	}
	return json.Marshal(nil)
}

func (s *String) UnmarshalJSON(data []byte) (err error) {
	var temp *string
	if err = json.Unmarshal(data, &temp); err != nil {
		return
	}
	if temp != nil {
		s.valid = true
		s.str = *temp
	} else {
		s.valid = false
	}
	return
}

func (s *String) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
		s.valid = false
		return
	case string:
		s.str, s.valid = v, true
		return
	case []uint8:
		s.str, s.valid = string(v), true
		return
	default:
		return fmt.Errorf("not supported: %T %v", value, value)
	}
}

func (s String) Value() (driver.Value, error) {
	if s.valid {
		return s.str, nil
	}
	return nil, nil
}

