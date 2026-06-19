package main

// LeetCode #65: Valid Number
// https://leetcode.com/problems/valid-number/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("65. Valid Number")
	testCases := []struct {
		s string
		r bool
	}{
		{"0", true},
		{"0.1", true},
		{"abc", false},
		{"2e10", true},
		{"2e", false},
		{"e3", false},
		{".", false},
		{"1.", true},
		{".1", true},
		{"-.1", true},
		{"+.", false},
		{"3.", true},
		{"3.e10", true},
		{"3.e", false},
		{".e1", false},
		{"+.8", true},
		{"46.e3", true},
		{" 1", true},
		{"1 ", true},
		{"1 1", false},
		{"", false},
	}
	for _, tc := range testCases {
		got := isNumber(tc.s)
		status := "OK"
		if got != tc.r {
			status = "FAIL"
		}
		fmt.Printf("  %q -> %5v (expected %5v) [%s]\n", tc.s, got, tc.r, status)
	}
}

// DFA states:
//
//	0 start (allow whitespace)
//	1 sign before digits
//	2 integer part digits
//	3 decimal point (no digits after yet)
//	4 fractional part digits
//	5 'e' or 'E'
//	6 sign after 'e'
//	7 exponent digits
//	8 trailing whitespace
//	-1 invalid
func isNumber(s string) bool {
	state := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch state {
		case 0: // leading whitespace
			if c == ' ' {
				continue
			} else if c == '+' || c == '-' {
				state = 1
			} else if c >= '0' && c <= '9' {
				state = 2
			} else if c == '.' {
				state = 3
			} else {
				return false
			}
		case 1: // sign before digits
			if c >= '0' && c <= '9' {
				state = 2
			} else if c == '.' {
				state = 3
			} else {
				return false
			}
		case 2: // integer part digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == '.' {
				state = 4
			} else if c == 'e' || c == 'E' {
				state = 5
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 3: // decimal point without integer digits
			if c >= '0' && c <= '9' {
				state = 4
			} else {
				return false
			}
		case 4: // fractional part digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == 'e' || c == 'E' {
				state = 5
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 5: // 'e'/'E'
			if c >= '0' && c <= '9' {
				state = 7
			} else if c == '+' || c == '-' {
				state = 6
			} else {
				return false
			}
		case 6: // sign after 'e'
			if c >= '0' && c <= '9' {
				state = 7
			} else {
				return false
			}
		case 7: // exponent digits
			if c >= '0' && c <= '9' {
				// stay
			} else if c == ' ' {
				state = 8
			} else {
				return false
			}
		case 8: // trailing whitespace
			if c == ' ' {
				// stay
			} else {
				return false
			}
		default:
			return false
		}
	}
	return state == 2 || state == 4 || state == 7 || state == 8
}
