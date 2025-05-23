// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type MinLength struct {
	// MyNullableString corresponds to the JSON schema field "myNullableString".
	MyNullableString *string `json:"myNullableString,omitempty" yaml:"myNullableString,omitempty" mapstructure:"myNullableString,omitempty"`

	// MyString corresponds to the JSON schema field "myString".
	MyString string `json:"myString" yaml:"myString" mapstructure:"myString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MinLength) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["myString"]; raw != nil && !ok {
		return fmt.Errorf("field myString in MinLength: required")
	}
	type Plain MinLength
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if plain.MyNullableString != nil && len(*plain.MyNullableString) < 10 {
		return fmt.Errorf("field %s length: must be >= %d", "myNullableString", 10)
	}
	if len(plain.MyString) < 5 {
		return fmt.Errorf("field %s length: must be >= %d", "myString", 5)
	}
	*j = MinLength(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *MinLength) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if _, ok := raw["myString"]; raw != nil && !ok {
		return fmt.Errorf("field myString in MinLength: required")
	}
	type Plain MinLength
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if plain.MyNullableString != nil && len(*plain.MyNullableString) < 10 {
		return fmt.Errorf("field %s length: must be >= %d", "myNullableString", 10)
	}
	if len(plain.MyString) < 5 {
		return fmt.Errorf("field %s length: must be >= %d", "myString", 5)
	}
	*j = MinLength(plain)
	return nil
}

// Verify checks all fields on the struct match the schema.
func (plain *MinLength) Verify() error {
	if plain.MyNullableString != nil && len(*plain.MyNullableString) < 10 {
		return fmt.Errorf("field %s length: must be >= %d", "myNullableString", 10)
	}
	if len(plain.MyString) < 5 {
		return fmt.Errorf("field %s length: must be >= %d", "myString", 5)
	}
	return nil
}
