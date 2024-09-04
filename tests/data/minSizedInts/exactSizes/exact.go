// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"

type Exact struct {
	// I16 corresponds to the JSON schema field "i16".
	I16 int16 `json:"i16" yaml:"i16" mapstructure:"i16"`

	// I32 corresponds to the JSON schema field "i32".
	I32 int32 `json:"i32" yaml:"i32" mapstructure:"i32"`

	// I64 corresponds to the JSON schema field "i64".
	I64 int64 `json:"i64" yaml:"i64" mapstructure:"i64"`

	// I8 corresponds to the JSON schema field "i8".
	I8 int8 `json:"i8" yaml:"i8" mapstructure:"i8"`

	// U16 corresponds to the JSON schema field "u16".
	U16 uint16 `json:"u16" yaml:"u16" mapstructure:"u16"`

	// U32 corresponds to the JSON schema field "u32".
	U32 uint32 `json:"u32" yaml:"u32" mapstructure:"u32"`

	// U64 corresponds to the JSON schema field "u64".
	U64 uint64 `json:"u64" yaml:"u64" mapstructure:"u64"`

	// U8 corresponds to the JSON schema field "u8".
	U8 uint8 `json:"u8" yaml:"u8" mapstructure:"u8"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Exact) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["i16"]; raw != nil && !ok {
		return fmt.Errorf("field i16 in Exact: required")
	}
	if _, ok := raw["i32"]; raw != nil && !ok {
		return fmt.Errorf("field i32 in Exact: required")
	}
	if _, ok := raw["i64"]; raw != nil && !ok {
		return fmt.Errorf("field i64 in Exact: required")
	}
	if _, ok := raw["i8"]; raw != nil && !ok {
		return fmt.Errorf("field i8 in Exact: required")
	}
	if _, ok := raw["u16"]; raw != nil && !ok {
		return fmt.Errorf("field u16 in Exact: required")
	}
	if _, ok := raw["u32"]; raw != nil && !ok {
		return fmt.Errorf("field u32 in Exact: required")
	}
	if _, ok := raw["u64"]; raw != nil && !ok {
		return fmt.Errorf("field u64 in Exact: required")
	}
	if _, ok := raw["u8"]; raw != nil && !ok {
		return fmt.Errorf("field u8 in Exact: required")
	}
	type Plain Exact
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 32767 < plain.I16 {
		return fmt.Errorf("field %s: must be <= %v", "i16", 32767)
	}
	if -32768 > plain.I16 {
		return fmt.Errorf("field %s: must be >= %v", "i16", -32768)
	}
	if 2147483647 < plain.I32 {
		return fmt.Errorf("field %s: must be <= %v", "i32", 2147483647)
	}
	if -2147483648 > plain.I32 {
		return fmt.Errorf("field %s: must be >= %v", "i32", -2147483648)
	}
	if 9223372036854775807 < plain.I64 {
		return fmt.Errorf("field %s: must be <= %v", "i64", 9223372036854775807)
	}
	if -9223372036854775808 > plain.I64 {
		return fmt.Errorf("field %s: must be >= %v", "i64", -9223372036854775808)
	}
	if 127 < plain.I8 {
		return fmt.Errorf("field %s: must be <= %v", "i8", 127)
	}
	if -128 > plain.I8 {
		return fmt.Errorf("field %s: must be >= %v", "i8", -128)
	}
	if 65535 < plain.U16 {
		return fmt.Errorf("field %s: must be <= %v", "u16", 65535)
	}
	if 0 > plain.U16 {
		return fmt.Errorf("field %s: must be >= %v", "u16", 0)
	}
	if 4294967295 < plain.U32 {
		return fmt.Errorf("field %s: must be <= %v", "u32", 4294967295)
	}
	if 0 > plain.U32 {
		return fmt.Errorf("field %s: must be >= %v", "u32", 0)
	}
	if 9223372036854775807 < plain.U64 {
		return fmt.Errorf("field %s: must be <= %v", "u64", 9223372036854775807)
	}
	if 0 > plain.U64 {
		return fmt.Errorf("field %s: must be >= %v", "u64", 0)
	}
	if 255 < plain.U8 {
		return fmt.Errorf("field %s: must be <= %v", "u8", 255)
	}
	if 0 > plain.U8 {
		return fmt.Errorf("field %s: must be >= %v", "u8", 0)
	}
	*j = Exact(plain)
	return nil
}