package jsonb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONB struct{}

func (m *JSONB) Value() (driver.Value, error) {
	val, err := json.Marshal(m)

	return driver.Value(string(val)), err
}

func (JSONB) GormDataType() string {
	return "JSONB"
}

func Scan(val interface{}, res interface{}) error {
	t, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value: %v", val))
	}

	return json.Unmarshal(t, &res)
}
