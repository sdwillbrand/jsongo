package test

import (
	"testing"

	"github.com/countersoda/jsongo"
)

type ArraySuit struct {
	content []rune
	result  []*jsongo.JSValue
}

var y_ArraySuites = [...](ArraySuit){
	// ArraySuit{content: []rune("[]"), result: jsongo.NewArray()},
	ArraySuit{content: []rune("[ true ]"), result: append(jsongo.NewArray(), jsongo.NewBoolean(true))},
	ArraySuit{content: []rune("[ false ]"), result: append(jsongo.NewArray(), jsongo.NewBoolean(false))},
	ArraySuit{content: []rune("[ null ]"), result: append(jsongo.NewArray(), jsongo.NewValue())},
	ArraySuit{content: []rune("[ 1 ]"), result: append(jsongo.NewArray(), jsongo.NewNumber(1))},
}

var n_ArraySuites = [...](ArraySuit){
	ArraySuit{content: []rune("]")},
	ArraySuit{content: []rune("[")},
	ArraySuit{content: []rune("[1 true]")},
}

func TestParseArray(t *testing.T) {
	for i, arraySuite := range y_ArraySuites {
		data := arraySuite.content
		arr := jsongo.ParseFromRune(data)
		if arr == nil {
			t.Fatalf("Test %d: Got nil, wanted %v", i, arraySuite.result)
		}
		result, err := arr.Array()
		if err != nil {
			t.Fatalf("Test %d, Got nil for array", i)
		}
		item := result[0]
		other := arraySuite.result[0]
		if item != nil && other != nil && !item.Equals(other) {
			t.Fatalf("Test %d: Got %v, wanted %v", i, result, arraySuite.result)
		}
	}
	for i, arraySuite := range n_ArraySuites {
		data := arraySuite.content
		arr := jsongo.ParseFromRune(data)
		if arr != nil {
			t.Fatalf("Test %d: Wanted error", i)
		}
	}
}
