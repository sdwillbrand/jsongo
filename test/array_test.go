package test

import (
	jq "jsongo/pkg"
	"testing"
)

type ArraySuit struct {
	content []rune
	result  jq.JSArray
}

var y_ArraySuites = [...](ArraySuit){
	ArraySuit{content: []rune("[]"), result: *jq.NewArray()},
	ArraySuit{content: []rune("[ true ]"), result: *jq.NewArray().Append(jq.NewBoolean(true))},
	ArraySuit{content: []rune("[ false ]"), result: *jq.NewArray().Append(jq.NewBoolean(false))},
	ArraySuit{content: []rune("[ null ]"), result: *jq.NewArray().Append(jq.NewValue())},
	ArraySuit{content: []rune("[ 1 ]"), result: *jq.NewArray().Append(jq.NewNumber(1))},
}

var n_ArraySuites = [...](ArraySuit){
	ArraySuit{content: []rune("]")},
	ArraySuit{content: []rune("[")},
	ArraySuit{content: []rune("[1 true]")},
}

func TestParseArray(t *testing.T) {
	for i, arraySuite := range y_ArraySuites {
		data := arraySuite.content
		arr := jq.ParseFromRune(data)
		if arr == nil {
			t.Fatalf("Test %d: Got nil, wanted %s", i, arraySuite.result)
		}
		result, _ := arr.GetArray()
		if result == nil {
			t.Fatalf("Test %d, Got nil for array", i)
		}
		item := (*result).Get(0)
		other := arraySuite.result.Get(0)
		if item != nil && other != nil && !item.Equals(other) {
			t.Fatalf("Test %d: Got %s, wanted %s", i, *result, arraySuite.result)
		}
	}
	for i, arraySuite := range n_ArraySuites {
		data := arraySuite.content
		arr := jq.ParseFromRune(data)
		if arr != nil {
			t.Fatalf("Test %d: Wanted error", i)
		}
	}
}
