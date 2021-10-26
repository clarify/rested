package jsonschema_test

import (
	"testing"

	"github.com/clarify/rested/schema"
)

func TestBoolValidatorEncode(t *testing.T) {
	testCase := encoderTestCase{
		name: ``,
		schema: schema.Schema{
			Fields: schema.Fields{
				"b": {
					Validator: &schema.Bool{},
				},
			},
		},
		customValidate: fieldValidator("b", `{"type": "boolean"}`),
	}
	testCase.Run(t)
}
