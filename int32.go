// This file was generated by cmd
package optional

import "encoding/json"

// Int32 represents an int32 that may be optional
type Int32 struct {
	value *int32
}

// NewInt32 creates an optional.Int32
func NewInt32(val int32) Int32 {
	return Int32{
		value: &val,
	}
}

// SetValue sets the int32 value
func (opt *Int32) SetValue(val int32) {
	opt.value = &val
}

func (opt Int32) initValue() int32 {
	var val int32
	return val
}

// Value returns the int32 value or init value if not present
func (opt Int32) Value() int32 {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Int32) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Int32) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Int32) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val int32
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}
