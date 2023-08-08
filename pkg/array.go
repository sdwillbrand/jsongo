package pkg

import "log"

type JSArray struct {
	objects []*JSValue
}

func NewArray() *JSArray {
	return &JSArray{objects: make([]*JSValue, 0, 10)}
}

func (arr *JSArray) Append(val *JSValue) *JSArray {
	arr.objects = append(arr.objects, val)
	return arr
}

func (arr *JSArray) String() string {
	result := "["
	for i, o := range arr.objects {
		switch o.kind {
		case Array:
			v, ok := o.value.(*JSArray)
			if !ok {
				log.Fatal("Cannot convert to JSArray struct in array")
			}
			result += v.String()
		case Object:
			v, ok := o.value.(*JSObject)
			if !ok {
				log.Fatal("Cannot convert to JSObject struct")
			}
			result += v.String()
		default:
			result += o.String()
		}
		if i < len(arr.objects)-1 {
			result += ","
		}
	}
	result += "]"
	return string(result)
}

func (arr *JSArray) Get(i int) *JSValue {
	if i < 0 || i >= len(arr.objects) {
		return nil
	} else {
		return arr.objects[i]
	}
}
