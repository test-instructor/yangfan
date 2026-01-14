package hrp

import (
	builtinJSON "encoding/json"
	"fmt"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/httprunner/funplugin"
	"github.com/httprunner/funplugin/shared"
	"github.com/maja42/goval"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/parsing/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/code"
)

func newParser() *Parser {
	return &Parser{}
}

type Parser struct {
	plugin funplugin.IPlugin // plugin is used to call functions
}

func buildURL(baseURL, stepURL string) string {
	// 解析url格式
	uStep, err := url.Parse(stepURL)
	if err != nil {
		log.Error().Str("stepURL", stepURL).Err(err).Msg("[buildURL] parse url failed")
		return ""
	}

	// step url is absolute url
	// 如果url存在域名直接返回url
	if uStep.Host != "" {
		return stepURL
	}

	// step url is relative, based on base url
	// 解析base url格式
	if baseURL != "" && !strings.Contains(baseURL, "://") {
		baseURL = "http://" + baseURL
	}
	uConfig, err := url.Parse(baseURL)
	if err != nil {
		log.Error().Str("baseURL", baseURL).Err(err).Msg("[buildURL] parse url failed")
		return ""
	}

	// merge url
	// 合并url
	uStep.Scheme = uConfig.Scheme
	uStep.Host = uConfig.Host
	uStep.Path = path.Join(uConfig.Path, uStep.Path)

	// base url missed
	// 返回合并后的url
	return uStep.String()
}

func (p *Parser) ParseHeaders(rawHeaders map[string]string, variablesMapping map[string]interface{}) (map[string]string, error) {
	parsedHeaders := make(map[string]string)
	headers, err := p.Parse(rawHeaders, variablesMapping)
	if err != nil {
		return rawHeaders, err
	}
	for k, v := range headers.(map[string]interface{}) {
		parsedHeaders[k] = convertString(v)
	}
	return parsedHeaders, nil
}

func convertString(raw interface{}) string {
	if str, ok := raw.(string); ok {
		return str
	}
	if float, ok := raw.(float64); ok {
		// f: avoid conversion to exponential notation
		return strconv.FormatFloat(float, 'f', -1, 64)
	}
	// convert to string
	return fmt.Sprintf("%v", raw)
}

// ParseString 将interface(变量、函数)内容转换为string
func (p *Parser) Parse(raw interface{}, variablesMapping map[string]interface{}) (interface{}, error) {
	rawValue := reflect.ValueOf(raw)
	switch rawValue.Kind() {
	case reflect.String:
		// raw 为string
		// json.Number
		// 如果为数字则转换为数字
		if rawValue, ok := raw.(builtinJSON.Number); ok {
			return parseJSONNumber(rawValue)
		}
		// other string
		// 如果为字符串则转换为字符串
		value := rawValue.String()
		value = strings.TrimSpace(value)
		return p.ParseString(value, variablesMapping)
	case reflect.Slice:
		// raw 为切片
		// 获取每个元素内容并获取元素中的函数\变量
		parsedSlice := make([]interface{}, rawValue.Len())
		for i := 0; i < rawValue.Len(); i++ {
			parsedValue, err := p.Parse(rawValue.Index(i).Interface(), variablesMapping)
			if err != nil {
				return raw, err
			}
			parsedSlice[i] = parsedValue
		}
		return parsedSlice, nil
	case reflect.Map: // convert any map to map[string]interface{}
		// raw 为map
		// 获取每个元素内容并获取元素中的函数\变量
		// map类型一般为header\body\params等格式
		parsedMap := make(map[string]interface{})
		for _, k := range rawValue.MapKeys() {
			// 获取key,并对key进行解析
			parsedKey, err := p.ParseString(k.String(), variablesMapping)
			if err != nil {
				return raw, err
			}
			v := rawValue.MapIndex(k)
			// 获取value,并对value进行解析
			parsedValue, err := p.Parse(v.Interface(), variablesMapping)
			if err != nil {
				return raw, err
			}

			key := convertString(parsedKey)
			parsedMap[key] = parsedValue
		}
		return parsedMap, nil
	default:
		// other types, e.g. nil, int, float, bool
		return builtin.TypeNormalization(raw), nil
	}
}

func parseJSONNumber(raw builtinJSON.Number) (value interface{}, err error) {
	if strings.Contains(raw.String(), ".") {
		// float64
		value, err = raw.Float64()
	} else {
		// int64
		value, err = raw.Int64()
	}
	if err != nil {
		return nil, errors.Wrap(code.ParseError,
			fmt.Sprintf("parse json number failed: %v", err))
	}
	return value, nil
}

