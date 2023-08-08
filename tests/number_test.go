package test

import (
	"testing"

	"github.com/countersoda/jsongo"
)

type NumberSuit struct {
	content []rune
	result  float64
}

var y_NumberSuites = [...](NumberSuit){
	NumberSuit{content: []rune("12.0"), result: 12.0},
	NumberSuit{content: []rune("-12.0"), result: -12.0},
	NumberSuit{content: []rune("12.0e1"), result: 120.0},
	NumberSuit{content: []rune("12.0e-1"), result: 1.20},
	NumberSuit{content: []rune("0e0"), result: 0.0},
	NumberSuit{content: []rune("123.456e-789"), result: 0},
}

var n_NumberSuites = [...](NumberSuit){
	NumberSuit{content: []rune("0e")},
	NumberSuit{content: []rune("1e+")},
	NumberSuit{content: []rune("1eE2")},
}

func TestParseNumber(t *testing.T) {
	for i, numberSuite := range y_NumberSuites {
		data := numberSuite.content
		number := jsongo.ParseFromRune(data)
		if number == nil {
			t.Fatalf("Test: %d, Got nil, wanted %f", i, numberSuite.result)
		}
		result, _ := number.Float()
		if result != numberSuite.result {
			t.Fatalf("Test: %d, Got %f, wanted %f", i, result, numberSuite.result)
		}
	}
	for i, numberSuite := range n_NumberSuites {
		data := numberSuite.content
		number := jsongo.ParseFromRune(data)
		if number != nil {
			t.Fatalf("Test: %d, Wanted error", i)
		}
	}
}
