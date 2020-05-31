package pdu

import (
	"fmt"
)

type Map map[Name]Field

func (m Map) Set(k Name, v interface{}) error {
	switch v.(type) {
	case string:
		m[k] = New(k, []byte(v.(string)))
	default:
		return fmt.Errorf("unsopported field data %#v", v)
	}

	return nil
}
