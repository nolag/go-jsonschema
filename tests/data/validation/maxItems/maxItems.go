// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type MaxItems struct {
	// MyNestedArray corresponds to the JSON schema field "myNestedArray".
	MyNestedArray [][]interface{} `json:"myNestedArray,omitempty" yaml:"myNestedArray,omitempty" mapstructure:"myNestedArray,omitempty"`

	// MyStringArray corresponds to the JSON schema field "myStringArray".
	MyStringArray []string `json:"myStringArray,omitempty" yaml:"myStringArray,omitempty" mapstructure:"myStringArray,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MaxItems) UnmarshalJSON(value []byte) error {
	type Plain MaxItems
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if len(plain.MyNestedArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myNestedArray", 5)
	}
	for i1 := range plain.MyNestedArray {
		if len(plain.MyNestedArray[i1]) > 5 {
			return fmt.Errorf("field %s length: must be <= %d", fmt.Sprintf("myNestedArray[%d]", i1), 5)
		}
	}
	if len(plain.MyStringArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myStringArray", 5)
	}
	*j = MaxItems(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *MaxItems) UnmarshalYAML(value *yaml.Node) error {
	type Plain MaxItems
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if len(plain.MyNestedArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myNestedArray", 5)
	}
	for i1 := range plain.MyNestedArray {
		if len(plain.MyNestedArray[i1]) > 5 {
			return fmt.Errorf("field %s length: must be <= %d", fmt.Sprintf("myNestedArray[%d]", i1), 5)
		}
	}
	if len(plain.MyStringArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myStringArray", 5)
	}
	*j = MaxItems(plain)
	return nil
}

// Verify checks all fields on the struct match the schema.
func (plain *MaxItems) Verify() error {
	if len(plain.MyNestedArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myNestedArray", 5)
	}
	for i1 := range plain.MyNestedArray {
		if len(plain.MyNestedArray[i1]) > 5 {
			return fmt.Errorf("field %s length: must be <= %d", fmt.Sprintf("myNestedArray[%d]", i1), 5)
		}
	}
	if len(plain.MyStringArray) > 5 {
		return fmt.Errorf("field %s length: must be <= %d", "myStringArray", 5)
	}
	return nil
}
