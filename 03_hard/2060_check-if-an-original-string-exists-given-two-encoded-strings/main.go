package main

// LeetCode #2060: Check if an Original String Exists Given Two Encoded Strings
// https://leetcode.com/problems/check-if-an-original-string-exists-given-two-encoded-strings/
// Difficulty: Hard
// Approach: DP with memoization on (i, j, diff)

import (
	"fmt"
	"unicode"
)

func possiblyEquals(s1 string, s2 string) bool {
	// Memoization: key = (i, j, diff) where diff = delta length from s1's perspective
	// If diff > 0, s1 has extra length (s2 needs to catch up)
	// If diff < 0, s2 has extra length
	// Offset diff by 2000 to make it non-negative
	memo := make(map[[3]int]bool)
	var dfs func(i, j, diff int) bool
	dfs = func(i, j, diff int) bool {
		if i == len(s1) && j == len(s2) {
			return diff == 0
		}

		key := [3]int{i, j, diff + 2000}
		if val, ok := memo[key]; ok {
			return val
		}

		// Case 1: s1 has a digit
		if i < len(s1) && unicode.IsDigit(rune(s1[i])) {
			end := i
			for end < len(s1) && unicode.IsDigit(rune(s1[end])) {
				end++
			}
			// Generate all possible numbers from digits s1[i:end]
			// Max 3 digits based on problem constraints
			num := 0
			for p := i; p < end; p++ {
				num = num*10 + int(s1[p]-'0')
				if dfs(p+1, j, diff-num) {
					memo[key] = true
					return true
				}
			}
			memo[key] = false
			return false
		}

		// Case 2: s2 has a digit
		if j < len(s2) && unicode.IsDigit(rune(s2[j])) {
			end := j
			for end < len(s2) && unicode.IsDigit(rune(s2[end])) {
				end++
			}
			num := 0
			for p := j; p < end; p++ {
				num = num*10 + int(s2[p]-'0')
				if dfs(i, p+1, diff+num) {
					memo[key] = true
					return true
				}
			}
			memo[key] = false
			return false
		}

		// Case 3: both have letters (or one is empty due to diff)
		if diff > 0 {
			// s1 has extra length, consume from s1
			if i < len(s1) && unicode.IsLetter(rune(s1[i])) {
				if dfs(i+1, j, diff-1) {
					memo[key] = true
					return true
				}
			}
		} else if diff < 0 {
			// s2 has extra length, consume from s2
			if j < len(s2) && unicode.IsLetter(rune(s2[j])) {
				if dfs(i, j+1, diff+1) {
					memo[key] = true
					return true
				}
			}
		} else {
			// diff == 0, both must be letters and equal
			if i < len(s1) && j < len(s2) && unicode.IsLetter(rune(s1[i])) && unicode.IsLetter(rune(s2[j])) && s1[i] == s2[j] {
				if dfs(i+1, j+1, 0) {
					memo[key] = true
					return true
				}
			}
		}

		memo[key] = false
		return false
	}

	return dfs(0, 0, 0)
}

func main() {
	fmt.Println("2060. Check if an Original String Exists Given Two Encoded Strings")

	// Example 1
	s1 := "internationalization"
	s2 := "i18n"
	fmt.Printf("s1=%q s2=%q → %v (expected true)\n", s1, s2, possiblyEquals(s1, s2))

	// Example 2
	s1 = "l123e"
	s2 = "44"
	fmt.Printf("s1=%q s2=%q → %v (expected true)\n", s1, s2, possiblyEquals(s1, s2))

	// Example 3
	s1 = "a5b"
	s2 = "c5b"
	fmt.Printf("s1=%q s2=%q → %v (expected false)\n", s1, s2, possiblyEquals(s1, s2))
}
