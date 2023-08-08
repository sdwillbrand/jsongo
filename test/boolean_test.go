package test

import (
	"testing"

	"github.com/countersoda/jsongo"
)

type BoolSuit struct {
	content []rune
	result  bool
}

var y_BoolSuites = [...](BoolSuit){
	BoolSuit{content: []rune("false"), result: false},
	BoolSuit{content: []rune("true"), result: true},
}

var n_BoolSuites = [...](BoolSuit){
	BoolSuit{content: []rune("fals")},
	BoolSuit{content: []rune("tru")},
}

func TestParseBoolean(t *testing.T) {
	for _, boolSuite := range y_BoolSuites {
		data := boolSuite.content
		b := jsongo.ParseFromRune(data)
		if b == nil {
			t.Fatalf("Got nil, wanted %t", boolSuite.result)
		}
		result, _ := b.Bool()
		if result != boolSuite.result {
			t.Fatalf("Got %t, wanted %t", result, boolSuite.result)
		}
	}
	for _, boolSuite := range n_BoolSuites {
		data := boolSuite.content
		b := jsongo.ParseFromRune(data)
		if b != nil {
			t.Fatalf("Wanted error")
		}
	}
}
