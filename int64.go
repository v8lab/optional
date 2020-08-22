// This file was generated by cmd
package optional

import "encoding/json"

// Int64 represents an int64 that may be optional
type Int64 struct {
	value *int64
}

// NewInt64 creates an optional.Int64
func NewInt64(val int64) Int64 {
	return Int64{
		value: &val,
	}
}

// SetValue sets the int64 value
func (opt *Int64) SetValue(val int64) {
	opt.value = &val
}

func (opt Int64) initValue() int64 {
	var val int64
	return val
}

// Value returns the int64 value or init value if not present
func (opt Int64) Value() int64 {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Int64) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Int64) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Int64) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val int64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}