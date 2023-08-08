package pkg

import (
	"fmt"
	"log"
	"sort"
)

type JSObject struct {
	objects map[string]*JSValue
}

func NewObject() *JSObject {
	return &JSObject{objects: make(map[string]*JSValue)}
}

func (obj *JSObject) String() string {
	result := "{"
	keys := make([]string, 0, 10)
	for key := range obj.objects {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		o := obj.objects[key]
		switch o.kind {
		case Array:
			v, ok := o.value.(*JSArray)
			if !ok {
				log.Fatal("Cannot convert to JSArray struct in object")
			}
			result += fmt.Sprintf("\"%s\": %s", key, v.String())
		case Object:
			v, ok := o.value.(*JSObject)
			if !ok {
				log.Fatal("Cannot convert to JSObject struct")
			}
			result += fmt.Sprintf("\"%s\": %s", key, v.String())
		default:
			result += fmt.Sprintf("\"%s\": %s", key, o.String())
		}
		if i < len(keys)-1 {
			result += ","
		}

	}
	result += "}"
	return string(result)
}

func (obj *JSObject) Get(key string) *JSValue {
	return obj.objects[key]
}

func (obj *JSObject) Insert(key string, value *JSValue) *JSObject {
	obj.objects[key] = value
	return obj
}

func (obj *JSObject) Delete(key string) *JSObject {
	obj.objects[key] = nil
	return obj
}
