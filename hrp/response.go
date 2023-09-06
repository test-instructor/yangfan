package hrp

import (
	"bytes"
	builtinJSON "encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/jmespath/go-jmespath"
	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/hrp/internal/json"
	"github.com/test-instructor/yangfan/hrp/pkg/uixt"
)

var fieldTags = []string{"proto", "status_code", "headers", "cookies", "body", textExtractorSubRegexp}

type httpRespObjMeta struct {
	Proto      string            `json:"proto"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Cookies    map[string]string `json:"cookies"`
	Body       interface{}       `json:"body"`
}

func newHttpResponseObject(t *testing.T, parser *Parser, resp *http.Response) (*responseObject, error) {
	// prepare response headers
	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// prepare response cookies
	cookies := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	// read response body
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse response body
	var body interface{}
	if err := json.Unmarshal(respBodyBytes, &body); err != nil {
		// response body is not json, use raw body
		body = string(respBodyBytes)
	}

	respObjMeta := httpRespObjMeta{
		Proto:      resp.Proto,
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Cookies:    cookies,
		Body:       body,
	}

	return convertToResponseObject(t, parser, respObjMeta)
}

type wsCloseRespObject struct {
	StatusCode int    `json:"status_code"`
	Text       string `json:"body"`
}

func newWsCloseResponseObject(t *testing.T, parser *Parser, resp *wsCloseRespObject) (*responseObject, error) {
	return convertToResponseObject(t, parser, resp)
}

type wsReadRespObject struct {
	Message     interface{} `json:"body"`
	messageType int
}

func newWsReadResponseObject(t *testing.T, parser *Parser, resp *wsReadRespObject) (*responseObject, error) {
	byteMessage, ok := resp.Message.([]byte)
	if !ok {
		return nil, errors.New("websocket message type should be []byte")
	}
	var msg interface{}
	if err := json.Unmarshal(byteMessage, &msg); err != nil {
		// response body is not json, use raw body
		msg = string(byteMessage)
	}
	resp.Message = msg
	return convertToResponseObject(t, parser, resp)
}

func convertToResponseObject(t *testing.T, parser *Parser, respObjMeta interface{}) (*responseObject, error) {
	respObjMetaBytes, _ := json.Marshal(respObjMeta)
	var data interface{}
	decoder := json.NewDecoder(bytes.NewReader(respObjMetaBytes))
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil {
		global.GVA_LOG.Error("[convertToResponseObject] convert respObjectMeta to interface{} failed", zap.String("respObjectMeta", string(respObjMetaBytes)), zap.Error(err))
		return nil, err
	}
	return &responseObject{
		t:           t,
		parser:      parser,
		respObjMeta: data,
	}, nil
}

type responseObject struct {
	t                 *testing.T
	parser            *Parser
	respObjMeta       interface{}
	validationResults []*ValidationResult
}

const textExtractorSubRegexp string = `(.*)`

func (v *responseObject) searchField(field string, variablesMapping map[string]interface{}) interface{} {
	var result interface{} = field
	if strings.Contains(field, "$") {
		// parse reference variables in field before search
		var err error
		result, err = v.parser.Parse(field, variablesMapping)
		if err != nil {
			global.GVA_LOG.Error("[searchField] fail to parse field before search", zap.String("field", field), zap.Error(err))
		}
	}
	// search field using jmespath or regex if parsed field is still string and contains specified fieldTags
	if parsedField, ok := result.(string); ok && checkSearchField(parsedField) {
		if strings.Contains(field, textExtractorSubRegexp) {
			result = v.searchRegexp(parsedField)
		} else {
			result = v.searchJmespath(parsedField)
		}
	}
	return result
}

func (v *responseObject) Extract(extractors map[string]string, variablesMapping map[string]interface{}) map[string]interface{} {
	if extractors == nil {
		return nil
	}

	extractMapping := make(map[string]interface{})
	for key, value := range extractors {
		extractedValue := v.searchField(value, variablesMapping)
		global.GVA_LOG.Info("[Extract] extract value", zap.String("from", value), zap.Any("value", extractedValue))
		global.GVA_LOG.Info("[Extract] set variable", zap.String("variable", key), zap.Any("value", extractedValue))
		extractMapping[key] = extractedValue
	}

	return extractMapping
}

func (v *responseObject) Validate(iValidators []interface{}, variablesMapping map[string]interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 发生了 panic，将其转化为错误信息
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	for _, iValidator := range iValidators {
		validator, ok := iValidator.(Validator)
		if !ok {
			return errors.New("validator type error")
		}
		// parse check value
		checkItem := validator.Check
		checkValue := v.searchField(checkItem, variablesMapping)

		// get assert method
		assertMethod := validator.Assert
		assertFunc, ok := builtin.Assertions[assertMethod]
		if !ok {
			return errors.New(fmt.Sprintf("unexpected assertMethod: %v", assertMethod))
		}

		// parse expected value
		expectValue, err := v.parser.Parse(validator.Expect, variablesMapping)
		if err != nil {
			return err
		}
		validResult := &ValidationResult{
			Validator: Validator{
				Check:   validator.Check,
				Expect:  expectValue,
				Assert:  assertMethod,
				Message: validator.Message,
			},
			CheckValue:  checkValue,
			CheckResult: "fail",
		}

		// do assertion
		global.GVA_LOG.Debug("assertFunc", zap.Any("checkValue", checkValue), zap.Any("expectValue", expectValue))
		var result bool
		if checkValue != nil {
			result = assertFunc(v.t, checkValue, expectValue)
		}
		if result {
			validResult.CheckResult = "pass"
		}
		global.GVA_LOG.Debug("result", zap.Bool("result", result))
		global.GVA_LOG.Debug("v.validationResults", zap.Any("", v.validationResults))
		global.GVA_LOG.Debug("v.validResult", zap.Any("", validResult))
		v.validationResults = append(v.validationResults, validResult)
		global.GVA_LOG.Info("[Validate] validate result",
			zap.String("checkExpr", validator.Check),
			zap.String("assertMethod", assertMethod),
			zap.Any("expectValue", expectValue),
			zap.Any("expectValueType", builtin.InterfaceType(expectValue)),
			zap.Any("checkValue", checkValue),
			zap.Any("checkValueType", builtin.InterfaceType(checkValue)),
			zap.Bool("result", result),
			zap.String("validate", checkItem),
		)
		if !result {
			global.GVA_LOG.Error("[Validate] assert failed",
				zap.String("checkExpr", validator.Check),
				zap.String("assertMethod", assertMethod),
				zap.Any("expectValue", expectValue),
				zap.Any("expectValueType", builtin.InterfaceType(expectValue)),
				zap.Any("checkValue", checkValue),
				zap.Any("checkValueType", builtin.InterfaceType(checkValue)),
				zap.Bool("result", result),
				zap.String("validate", checkItem),
			)
			return errors.New("step validation failed")
		}
	}
	return nil
}

func checkSearchField(expr string) bool {
	for _, t := range fieldTags {
		if strings.Contains(expr, t) {
			return true
		}
	}
	return false
}

func (v *responseObject) searchJmespath(expr string) interface{} {
	checkValue, err := jmespath.Search(expr, v.respObjMeta)
	if err != nil {
		global.GVA_LOG.Error("[searchJmespath] search jmespath failed", zap.String("expr", expr), zap.Error(err))
		return expr // jmespath not found, return the expression
	}
	if number, ok := checkValue.(builtinJSON.Number); ok {
		checkNumber, err := parseJSONNumber(number)
		if err != nil {
			global.GVA_LOG.Error("[searchJmespath] convert json number failed", zap.String("expr", expr), zap.Error(err))
		}
		return checkNumber
	}
	return checkValue
}

func (v *responseObject) searchRegexp(expr string) interface{} {
	respMap, ok := v.respObjMeta.(map[string]interface{})
	if !ok {
		global.GVA_LOG.Error("[searchRegexp] convert respObjMeta to map failed", zap.Any("resp", v.respObjMeta))
		return expr
	}
	bodyStr, ok := respMap["body"].(string)
	if !ok {
		global.GVA_LOG.Error("[searchRegexp] convert body to string failed", zap.Any("resp", respMap))
		return expr
	}
	regexpCompile, err := regexp.Compile(expr)
	if err != nil {
		global.GVA_LOG.Error("[searchRegexp] compile regexp failed", zap.String("expr", expr), zap.Error(err))
		return expr
	}
	match := regexpCompile.FindStringSubmatch(bodyStr)
	if len(match) > 1 {
		return match[1] // return first matched result in parentheses
	}
	global.GVA_LOG.Error("[searchRegexp] search regexp failed", zap.String("expr", expr), zap.Any("match", match), zap.Any("resp", respMap))
	return expr
}

func validateUI(ud *uixt.DriverExt, iValidators []interface{}) (validateResults []*ValidationResult, err error) {
	for _, iValidator := range iValidators {
		validator, ok := iValidator.(Validator)
		if !ok {
			return nil, errors.New("validator type error")
		}

		validataResult := &ValidationResult{
			Validator:   validator,
			CheckResult: "fail",
		}

		// parse check value
		if !strings.HasPrefix(validator.Check, "ui_") {
			validataResult.CheckResult = "skip"
			global.GVA_LOG.Warn("[Validate] skip validator", zap.Any("validator", validator))
			validateResults = append(validateResults, validataResult)
			continue
		}

		expected, ok := validator.Expect.(string)
		if !ok {
			return nil, errors.New("validator expect should be string")
		}

		if !ud.DoValidation(validator.Check, validator.Assert, expected, validator.Message) {
			return validateResults, errors.New("step validation failed")
		}

		validataResult.CheckResult = "pass"
		validateResults = append(validateResults, validataResult)
	}
	return validateResults, nil
}

func newSkipObject(t *testing.T, parser *Parser) (*skipObject, error) {
	return &skipObject{
		t:           t,
		parser:      parser,
		respObjMeta: nil,
	}, nil
}

type skipObject struct {
	t                 *testing.T
	parser            *Parser
	respObjMeta       interface{}
	validationResults []*ValidationResult
}

func (v *skipObject) Validate(iValidators []interface{}, variablesMapping map[string]interface{}) (err error) {
	for _, Validators := range iValidators {
		iValidator := Validators.(map[string]interface{})
		validator := Validator{
			Check:   iValidator["check"].(string),
			Assert:  iValidator["assert"].(string),
			Message: iValidator["msg"].(string),
			Expect:  iValidator["check"],
		}
		// parse check value
		checkItem := validator.Check
		checkValue := v.searchField(checkItem, variablesMapping)

		// get assert method
		assertMethod := validator.Assert
		assertFunc, ok := builtin.Assertions[assertMethod]
		if !ok {
			return errors.New(fmt.Sprintf("unexpected assertMethod: %v", assertMethod))
		}

		// parse expected value
		expectValue, err := v.parser.Parse(validator.Expect, variablesMapping)
		if err != nil {
			return err
		}
		validResult := &ValidationResult{
			Validator: Validator{
				Check:   validator.Check,
				Expect:  expectValue,
				Assert:  assertMethod,
				Message: validator.Message,
			},
			CheckValue:  checkValue,
			CheckResult: "fail",
		}

		// do assertion
		result := assertFunc(v.t, checkValue, expectValue)
		if result {
			validResult.CheckResult = "pass"
		}
		v.validationResults = append(v.validationResults, validResult)
		global.GVA_LOG.Info("[Validate] validate result",
			zap.String("checkExpr", validator.Check),
			zap.String("assertMethod", assertMethod),
			zap.Any("expectValue", expectValue),
			zap.Any("expectValueType", builtin.InterfaceType(expectValue)),
			zap.Any("checkValue", checkValue),
			zap.Any("checkValueType", builtin.InterfaceType(checkValue)),
			zap.Bool("result", result),
			zap.String("validate", checkItem),
		)
		if !result {
			global.GVA_LOG.Error("[Validate] assert failed",
				zap.String("checkExpr", validator.Check),
				zap.String("assertMethod", assertMethod),
				zap.Any("expectValue", expectValue),
				zap.Any("expectValueType", builtin.InterfaceType(expectValue)),
				zap.Any("checkValue", checkValue),
				zap.Any("checkValueType", builtin.InterfaceType(checkValue)),
				zap.Bool("result", result),
				zap.String("validate", checkItem),
			)
			return errors.New("step validation failed")
		}
	}
	return nil
}

func (v *skipObject) searchField(field string, variablesMapping map[string]interface{}) interface{} {
	var result interface{} = field
	if strings.Contains(field, "$") {
		// parse reference variables in field before search
		var err error
		result, err = v.parser.Parse(field, variablesMapping)
		if err != nil {
			global.GVA_LOG.Error("[searchField] fail to parse field before search", zap.String("field", field), zap.Error(err))
		}
	}
	// search field using jmespath or regex if parsed field is still string and contains specified fieldTags
	if parsedField, ok := result.(string); ok && checkSearchField(parsedField) {
		if strings.Contains(field, textExtractorSubRegexp) {
			result = v.searchRegexp(parsedField)
		} else {
			result = v.searchJmespath(parsedField)
		}
	}
	return result
}

func (v *skipObject) searchJmespath(expr string) interface{} {
	checkValue, err := jmespath.Search(expr, v.respObjMeta)
	if err != nil {
		global.GVA_LOG.Error("[searchJmespath] search jmespath failed", zap.String("expr", expr), zap.Error(err))
		return expr // jmespath not found, return the expression
	}
	if number, ok := checkValue.(builtinJSON.Number); ok {
		checkNumber, err := parseJSONNumber(number)
		if err != nil {
			global.GVA_LOG.Error("[searchJmespath] convert json number failed", zap.String("expr", expr), zap.Error(err))
		}
		return checkNumber
	}
	return checkValue
}

func (v *skipObject) searchRegexp(expr string) interface{} {
	respMap, ok := v.respObjMeta.(map[string]interface{})
	if !ok {
		global.GVA_LOG.Error("[searchRegexp] convert respObjMeta to map failed", zap.Any("resp", v.respObjMeta))
		return expr
	}
	bodyStr, ok := respMap["body"].(string)
	if !ok {
		global.GVA_LOG.Error("[searchRegexp] convert body to string failed", zap.Any("resp", respMap))
		return expr
	}
	regexpCompile, err := regexp.Compile(expr)
	if err != nil {
		global.GVA_LOG.Error("[searchRegexp] compile regexp failed", zap.String("expr", expr), zap.Error(err))
		return expr
	}
	match := regexpCompile.FindStringSubmatch(bodyStr)
	if len(match) > 1 {
		return match[1] // return first matched result in parentheses
	}
	global.GVA_LOG.Error("[searchRegexp] search regexp failed", zap.String("expr", expr), zap.Any("match", match), zap.Any("resp", respMap))
	return expr
}
