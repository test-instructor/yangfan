package customType

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type TypeArgsMap []map[string]interface{}

// Scan Scanner
func (argsMap *TypeArgsMap) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	err := json.Unmarshal(b, &argsMap)
	return err
}

// Value Valuer
func (argsMap TypeArgsMap) Value() (driver.Value, error) {
	if argsMap == nil {
		return nil, nil
	}
	argsStr, _ := json.Marshal(argsMap)
	return argsStr, nil
}

type TypeArgs []string

func (a TypeArgs) Value() (value driver.Value, err error) {
	if len(a) > 0 {
		var str string = a[0]
		for _, v := range a[1:] {
			str += "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}

func (a *TypeArgs) Scan(value interface{}) error {
	str, ok := value.([]byte)
	strs := string(str)
	if !ok {
		return errors.New("数据格式无法解析")
	}
	tempStr := ""
	for k, v := range strs {
		if k < len(strs)-1 {
			if string(v) == "," && string(strs[k+1]) == "$" {
				*a = append(*a, tempStr)
				tempStr = ""
				fmt.Println(string(v))

			} else {
				tempStr += string(v)
			}
		} else {
			tempStr += string(v)
			*a = append(*a, tempStr)
			break
		}
	}

	return nil
}

type TypeJson []byte

func (j TypeJson) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *TypeJson) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], string(s)...)
	return nil
}
func (m TypeJson) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *TypeJson) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j TypeJson) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j TypeJson) Equals(j1 TypeJson) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
