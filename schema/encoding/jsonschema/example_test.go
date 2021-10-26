package jsonschema_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/clarify/rested/schema"
	"github.com/clarify/rested/schema/encoding/jsonschema"
)

func ExampleEncoder() {
	s := schema.Schema{
		Fields: schema.Fields{
			"foo": schema.Field{
				Required:  true,
				Validator: &schema.Float{Boundaries: &schema.Boundaries{Min: 0, Max: math.Inf(1)}},
			},
			"bar": schema.Field{
				Validator: &schema.Integer{},
			},
			"baz": schema.Field{
				Description: "baz can not be set by the user",
				ReadOnly:    true,
				Validator:   &schema.String{MaxLen: 42},
			},
			"foobar": schema.Field{
				Description: "foobar can hold any valid JSON value",
			},
		},
	}
	b := new(bytes.Buffer)
	enc := jsonschema.NewEncoder(b)
	_ = enc.Encode(&s)
	b2 := new(bytes.Buffer)
	_ = json.Indent(b2, b.Bytes(), "", "| ")
	fmt.Println(b2)
	// Output: {
	// | "additionalProperties": false,
	// | "properties": {
	// | | "bar": {
	// | | | "type": "integer"
	// | | },
	// | | "baz": {
	// | | | "description": "baz can not be set by the user",
	// | | | "maxLength": 42,
	// | | | "readOnly": true,
	// | | | "type": "string"
	// | | },
	// | | "foo": {
	// | | | "minimum": 0,
	// | | | "type": "number"
	// | | },
	// | | "foobar": {
	// | | | "description": "foobar can hold any valid JSON value"
	// | | }
	// | },
	// | "required": [
	// | | "foo"
	// | ],
	// | "type": "object"
	// }
}
