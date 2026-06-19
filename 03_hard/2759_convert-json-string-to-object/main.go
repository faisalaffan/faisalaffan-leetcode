package main

// LeetCode #2759: Convert JSON String to Object
// https://leetcode.com/problems/convert-json-string-to-object/
// Difficulty: Hard [Paid]
//
// Approach: Recursive descent parser. Supports objects, arrays,
// strings, numbers, booleans, and null.

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type jsonValue interface{}

func parseJSON(s string) (jsonValue, int) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil, 0
	}

	switch {
	case s[0] == '{':
		return parseObject(s)
	case s[0] == '[':
		return parseArray(s)
	case s[0] == '"':
		return parseString(s)
	case s[0] == 't' || s[0] == 'f':
		return parseBool(s)
	case s[0] == 'n':
		return parseNull(s)
	default:
		return parseNumber(s)
	}
}

func parseObject(s string) (map[string]jsonValue, int) {
	obj := make(map[string]jsonValue)
	pos := 1 // skip '{'
	if pos < len(s) && s[pos] == '}' {
		return obj, pos + 1
	}
	for pos < len(s) {
		for pos < len(s) && (s[pos] == ' ' || s[pos] == ',' || s[pos] == '\n' || s[pos] == '\t' || s[pos] == '\r') {
			pos++
		}
		if pos >= len(s) || s[pos] == '}' {
			return obj, pos + 1
		}
		key, n := parseString(s[pos:])
		pos += n
		for pos < len(s) && s[pos] == ' ' {
			pos++
		}
		pos++ // skip ':'
		for pos < len(s) && s[pos] == ' ' {
			pos++
		}
		val, n := parseJSON(s[pos:])
		obj[key] = val
		pos += n
	}
	return obj, pos
}

func parseArray(s string) ([]jsonValue, int) {
	arr := make([]jsonValue, 0)
	pos := 1 // skip '['
	if pos < len(s) && s[pos] == ']' {
		return arr, pos + 1
	}
	for pos < len(s) {
		for pos < len(s) && (s[pos] == ' ' || s[pos] == ',' || s[pos] == '\n' || s[pos] == '\t' || s[pos] == '\r') {
			pos++
		}
		if pos >= len(s) || s[pos] == ']' {
			return arr, pos + 1
		}
		val, n := parseJSON(s[pos:])
		arr = append(arr, val)
		pos += n
	}
	return arr, pos
}

func parseString(s string) (string, int) {
	pos := 1 // skip opening '"'
	var sb strings.Builder
	for pos < len(s) {
		if s[pos] == '\\' {
			pos++
			switch s[pos] {
			case '"', '\\', '/':
				sb.WriteByte(s[pos])
			case 'b':
				sb.WriteByte('\b')
			case 'f':
				sb.WriteByte('\f')
			case 'n':
				sb.WriteByte('\n')
			case 'r':
				sb.WriteByte('\r')
			case 't':
				sb.WriteByte('\t')
			case 'u':
				hex := s[pos+1 : pos+5]
				val, _ := strconv.ParseInt(hex, 16, 32)
				sb.WriteRune(rune(val))
				pos += 4
			}
			pos++
		} else if s[pos] == '"' {
			return sb.String(), pos + 1
		} else {
			sb.WriteByte(s[pos])
			pos++
		}
	}
	return sb.String(), pos
}

func parseNumber(s string) (float64, int) {
	pos := 0
	for pos < len(s) && (unicode.IsDigit(rune(s[pos])) || s[pos] == '.' || s[pos] == '-' || s[pos] == '+' || s[pos] == 'e' || s[pos] == 'E') {
		pos++
	}
	val, _ := strconv.ParseFloat(s[:pos], 64)
	return val, pos
}

func parseBool(s string) (bool, int) {
	if s[0] == 't' {
		return true, 4
	}
	return false, 5
}

func parseNull(s string) (any, int) {
	return nil, 4
}

func jsonStringify(v jsonValue) string {
	switch val := v.(type) {
	case nil:
		return "null"
	case bool:
		if val {
			return "true"
		}
		return "false"
	case float64:
		if val == float64(int(val)) {
			return strconv.Itoa(int(val))
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return `"` + val + `"`
	case []jsonValue:
		parts := make([]string, len(val))
		for i, item := range val {
			parts[i] = jsonStringify(item)
		}
		return "[" + strings.Join(parts, ",") + "]"
	case map[string]jsonValue:
		parts := make([]string, 0, len(val))
		for k, item := range val {
			parts = append(parts, `"`+k+`":`+jsonStringify(item))
		}
		return "{" + strings.Join(parts, ",") + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func convertJSONStringToObject(s string) jsonValue {
	v, _ := parseJSON(s)
	return v
}

func main() {
	// Object with array
	fmt.Println(jsonStringify(convertJSONStringToObject(`{"a":1,"b":[2,3]}`)))
	// Boolean
	fmt.Println(jsonStringify(convertJSONStringToObject(`true`)))
	// Array with mixed types
	fmt.Println(jsonStringify(convertJSONStringToObject(`[1,"hello",null]`)))
	// Nested object
	fmt.Println(jsonStringify(convertJSONStringToObject(`{"nested":{"key":"value"}}`)))
	// Number
	fmt.Println(jsonStringify(convertJSONStringToObject(`42`)))
	// String
	fmt.Println(jsonStringify(convertJSONStringToObject(`"hello, world"`)))
}
