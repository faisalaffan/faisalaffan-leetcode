package main

// LeetCode #1807: Evaluate the Bracket Pairs of a String
// https://leetcode.com/problems/evaluate-the-bracket-pairs-of-a-string/
// Difficulty: Medium
// Time: O(n), Space: O(k) where k = number of knowledge pairs

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string]string)
	for _, kv := range knowledge {
		dict[kv[0]] = kv[1]
	}

	var result strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			j := i + 1
			for s[j] != ')' {
				j++
			}
			key := s[i+1 : j]
			if val, ok := dict[key]; ok {
				result.WriteString(val)
			} else {
				result.WriteByte('?')
			}
			i = j + 1
		} else {
			result.WriteByte(s[i])
			i++
		}
	}
	return result.String()
}

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}})) // Expected: "bobistwoyearsold"
	fmt.Println(evaluate("hi(name)", [][]string{{"a", "b"}})) // Expected: "hi?"
	fmt.Println(evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}})) // Expected: "yesyesyesaaa"
}
