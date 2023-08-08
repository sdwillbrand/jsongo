package pkg

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

func ParseFromRune(data []rune) *JSValue {
	cursor := 0
	val := NewValue()
	for cursor < len(data) {
		c := data[cursor]
		if val.value != nil && !unicode.IsSpace(c) {
			// log.Fatal("End of file expected.")
			return nil
		}
		if unicode.IsSpace(c) {
			cursor++
			continue
		} else if c == '"' {
			val, cursor = parseString(data, cursor)
		} else if c == '-' || unicode.IsDigit(c) {
			val, cursor = parseNumber(data, cursor)
		} else if c == 'f' {
			val, cursor = parseFalse(data, cursor)
		} else if c == 't' {
			val, cursor = parseTrue(data, cursor)
		} else if c == 'n' {
			val, cursor = parseNull(data, cursor)
		} else if c == '[' {
			val, cursor = parseArray(data, cursor)
		} else if c == '{' {
			val, cursor = parseObject(data, cursor)
		} else {
			// log.Fatalf("Expected a JSON object, array or literal.\nError: %s", string(c))
			return nil
		}
		cursor++
	}
	return val
}

func ParseFromByte(data []byte) *JSValue {
	content := bytes.Runes(data)
	return ParseFromRune(content)
}

func ParseFromString(data string) *JSValue {
	content := bytes.Runes([]byte(data))
	return ParseFromRune(content)
}

func parseValue(data []rune, cursor int) (*JSValue, int) {
	cursor = skipWhitespace(data, cursor)
	var val *JSValue
	if cursor < len(data) && data[cursor] == '"' {
		val, cursor = parseString(data, cursor)
	} else if cursor < len(data) && (data[cursor] == '-' || unicode.IsDigit(data[cursor])) {
		val, cursor = parseNumber(data, cursor)
	} else if cursor < len(data) && data[cursor] == 'f' {
		val, cursor = parseFalse(data, cursor)
	} else if cursor < len(data) && data[cursor] == 't' {
		val, cursor = parseTrue(data, cursor)
	} else if cursor < len(data) && data[cursor] == 'n' {
		val, cursor = parseNull(data, cursor)
	} else if cursor < len(data) && data[cursor] == '[' {
		val, cursor = parseArray(data, cursor)
	} else if cursor < len(data) && data[cursor] == '{' {
		val, cursor = parseObject(data, cursor)
	} else {
		// log.Fatalf("Expected a JSON object, array or literal.\nError: %s", string(c))
		return nil, len(data)
	}
	return val, cursor
}

func parseArray(data []rune, cursor int) (*JSValue, int) {
	cursor++ // Skip [
	arr := NewArray()
	value := NewValue()
	value.kind = Array
	if cursor >= len(data) {
		return nil, len(data)
	}
	// Case 1: Empty Array
	if data[cursor] == ']' {
		value.value = arr
		return value, cursor
	}
	// Case 2: Single item
	var val *JSValue
	val, cursor = parseValue(data, cursor)
	if val == nil {
		return nil, len(data)
	}
	arr.objects = append(arr.objects, val)
	cursor++
	// Case 3: Multiple items
	for cursor < len(data) {
		cursor = skipWhitespace(data, cursor)
		if cursor < len(data) && data[cursor] == ',' ||
			cursor+1 < len(data) && data[cursor+1] == ',' {
			cursor++
			cursor = skipWhitespace(data, cursor)
			val, cursor = parseValue(data, cursor)
			if val == nil {
				return nil, len(data)
			}
			arr.objects = append(arr.objects, val)
		} else if cursor+1 < len(data) && unicode.IsSpace(data[cursor+1]) {
			cursor++
		} else if cursor < len(data) && data[cursor] == ']' || cursor+1 < len(data) && data[cursor+1] == ']' {
			cursor++
			break
		} else {
			return nil, len(data)
		}
	}
	value.value = arr
	return value, cursor
}

func parsePair(data []rune, cursor int) (string, *JSValue, int) {
	var key string
	if data[cursor] == '"' {
		var res *JSValue
		res, cursor = parseString(data, cursor)
		var err error
		key, err = res.GetString()
		if err != nil {
			return "", nil, len(data)
		}
	}
	cursor++ // Skip " of key
	cursor = skipWhitespace(data, cursor)
	if cursor < len(data) && data[cursor] == ':' {
		cursor++
	}
	if cursor+1 < len(data) && data[cursor+1] == ':' {
		cursor += 2
	}
	var val *JSValue
	val, cursor = parseValue(data, cursor)
	if val == nil {
		return "", nil, len(data)
	}
	return key, val, cursor
}

