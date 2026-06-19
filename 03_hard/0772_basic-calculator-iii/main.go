package main

// LeetCode #772: Basic Calculator III
// https://leetcode.com/problems/basic-calculator-iii/
// Difficulty: Hard [Paid]
//
// Expression with +, -, *, /, parentheses, and non-negative integers.
// Evaluate and return the result as an integer (division truncates toward zero).
//
// Approach: Two-stack (recursive descent)
// Use a single-pass parser with precedence handling.

import "fmt"

func main() {
	fmt.Println(calculate("2*(5+5*2)/3+(6/2+8)"))    // 21
	fmt.Println(calculate("1+1"))                      // 2
	fmt.Println(calculate("6-4/2"))                    // 4
	fmt.Println(calculate("2*(5+5*2)/3"))              // 10
	fmt.Println(calculate("(2+6*3+5-(3*14/7+2)*5)+3")) // -12
}

func calculate(s string) int {
	// We'll use a recursive approach:
	// parseExpr handles +, -
	// parseTerm handles *, /
	// parseFactor handles numbers, parentheses, and unary minus
	idx := 0
	return parseExpr(s, &idx)
}

// parseExpr parses addition and subtraction
func parseExpr(s string, idx *int) int {
	result := parseTerm(s, idx)
	for *idx < len(s) {
		ch := s[*idx]
		if ch == '+' || ch == '-' {
			*idx++
			right := parseTerm(s, idx)
			if ch == '+' {
				result += right
			} else {
				result -= right
			}
		} else {
			break
		}
	}
	return result
}

// parseTerm parses multiplication and division
func parseTerm(s string, idx *int) int {
	result := parseFactor(s, idx)
	for *idx < len(s) {
		ch := s[*idx]
		if ch == '*' || ch == '/' {
			*idx++
			right := parseFactor(s, idx)
			if ch == '*' {
				result *= right
			} else {
				result /= right
			}
		} else {
			break
		}
	}
	return result
}

// parseFactor parses numbers, parentheses, and unary operators
func parseFactor(s string, idx *int) int {
	// Skip spaces
	for *idx < len(s) && s[*idx] == ' ' {
		*idx++
	}

	// Handle unary minus
	sign := 1
	if *idx < len(s) && s[*idx] == '-' {
		sign = -1
		*idx++
	}
	if *idx < len(s) && s[*idx] == '+' {
		*idx++
	}

	// Skip spaces after unary operator
	for *idx < len(s) && s[*idx] == ' ' {
		*idx++
	}

	var result int
	if *idx < len(s) && s[*idx] == '(' {
		*idx++ // skip '('
		result = parseExpr(s, idx)
		// skip ')'
		for *idx < len(s) && s[*idx] == ' ' {
			*idx++
		}
		*idx++
	} else {
		// Parse number
		for *idx < len(s) && s[*idx] >= '0' && s[*idx] <= '9' {
			result = result*10 + int(s[*idx]-'0')
			*idx++
		}
	}

	return sign * result
}
