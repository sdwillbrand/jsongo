package pkg

import (
	"fmt"
	"log"
	"strconv"
)

type JSValue struct {
	value interface{}
	kind  JSType
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

func (val *JSValue) Equals(other *JSValue) bool {
	if val.kind != other.kind {
		return false
	} else if val.kind == Array || val.kind == Object {
		return false
	} else {
		if val.kind == Boolean {
			v1, _ := val.GetBool()
			v2, _ := other.GetBool()
			return v1 == v2
		} else if val.kind == Null {
			return true
		} else if val.kind == Number {
			v1, _ := val.GetFloat()
			v2, _ := other.GetFloat()
			return v1 == v2
		} else if val.kind == String {
			v1, _ := val.GetString()
			v2, _ := other.GetString()
			return v1 == v2
		}
		return false
	}
}

func (val *JSValue) String() string {
	result := ""
	switch val.kind {
	case Array:
		v, ok := val.value.(*JSArray)
		if !ok {
			log.Fatal("Cannot convert to JSArray struct in value")
		}
		result += v.String()
	case Object:
		v, ok := val.value.(*JSObject)
		if !ok {
			log.Fatal("Cannot convert to JSObject struct")
		}
		result += v.String()
	case String:
		v, ok := val.value.(string)
		if !ok {
			log.Fatal("Cannot convert to string")
		}
		result += "\"" + string(v) + "\""
	case Boolean:
		v, ok := val.value.(bool)
		if !ok {
			log.Fatal("Cannot convert to bool")
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
			log.Fatal("Cannot convert to float")
		}
		result += strconv.FormatFloat(v, 'f', -1, 64)
	default:
		log.Fatal("Cannot convert to string")
	}
	return string(result)
}

func (v *JSValue) GetFloat() (float64, error) {
	res, ok := v.value.(float64)
	if !ok {
		return 0, fmt.Errorf("%s", "Error: Not a number value")
	} else {
		return res, nil
	}
}

func (v *JSValue) GetString() (string, error) {
	res, ok := v.value.(string)
	if !ok {
		return "", fmt.Errorf("%s", "Error: Not a string value")
	} else {
		return res, nil
	}
}

func (v *JSValue) GetBool() (bool, error) {
	res, ok := v.value.(bool)
	if !ok {
		return false, fmt.Errorf("%s", "Error: Not a bool value")
	} else {
		return res, nil
	}
}

func (v *JSValue) GetArray() (*JSArray, error) {
	res, ok := v.value.(*JSArray)
	if !ok {
		return nil, fmt.Errorf("%s", "Error: Not an array value")
	} else {
		return res, nil
	}
}

func (v *JSValue) GetObject() (*JSObject, error) {
	res, ok := v.value.(*JSObject)
	if !ok {
		return nil, fmt.Errorf("%s", "Error: Not an object value")
	} else {
		return res, nil
	}
}
