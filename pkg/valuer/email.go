package valuer

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"regexp"
	"strings"

	"github.com/gear6io/pragmata/pkg/errors"
)

const (
	emailRegexString string = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	emailRegex        = regexp.MustCompile(emailRegexString)
	_          Valuer = (*Email)(nil)

	CodeInvalidEmail = errors.MustNewCode("invalid_email")
)

type Email struct {
	val string
}

func NewEmail(val string) (Email, error) {
	if !emailRegex.MatchString(strings.ToLower(strings.TrimSpace(val))) {
		return Email{}, errors.NewInvalidInputf(CodeInvalidEmail, "invalid email %s", val)
	}

	return Email{val: strings.ToLower(strings.TrimSpace(val))}, nil
}

func MustNewEmail(val string) Email {
	email, err := NewEmail(val)
	if err != nil {
		panic(err)
	}

	return email
}

func (enum Email) IsZero() bool {
	return enum.val == ""
}

func (enum Email) StringValue() string {
	return enum.val
}

func (enum Email) String() string {
	return enum.val
}

func (enum Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.StringValue())
}

func (enum *Email) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	var err error
	*enum, err = NewEmail(str)
	if err != nil {
		return err
	}

	return nil
}

func (enum Email) Value() (driver.Value, error) {
	return enum.StringValue(), nil
}

func (enum *Email) Scan(val interface{}) error {
	if enum == nil {
		return errors.NewInternalf(errors.CodeInternal, "email: (nil \"%s\")", reflect.TypeOf(enum).String())
	}

	str, ok := val.(string)
	if !ok {
		return errors.NewInternalf(errors.CodeInternal, "email: (non-string \"%s\")", reflect.TypeOf(val).String())
	}

	if str == "" {
		*enum = Email{}
		return nil
	}

	var err error
	*enum, err = NewEmail(str)
	if err != nil {
		return err
	}

	return nil
}

func (enum *Email) UnmarshalText(text []byte) error {
	var err error
	*enum, err = NewEmail(string(text))
	if err != nil {
		return err
	}

	return nil
}

func (enum Email) MarshalText() (text []byte, err error) {
	return []byte(enum.StringValue()), nil
}

func (enum *Email) UnmarshalParam(param string) error {
	email, err := NewEmail(param)
	if err != nil {
		return err
	}

	*enum = email
	return nil
}
