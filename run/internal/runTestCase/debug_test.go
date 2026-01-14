package runTestCase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"gorm.io/datatypes"
)

func TestConvertRequestJsonMap(t *testing.T) {
	// Construct a Request with specific Json content
	req := &automation.Request{
		Method: "POST",
		URL:    "http://localhost",
		Json: datatypes.JSONMap{
			"UserName": "Taylor",
			"age":      18,
			"Nested": map[string]interface{}{
				"Field_A": "Value_A",
			},
		},
	}

	// 1. Test StructToMap
	m := StructToMap(req)
	fmt.Printf("StructToMap Result: %+v\n", m)

	bodyVal, ok := m["body"]
	assert.True(t, ok, "key 'body' should exist in map")

	bodyMap, ok := bodyVal.(map[string]interface{})
	assert.True(t, ok, "value should be map[string]interface{}")

	assert.Equal(t, "Taylor", bodyMap["UserName"])
	assert.Equal(t, 18, bodyMap["age"])

	// 2. Test Mapstructure Decode
	hrpReq := &hrp.Request{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           hrpReq,
	})
	assert.NoError(t, err)

	err = decoder.Decode(m)
	assert.NoError(t, err)

	fmt.Printf("HRP Request Body: %+v\n", hrpReq.Body)

	// Verify hrpReq.Body
	hrpBodyMap, ok := hrpReq.Body.(map[string]interface{})
	assert.True(t, ok, "hrpReq.Body should be map[string]interface{}")

	assert.Equal(t, "Taylor", hrpBodyMap["UserName"])
	nestedMap, ok := hrpBodyMap["Nested"].(map[string]interface{})
	assert.True(t, ok, "Nested should be map[string]interface{}")
	assert.Equal(t, "Value_A", nestedMap["Field_A"])
}

func TestStructToMap_DatatypesJSONMap(t *testing.T) {
	// Verify reflect type logic
	typ := reflect.TypeOf(datatypes.JSONMap{})
	fmt.Printf("datatypes.JSONMap Type Name: %s, Kind: %s\n", typ.Name(), typ.Kind())

	val := datatypes.JSONMap{"k": "v"}
	v := reflect.ValueOf(val)

	if v.Type() == typ {
		fmt.Println("Type match!")
	} else {
		fmt.Println("Type mismatch!")
	}
}
