package valuer

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

var _ Valuer = (*RawString)(nil)

type RawString struct {
	val string
}

func NewRawString(val string) RawString {
	return RawString{val: strings.TrimSpace(val)}
}

func (enum RawString) IsZero() bool {
	return enum.val == ""
}

func (enum RawString) StringValue() string {
	return enum.val
}

func (enum RawString) String() string {
	return enum.val
}

func (enum RawString) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.StringValue())
}

func (enum *RawString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*enum = NewRawString(str)
	return nil
}

func (enum RawString) Value() (driver.Value, error) {
	return enum.StringValue(), nil
}

func (enum *RawString) Scan(val any) error {
	if enum == nil {
		return errors.New("string: (nil \"%s\")" + reflect.TypeOf(enum).String())
	}

	if val == nil {
		// Return an empty string
		*enum = NewRawString("")
		return nil
	}

	str, ok := val.(string)
	if !ok {
		return errors.New("string: (non-string \"%s\")" + reflect.TypeOf(val).String())
	}

	*enum = NewRawString(str)
	return nil
}

func (enum *RawString) UnmarshalText(text []byte) error {
	*enum = NewRawString(string(text))
	return nil
}

func (enum RawString) MarshalText() (text []byte, err error) {
	return []byte(enum.StringValue()), nil
}

// Implement Gin's BindUnmarshaler interface
// See https://github.com/SigNoz/signoz/pull/10219 description for additional details.
func (enum *RawString) UnmarshalParam(param string) error {
	*enum = NewRawString(param)
	return nil
}
