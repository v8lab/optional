// This file was generated by cmd
package optional

import "encoding/json"

// Bool represents an bool that may be optional
type Bool struct {
	value *bool
}

// NewBool creates an optional.Bool
func NewBool(val bool) Bool {
	return Bool{
		value: &val,
	}
}

// SetValue sets the bool value
func (opt *Bool) SetValue(val bool) {
	opt.value = &val
}

func (opt Bool) initValue() bool {
	var val bool
	return val
}

// Value returns the bool value or init value if not present
func (opt Bool) Value() bool {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Bool) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Bool) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Bool) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val bool
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}
