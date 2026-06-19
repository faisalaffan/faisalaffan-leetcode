package main

// LeetCode #1106: Parsing A Boolean Expression
// https://leetcode.com/problems/parsing-a-boolean-expression/
// Difficulty: Hard
//
// Recursive descent parser for boolean expressions:
//   't' → true
//   'f' → false
//   '!(expr)' → NOT
//   '&(expr,expr,...)' → AND
//   '|(expr,expr,...)' → OR

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("&(|(f))"))
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
}

func parseBoolExpr(expression string) bool {
	idx := 0

	var parse func() bool
	parse = func() bool {
		ch := expression[idx]
		idx++

		if ch == 't' {
			return true
		}
		if ch == 'f' {
			return false
		}

		// ch is '!', '&', or '|'
		idx++ // skip '('

		var result bool
		if ch == '!' {
			result = !parse()
		} else {
			result = parse()
			for idx < len(expression) && expression[idx] == ',' {
				idx++ // skip ','
				val := parse()
				if ch == '&' {
					result = result && val
				} else { // '|'
					result = result || val
				}
			}
		}

		idx++ // skip ')'
		return result
	}

	return parse()
}
