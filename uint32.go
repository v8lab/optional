// This file was generated by cmd
package optional

import "encoding/json"

// Uint32 represents an uint32 that may be optional
type Uint32 struct {
	value *uint32
}

// NewUint32 creates an optional.Uint32
func NewUint32(val uint32) Uint32 {
	return Uint32{
		value: &val,
	}
}

// SetValue sets the uint32 value
func (opt *Uint32) SetValue(val uint32) {
	opt.value = &val
}

func (opt Uint32) initValue() uint32 {
	var val uint32
	return val
}

// Value returns the uint32 value or init value if not present
func (opt Uint32) Value() uint32 {
	if opt.value != nil {
		return *opt.value
	} else {
		return opt.initValue()
	}
}

// IsPresent returns whether or not the value is present
func (opt Uint32) IsPresent() bool {
	return opt.value != nil
}

// MarshalJSON implements the json.MarshalJSON interface
func (opt Uint32) MarshalJSON() ([]byte, error) {
	if opt.value != nil {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (opt *Uint32) UnmarshalJSON(data []byte) error {
	opt.value = nil

	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var val uint32
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	opt.value = &val
	return nil
}