const (
	regexVariable     = `[a-zA-Z_]\w*`    // variable name should start with a letter or underscore
	regexFunctionName = `[a-zA-Z_]\w*`    // function name should start with a letter or underscore
	regexNumber       = `^-?\d+(\.\d+)?$` // match number, e.g. 123, -123, 1.23, -1.23
)

var (
	regexCompileVariable = regexp.MustCompile(fmt.Sprintf(`\$\{(%s)\}|\$(%s)`, regexVariable, regexVariable))     // parse ${var} or $var
	regexCompileFunction = regexp.MustCompile(fmt.Sprintf(`\$\{(%s)\(([\$\w\.\-/\s=,]*)\)\}`, regexFunctionName)) // parse ${func1($a, $b)}
	regexCompileNumber   = regexp.MustCompile(regexNumber)                                                        // parse number
)

// ParseString parse string with variables
// 将变量解析为常量
// raw:需要解析函数、变量的字符串
// variablesMapping:变量映射，key为变量名，value为变量值
func (p *Parser) ParseString(raw string, variablesMapping map[string]interface{}) (interface{}, error) {
	matchStartPosition := 0
	parsedString := ""
	remainedString := raw
	// 在字符串中提取变量格式的字符串
	for matchStartPosition < len(raw) {
		// locate $ char position
		startPosition := strings.Index(remainedString, "$")
		if startPosition == -1 { // no $ found
			// append remained string
			parsedString += remainedString
			break
		}

		// found $, check if variable or function
		matchStartPosition += startPosition
		parsedString += remainedString[0:startPosition]
		remainedString = remainedString[startPosition:]

		// Notice: notation priority
		// $$ > ${func($a, $b)} > $var

		// search $$, use $$ to escape $ notation
		if strings.HasPrefix(remainedString, "$$") { // found $$
			matchStartPosition += 2
			parsedString += "$"
			remainedString = remainedString[2:]
			continue
		}

		// search function like ${func($a, $b)}
		// 提取字符串中的的函数
		// 返回内容为[]string，第一个元素为函数名，第二个元素为函数参数
		funcMatched := regexCompileFunction.FindStringSubmatch(remainedString)
		if len(funcMatched) == 3 {
			funcName := funcMatched[1]
			argsStr := funcMatched[2]
			// 获取变量，返回内容为[]interface{}为函数参数的值
			// 变量值类型不确定，故采用interface{}类型
			arguments, err := parseFunctionArguments(argsStr)
			if err != nil {
				return raw, errors.Wrap(code.ParseFunctionError, err.Error())
			}
			// 解析变量
			parsedArgs, err := p.Parse(arguments, variablesMapping)
			if err != nil {
				return raw, err
			}
			// 调用函数
			result, err := p.callFunc(funcName, parsedArgs.([]interface{})...)
			if err != nil {
				// 解析错误时直接返回
				log.Error().Str("funcName", funcName).Interface("arguments", arguments).
					Err(err).Msg("call function failed")
				return raw, errors.Wrap(code.CallFunctionError, err.Error())
			}
			log.Info().Str("funcName", funcName).Interface("arguments", arguments).
				Interface("output", result).Msg("call function success")

			if funcMatched[0] == raw {
				// raw_string is a function, e.g. "${add_one(3)}", return its eval value directly
				// 如果返回值和原始字符串相同，则直接返回
				return result, nil
			}

			// raw_string contains one or many functions, e.g. "abc${add_one(3)}def"
			matchStartPosition += len(funcMatched[0])
			parsedString += convertString(result)
			remainedString = raw[matchStartPosition:]
			log.Debug().
				Str("parsedString", parsedString).
				Int("matchStartPosition", matchStartPosition).
				Msg("[parseString] parse function")
			continue
		}

		// search variable like ${var} or $var
		varMatched := regexCompileVariable.FindStringSubmatch(remainedString)
		if len(varMatched) == 3 {
			var varName string
			if varMatched[1] != "" {
				varName = varMatched[1] // match ${var}
			} else {
				varName = varMatched[2] // match $var
			}
			varValue, ok := variablesMapping[varName]
			if !ok {
				return raw, errors.Wrap(code.VariableNotFound,
					fmt.Sprintf("variable %s not found", varName))
			}

			if fmt.Sprintf("${%s}", varName) == raw || fmt.Sprintf("$%s", varName) == raw {
				// raw string is a variable, $var or ${var}, return its value directly
				return varValue, nil
			}

			matchStartPosition += len(varMatched[0])
			parsedString += convertString(varValue)
			remainedString = raw[matchStartPosition:]
			log.Debug().
				Str("parsedString", parsedString).
				Int("matchStartPosition", matchStartPosition).
				Msg("[parseString] parse variable")
			continue
		}

		parsedString += remainedString
		break
	}

	return parsedString, nil
}