func parseObject(data []rune, cursor int) (*JSValue, int) {
	cursor++ // Skip {
	obj := NewObject()
	result := NewValue()
	result.kind = Object
	if cursor >= len(data) {
		return nil, len(data)
	}
	// Case 1: Empty Array
	if data[cursor] == '}' {
		result.value = obj
		return result, cursor
	}
	cursor = skipWhitespace(data, cursor)
	// Case 2: Single item
	var key string
	var val *JSValue
	key, val, cursor = parsePair(data, cursor)
	if val == nil {
		return nil, len(data)
	}
	obj.objects[key] = val
	cursor++
	for cursor < len(data) {
		cursor = skipWhitespace(data, cursor)
		if cursor+1 < len(data) && data[cursor+1] == ',' {
			cursor++
		}
		if cursor < len(data) && data[cursor] == ',' {
			cursor++
			cursor = skipWhitespace(data, cursor)
			var key string
			var val *JSValue
			key, val, cursor = parsePair(data, cursor)
			if val == nil {
				return nil, len(data)
			}
			obj.objects[key] = val
		} else if cursor < len(data) && data[cursor] == '}' {
			break
		} else if cursor+1 < len(data) && data[cursor+1] == '}' {
			cursor++
			break
		} else {
			return nil, len(data)
		}
	}
	result.value = obj
	return &JSValue{value: obj, kind: Object}, cursor
}

func parseFalse(data []rune, cursor int) (*JSValue, int) {
	if cursor+5 <= len(data) && bytes.Equal([]byte(string(data[cursor:cursor+5])), []byte("false")) {
		return &JSValue{value: false, kind: Boolean}, cursor + 4
	} else {
		return nil, len(data)
	}
}
func parseTrue(data []rune, cursor int) (*JSValue, int) {
	if cursor+4 <= len(data) && bytes.Equal([]byte(string(data[cursor:cursor+4])), []byte("true")) {
		return &JSValue{value: true, kind: Boolean}, cursor + 3
	} else {
		return nil, len(data)
	}
}
func parseNull(data []rune, cursor int) (*JSValue, int) {
	if cursor+4 < len(data) && bytes.Equal([]byte(string(data[cursor:cursor+4])), []byte("null")) {
		return &JSValue{value: nil, kind: Null}, cursor + 3
	} else {
		return nil, len(data)
	}
}

func parseDigit(data []rune, cursor int) (string, int) {
	r := ""
	j := 0
	for ; cursor+j < len(data) && unicode.IsDigit(rune(data[cursor+j])); j++ {
		r += string(data[cursor+j])
	}
	return r, cursor + j - 1
}

func parseNumber(data []rune, cursor int) (*JSValue, int) {
	sign := ""
	base := ""
	frac := ""
	exp := ""
	expSign := ""
	if rune(data[cursor]) == '-' {
		sign = "-"
		cursor++
	}
	base, cursor = parseDigit(data, cursor)
	if cursor+1 < len(data) && data[cursor+1] == '.' {
		cursor = cursor + 2
		frac, cursor = parseDigit(data, cursor)
	}
	if cursor+1 < len(data) && (data[cursor+1] == 'e' || data[cursor+1] == 'E') {
		cursor += 2
		if cursor >= len(data) {
			// log.Fatal("Unexpected end of number.")
			return nil, len(data)
		}
		if cursor < len(data) && data[cursor] == '-' {
			expSign = "-"
			cursor++
		} else if cursor < len(data) && data[cursor] == '+' {
			cursor++
		} else if !unicode.IsDigit(rune(data[cursor])) {
			// log.Fatal("Unexpected end of number.")
			return nil, len(data)
		}
		exp, cursor = parseDigit(data, cursor)
		if exp == "" {
			return nil, len(data)
		}
	}
	if frac == "" {
		frac = "0"
	}
	if exp == "" {
		exp = "0"
	}
	result := fmt.Sprintf("%s%s.%se%s%s", sign, base, frac, expSign, exp)
	value, _ := strconv.ParseFloat(result, 64)
	return &JSValue{value: value, kind: Number}, cursor
}

func parseString(data []rune, cursor int) (*JSValue, int) {
	cursor++ // Skip first "
	var value string
	if cursor >= len(data) {
		return nil, len(data)
	}
	token := data[cursor]
	for cursor < len(data) && token != '"' {
		if token == '\\' {
			cursor++
			token = data[cursor]
			switch token {
			case '"':
				value += string('"')
				cursor++
			case 'b':
				value += string('\b')
				cursor++
			case 'f':
				value += string('\f')
				cursor++
			case 'n':
				value += string('\n')
				cursor++
			case 'r':
				value += string('\r')
				cursor++
			case 't':
				value += string('\t')
				cursor++
			case '\\':
				value += string('\\')
				cursor++
			case '/':
				value += string('/')
				cursor++
			case 'u':
				hex := data[cursor-1 : cursor+5]
				if !isHex(data[cursor+1 : cursor+5]) {
					// log.Fatal("Error: Invalid unicode sequence in string")
					return nil, len(data)
				}
				r, _, _, err := strconv.UnquoteChar(string(hex), 0)
				if err == nil {
					value += string(r)
				} else {
					value += string(hex)
				}
				cursor += 5
			default:
				// log.Fatal("Error: Invalid escape character")
				return nil, len(data)
			}
		} else {
			value += string(token)
			cursor++
		}
		token = data[cursor]
	}
	return &JSValue{value: value, kind: String}, cursor
}

func isHex(s []rune) bool {
	for _, b := range s {
		if !(b >= '0' && b <= '9' || b >= 'a' && b <= 'f' || b >= 'A' && b <= 'F') {
			return false
		}
	}
	return true
}

func skipWhitespace(data []rune, cursor int) int {
	for cursor < len(data) && unicode.IsSpace(data[cursor]) {
		cursor++
	}
	return cursor
}
