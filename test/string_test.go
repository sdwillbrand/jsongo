package test

import (
	"testing"

	"github.com/countersoda/jsongo"
)

type StringSuit struct {
	content []rune
	result  string
}

var y_StringSuites = [...](StringSuit){
	StringSuit{content: []rune("\"\""), result: ""},
	StringSuit{content: []rune("\"Hello\""), result: "Hello"},
	StringSuit{content: []rune{'"', '\\', 'u', '1', '2', '1', '2', '"'}, result: "\u1212"},
	StringSuit{content: []rune{'"', 'a', '\\', 'n', '"'}, result: "a\n"},
	StringSuit{content: []rune("\"\u1212\""), result: "\u1212"},
	StringSuit{content: []rune("\"\b\""), result: "\b"},
}

var n_StringSuites = [...](StringSuit){
	StringSuit{content: []rune("\"")},
	StringSuit{content: []rune("\"\\uqqqq\"")},
	StringSuit{content: []rune("\"\\u00a\"")},
}

func TestParseString(t *testing.T) {
	for i, stringSuite := range y_StringSuites {
		data := stringSuite.content
		str := jsongo.ParseFromRune(data)
		if str == nil {
			t.Fatalf("Test %d: Got nil, wanted %s", i, stringSuite.result)
		}
		result, _ := str.String()
		if result != stringSuite.result {
			t.Fatalf("Test %d: Got %s, wanted %s", i, result, stringSuite.result)
		}
	}
	for i, stringSuite := range n_StringSuites {
		data := stringSuite.content
		str := jsongo.ParseFromRune(data)
		if str != nil {
			t.Fatalf("Test %d: Wanted error", i)
		}
	}
}
