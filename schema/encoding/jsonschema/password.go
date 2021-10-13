package jsonschema

import "github.com/searis/rested/schema"

type passwordBuilder schema.Password

func (v passwordBuilder) BuildJSONSchema() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"type":   "string",
		"format": "password",
	}
	if v.MinLen > 0 {
		m["minLength"] = v.MinLen
	}
	if v.MaxLen > 0 {
		m["maxLength"] = v.MaxLen
	}
	return m, nil
}
