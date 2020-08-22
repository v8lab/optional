// This file was generated by cmd
package optional

import "encoding/json"

// Int16 represents an int16 that may be optional
type Int16 struct {
	value *int16
}

// NewInt16 creates an optional.Int16
func NewInt16(val int16) Int16 {
	return Int16{
		value: &val,
	}
}

// SetValue sets the int16 value
func (opt *Int16) SetValue(val int16) {
	opt.value = &val
}

func (opt Int16) initValue() int16 {
	var val int16
	return val
}

// Value returns the int16 value or init value if not present
func (opt Int16) Value() int16 {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Int16) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Int16) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Int16) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val int16
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}