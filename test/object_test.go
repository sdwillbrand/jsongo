package test

import (
	jq "jsongo/pkg"
	"testing"
)

type ObjectSuit struct {
	content []rune
	result  jq.JSObject
}

var y_ObjectSuites = [...](ObjectSuit){
	ObjectSuit{content: []rune("{}"), result: *jq.NewObject()},
	ObjectSuit{content: []rune("{ \"a\" : true }"), result: *jq.NewObject().Insert("a", jq.NewBoolean(true))},
	ObjectSuit{content: []rune("{ \"a\" : false }"), result: *jq.NewObject().Insert("a", jq.NewBoolean(false))},
	ObjectSuit{content: []rune("{ \"a\" : null }"), result: *jq.NewObject().Insert("a", jq.NewValue())},
	ObjectSuit{content: []rune("{ \"a\" : 1 }"), result: *jq.NewObject().Insert("a", jq.NewNumber(1))},
}

var n_ObjectSuites = [...](ObjectSuit){
	ObjectSuit{content: []rune("}")},
	ObjectSuit{content: []rune("{")},
}

func TestParseObject(t *testing.T) {
	for i, arraySuite := range y_ObjectSuites {
		data := arraySuite.content
		arr := jq.ParseFromRune(data)
		if arr == nil {
			t.Fatalf("Test %d: Got nil, wanted %s", i, arraySuite.result)
		}
		result, _ := arr.GetObject()
		if result == nil {
			t.Fatalf("Test %d, Got nil for array", i)
		}
		item := (*result).Get("a")
		other := arraySuite.result.Get("a")
		if item != nil && other != nil && !item.Equals(other) {
			t.Fatalf("Test %d: Got %s, wanted %s", i, *result, arraySuite.result)
		}
	}
	for i, arraySuite := range n_ObjectSuites {
		data := arraySuite.content
		arr := jq.ParseFromRune(data)
		if arr != nil {
			t.Fatalf("Test %d: Wanted error", i)
		}
	}
}