// callFunc calls function with arguments
// only support return at most one result value
func (p *Parser) callFunc(funcName string, arguments ...interface{}) (interface{}, error) {
	// call with plugin function
	if p.plugin != nil {
		if p.plugin.Has(funcName) {
			return p.plugin.Call(funcName, arguments...)
		}
		commonName := shared.ConvertCommonName(funcName)
		if p.plugin.Has(commonName) {
			return p.plugin.Call(commonName, arguments...)
		}
	}

	// get builtin function
	function, ok := builtin.Functions[funcName]
	if !ok {
		return nil, fmt.Errorf("function %s is not found", funcName)
	}
	fn := reflect.ValueOf(function)

	// call with builtin function
	return shared.CallFunc(fn, arguments...)
}

// merge two variables mapping, the first variables have higher priority
func mergeVariables(variables, overriddenVariables map[string]interface{}) map[string]interface{} {
	if overriddenVariables == nil {
		return variables
	}
	if variables == nil {
		return overriddenVariables
	}

	mergedVariables := make(map[string]interface{})
	for k, v := range overriddenVariables {
		mergedVariables[k] = v
	}
	for k, v := range variables {
		if fmt.Sprintf("${%s}", k) == v || fmt.Sprintf("$%s", k) == v {
			// e.g. {"base_url": "$base_url"}
			// or {"base_url": "${base_url}"}
			continue
		}

		mergedVariables[k] = v
	}
	return mergedVariables
}

// merge two map, the first map have higher priority
func mergeMap(m, overriddenMap map[string]string) map[string]string {
	if overriddenMap == nil {
		return m
	}
	if m == nil {
		return overriddenMap
	}

	mergedMap := make(map[string]string)
	for k, v := range overriddenMap {
		mergedMap[k] = v
	}
	for k, v := range m {
		mergedMap[k] = v
	}
	return mergedMap
}

// merge two validators slice, the first validators have higher priority
func mergeValidators(validators, overriddenValidators []interface{}) []interface{} {
	if validators == nil {
		return overriddenValidators
	}
	if overriddenValidators == nil {
		return validators
	}
	var mergedValidators []interface{}
	validators = append(validators, overriddenValidators...)
	for _, validator := range validators {
		flag := true
		for _, mergedValidator := range mergedValidators {
			if validator.(Validator).Check == mergedValidator.(Validator).Check {
				flag = false
				break
			}
		}
		if flag {
			mergedValidators = append(mergedValidators, validator)
		}
	}
	return mergedValidators
}

// merge two slices, the first slice have higher priority
func mergeSlices(slice, overriddenSlice []string) []string {
	if slice == nil {
		return overriddenSlice
	}
	if overriddenSlice == nil {
		return slice
	}

	for _, value := range overriddenSlice {
		if !builtin.Contains(slice, value) {
			slice = append(slice, value)
		}
	}
	return slice
}

var eval = goval.NewEvaluator()

// literalEval parse string to number if possible
func literalEval(raw string) (interface{}, error) {
	raw = strings.TrimSpace(raw)

	// return raw string if not number
	if !regexCompileNumber.Match([]byte(raw)) {
		return raw, nil
	}

	// eval string to number
	result, err := eval.Evaluate(raw, nil, nil)
	if err != nil {
		log.Error().Err(err).Msgf("[literalEval] eval %s failed", raw)
		return raw, err
	}
	return result, nil
}

// parseFunctionArguments 把字符串分割成变量、字符串和数字
func parseFunctionArguments(argsStr string) ([]interface{}, error) {
	// 调整args格式，用于去除字符串的首尾空白字符（空格、制表符、换行符等），并返回新的字符串
	argsStr = strings.TrimSpace(argsStr)
	if argsStr == "" {
		return []interface{}{}, nil
	}

	// split arguments by comma
	// 将字符串按照逗号分割成字符串数组
	args := strings.Split(argsStr, ",")
	// 创建一个长度为len(args)的interface{}类型的切片
	arguments := make([]interface{}, len(args))
	// 获取变量对应的值
	for index, arg := range args {
		// 调整arg格式，用于去除字符串的首尾空白字符（空格、制表符、换行符等），并返回新的字符串
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}

		// parse argument to number if possible
		// 将字符串转换为数字
		arg, err := literalEval(arg)
		if err != nil {
			return nil, err
		}
		arguments[index] = arg
	}

	return arguments, nil
}

