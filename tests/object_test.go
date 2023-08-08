package test

import (
	"testing"

	"github.com/countersoda/jsongo"
)

type ObjectSuit struct {
	content []rune
	result  map[string]*jsongo.JSValue
}

func insert(m map[string]*jsongo.JSValue, k string, v *jsongo.JSValue) map[string]*jsongo.JSValue {
	m[k] = v
	return m
}

var y_ObjectSuites = [...](ObjectSuit){
	ObjectSuit{content: []rune("{}"), result: jsongo.NewObject()},
	ObjectSuit{content: []rune("{ \"a\" : true }"), result: insert(jsongo.NewObject(), "a", jsongo.NewBoolean(true))},
	ObjectSuit{content: []rune("{ \"a\" : false }"), result: insert(jsongo.NewObject(), "a", jsongo.NewBoolean(false))},
	ObjectSuit{content: []rune("{ \"a\" : null }"), result: insert(jsongo.NewObject(), "a", jsongo.NewValue())},
	ObjectSuit{content: []rune("{ \"a\" : 1 }"), result: insert(jsongo.NewObject(), "a", jsongo.NewNumber(1))},
}

var n_ObjectSuites = [...](ObjectSuit){
	ObjectSuit{content: []rune("}")},
	ObjectSuit{content: []rune("{")},
}

func TestParseObject(t *testing.T) {
	for i, arraySuite := range y_ObjectSuites {
		data := arraySuite.content
		arr := jsongo.ParseFromRune(data)
		if arr == nil {
			t.Fatalf("Test %d: Got nil, wanted %v", i, arraySuite.result)
		}
		result, err := arr.Object()
		if err != nil {
			t.Fatalf("Test %d, Got nil for array", i)
		}
		item := result["a"]
		other := arraySuite.result["a"]
		if item != nil && other != nil && !item.Equals(other) {
			t.Fatalf("Test %d: Got %v, wanted %v", i, result, arraySuite.result)
		}
	}
	for i, arraySuite := range n_ObjectSuites {
		data := arraySuite.content
		arr := jsongo.ParseFromRune(data)
		if arr != nil {
			t.Fatalf("Test %d: Wanted error", i)
		}
	}
}
