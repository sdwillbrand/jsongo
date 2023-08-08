package jsongo

import (
	"fmt"
	"strconv"
)

type JSType int64

const (
	String JSType = iota + 1
	Number
	Boolean
	Null
	Array
	Object
)

type JSValue struct {
	value interface{}
	kind  JSType
}

func NewArray() []*JSValue {
	return make([]*JSValue, 0, 10)
}

func NewValue() *JSValue {
	return &JSValue{kind: Null, value: nil}
}

func NewBoolean(b bool) *JSValue {
	return &JSValue{kind: Boolean, value: b}
}

func NewNumber(n float64) *JSValue {
	return &JSValue{kind: Number, value: n}
}

func NewObject() map[string]*JSValue {
	return make(map[string]*JSValue, 0)
}

func (val *JSValue) Equals(other *JSValue) bool {
	if val.kind != other.kind {
		return false
	} else if val.kind == Array || val.kind == Object {
		return false
	} else {
		if val.kind == Boolean {
			v1, _ := val.Bool()
			v2, _ := other.Bool()
			return v1 == v2
		} else if val.kind == Null {
			return true
		} else if val.kind == Number {
			v1, _ := val.Float()
			v2, _ := other.Float()
			return v1 == v2
		} else if val.kind == String {
			v1, _ := val.String()
			v2, _ := other.String()
			return v1 == v2
		}
		return false
	}
}

func (val *JSValue) ToString() (string, error) {
	result := ""
	switch val.kind {
	case Array:
		v, ok := val.value.([]*JSValue)
		if !ok {
			return "", fmt.Errorf("%s", "Cannot convert to []*JSValue in value")
		}
		for index, i := range v {
			s, _ := i.ToString()
			result += s
			if index+1 != len(v) {
				result += ", "
			}
		}
	case Object:
		v, ok := val.value.(map[string]*JSValue)
		if !ok {
			return "", fmt.Errorf("%s", "Cannot convert to map[string]*JSValue struct")
		}
		result += "{"
		amount := len(v)
		j := 0
		for k, v := range v {
			s, _ := v.ToString()
			result += fmt.Sprintf("\"%s\" : %s", k, s)
			if j+1 != amount {
				result += ", "
			}
			j++
		}
		result += "}"
	case String:
		v, ok := val.value.(string)
		if !ok {
			return "", fmt.Errorf("%s", "Cannot convert to string")
		}
		result += "\"" + string(v) + "\""
	case Boolean:
		v, ok := val.value.(bool)
		if !ok {
			return "", fmt.Errorf("%s", "Cannot convert to bool")
		}
		if v {
			result += "true"
		} else {
			result += "false"
		}
	case Null:
		result += "null"
	case Number:
		v, ok := val.value.(float64)
		if !ok {
			return "", fmt.Errorf("%s", "Cannot convert to float")
		}
		result += strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return "", fmt.Errorf("%s", "Cannot convert to string")
	}
	return string(result), nil
}

func (v *JSValue) Float() (float64, error) {
	res, ok := v.value.(float64)
	if !ok {
		return 0, fmt.Errorf("%s", "Error: Not a number value")
	} else {
		return res, nil
	}
}

func (v *JSValue) Int() (int64, error) {
	res, ok := v.value.(float64)
	if !ok {
		return 0, fmt.Errorf("%s", "Error: Not a number value")
	} else {
		return int64(res), nil
	}
}

func (v *JSValue) String() (string, error) {
	res, ok := v.value.(string)
	if !ok {
		return "", fmt.Errorf("%s", "Error: Not a string value")
	} else {
		return res, nil
	}
}

func (v *JSValue) Bool() (bool, error) {
	res, ok := v.value.(bool)
	if !ok {
		return false, fmt.Errorf("%s", "Error: Not a bool value")
	} else {
		return res, nil
	}
}

func (v *JSValue) Array() ([]*JSValue, error) {
	res, ok := v.value.([]*JSValue)
	if !ok {
		return nil, fmt.Errorf("%s", "Error: Not an array value")
	} else {
		return res, nil
	}
}

func (v *JSValue) Object() (map[string]*JSValue, error) {
	res, ok := v.value.(map[string]*JSValue)
	if !ok {
		return nil, fmt.Errorf("%s", "Error: Not an object value")
	} else {
		return res, nil
	}
}

func (v *JSValue) Type() JSType {
	return v.kind
}
