package runTestCase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/server/v2/model/automation"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
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

	jsonVal, ok := m["json"]
	assert.True(t, ok, "key 'json' should exist in map")
	
	jsonMap, ok := jsonVal.(map[string]interface{})
	assert.True(t, ok, "value should be map[string]interface{}")
	
	assert.Equal(t, "Taylor", jsonMap["UserName"])
	assert.Equal(t, 18, jsonMap["age"])

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

	fmt.Printf("HRP Request Json: %+v\n", hrpReq.Json)

	// Verify hrpReq.Json
	hrpJsonMap, ok := hrpReq.Json.(map[string]interface{})
	assert.True(t, ok, "hrpReq.Json should be map[string]interface{}")
	
	assert.Equal(t, "Taylor", hrpJsonMap["UserName"])
	assert.Equal(t, "Value_A", hrpJsonMap["Nested"].(map[string]interface{})["Field_A"])
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