func (p *Parser) ParseVariables(variables map[string]interface{}) (map[string]interface{}, error) {
	parsedVariables := make(map[string]interface{})
	var traverseRounds int

	for len(parsedVariables) != len(variables) {
		for varName, varValue := range variables {
			// skip parsed variables
			if _, ok := parsedVariables[varName]; ok {
				continue
			}

			// extract variables from current value
			extractVarsSet := extractVariables(varValue)

			// check if reference variable itself
			// e.g.
			// variables = {"token": "abc$token"}
			// variables = {"key": ["$key", 2]}
			if _, ok := extractVarsSet[varName]; ok {
				log.Error().Interface("variables", variables).Msg("[parseVariables] variable self reference error")
				return variables, errors.Wrap(code.ParseVariablesError,
					fmt.Sprintf("variable self reference: %v", varName))
			}

			// check if reference variable not in variables mapping
			// e.g.
			// {"varA": "123$varB", "varB": "456$varC"} => $varC not defined
			// {"varC": "${sum_two($a, $b)}"} => $a, $b not defined
			var undefinedVars []string
			for extractVar := range extractVarsSet {
				if _, ok := variables[extractVar]; !ok { // not in variables mapping
					undefinedVars = append(undefinedVars, extractVar)
				}
			}
			if len(undefinedVars) > 0 {
				log.Error().Interface("undefinedVars", undefinedVars).Msg("[parseVariables] variable not defined error")
				return variables, errors.Wrap(code.ParseVariablesError,
					fmt.Sprintf("variable not defined: %v", undefinedVars))
			}

			parsedValue, err := p.Parse(varValue, parsedVariables)
			if err != nil {
				continue
			}
			parsedVariables[varName] = parsedValue
		}
		traverseRounds += 1
		// check if circular reference exists
		if traverseRounds > len(variables) {
			log.Error().Msg("[parseVariables] circular reference error, break infinite loop!")
			return variables, errors.Wrap(code.ParseVariablesError, "circular reference")
		}
	}

	return parsedVariables, nil
}

type variableSet map[string]struct{}

func extractVariables(raw interface{}) variableSet {
	rawValue := reflect.ValueOf(raw)
	switch rawValue.Kind() {
	case reflect.String:
		return findallVariables(rawValue.String())
	case reflect.Slice:
		varSet := make(variableSet)
		for i := 0; i < rawValue.Len(); i++ {
			for extractVar := range extractVariables(rawValue.Index(i).Interface()) {
				varSet[extractVar] = struct{}{}
			}
		}
		return varSet
	case reflect.Map:
		varSet := make(variableSet)
		for _, key := range rawValue.MapKeys() {
			value := rawValue.MapIndex(key)
			for extractVar := range extractVariables(value.Interface()) {
				varSet[extractVar] = struct{}{}
			}
		}
		return varSet
	default:
		// other types, e.g. nil, int, float, bool
		return make(variableSet)
	}
}

func findallVariables(raw string) variableSet {
	matchStartPosition := 0
	remainedString := raw
	varSet := make(variableSet)

	for matchStartPosition < len(raw) {
		// locate $ char position
		startPosition := strings.Index(remainedString, "$")
		if startPosition == -1 { // no $ found
			return varSet
		}

		// found $, check if variable or function
		matchStartPosition += startPosition
		remainedString = remainedString[startPosition:]

		// Notice: notation priority
		// $$ > $var

		// search $$, use $$ to escape $ notation
		if strings.HasPrefix(remainedString, "$$") { // found $$
			matchStartPosition += 2
			remainedString = remainedString[2:]
			continue
		}

		// search variable like ${var} or $var
		varMatched := regexCompileVariable.FindStringSubmatch(remainedString)
		if len(varMatched) == 3 {
			var varName string
			if varMatched[1] != "" {
				varName = varMatched[1] // match ${var}
			} else {
				varName = varMatched[2] // match $var
			}
			varSet[varName] = struct{}{}

			matchStartPosition += len(varMatched[0])
			remainedString = raw[matchStartPosition:]
			continue
		}

		break
	}

	return varSet
}
