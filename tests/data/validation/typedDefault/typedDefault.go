// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import yaml "gopkg.in/yaml.v3"

type TypedDefault struct {
	// TopLevelDomains corresponds to the JSON schema field "topLevelDomains".
	TopLevelDomains []string `json:"topLevelDomains,omitempty" yaml:"topLevelDomains,omitempty" mapstructure:"topLevelDomains,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TypedDefault) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain TypedDefault
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if v, ok := raw["topLevelDomains"]; !ok || v == nil {
		plain.TopLevelDomains = []string{
			".com",
			".org",
			".info",
			".gov",
		}
	}
	*j = TypedDefault(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *TypedDefault) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain TypedDefault
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if v, ok := raw["topLevelDomains"]; !ok || v == nil {
		plain.TopLevelDomains = []string{
			".com",
			".org",
			".info",
			".gov",
		}
	}
	*j = TypedDefault(plain)
	return nil
}

// Verify checks all fields on the struct match the schema.
func (plain *TypedDefault) Verify() error {
	return nil
}
