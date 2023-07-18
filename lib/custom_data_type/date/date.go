package date

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Date time.Time

var _ json.Unmarshaler = &Date{}
var _ json.Marshaler = &Date{}

func (d *Date) Time() time.Time {
	return time.Time(*d)
}

func (d *Date) String() string {
	return d.Time().Format("2006-01-02")
}

func (d *Date) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	tmpDate := Date(t)
	*d = tmpDate

	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return driver.Value(d.Time().Format("2006-01-02")), nil
}

func (d *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time().UTC().Format("2006-01-02") + `"`), nil
}

func (Date) GormDataType() string {
	return "date"
}

func (Date) FromString(date string) (Date, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Date{}, err
	}

	return Date(t), nil
}
