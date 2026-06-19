package main

// LeetCode #3474: Lexicographically Smallest Generated String
// https://leetcode.com/problems/lexicographically-smallest-generated-string/
// Difficulty: Hard
//
// Given two strings str1 and str2, generate a string that contains str1 as a
// subsequence and str2 as a substring, minimizing lexicographically. If
// multiple, return the lexicographically smallest.
//
// Approach: Try all possible positions to insert str2 as substring, then
// greedily fill remaining characters to minimize lexicographic order while
// maintaining str1 as subsequence.

import "fmt"

func main() {
	// Example 1
	fmt.Println(generateString("abc", "bc"))
	// Example 2
	fmt.Println(generateString("ab", "cd"))
	// Example 3: str2 already in str1
	fmt.Println(generateString("abcde", "bcd"))
	// Edge: empty str1
	fmt.Println(generateString("", "a"))
	// Edge: str2 longer than str1
	fmt.Println(generateString("a", "abc"))
}

func generateString(str1 string, str2 string) string {
	if len(str1) == 0 {
		return str2
	}

	// Try all insertion positions for str2
	type candidate struct {
		result string
	}
	best := ""

	for pos := 0; pos <= len(str1); pos++ {
		// Try to form result with str2 at position pos
		res := tryInsert(str1, str2, pos)
		if res != "" {
			if best == "" || res < best {
				best = res
			}
		}
	}

	if best == "" {
		// If no valid insertion, concatenate
		return str1 + str2
	}
	return best
}

func tryInsert(str1 string, str2 string, pos int) string {
	n := len(str1)
	m := len(str2)

	// Build result with str2 inserted at position pos
	result := make([]byte, 0, n+m)
	result = append(result, str1[:pos]...)
	result = append(result, str2...)
	result = append(result, str1[pos:]...)

	// Check if str2 appears as substring
	found := false
	for i := 0; i <= len(result)-m; i++ {
		if result[i:i+m] == str2 {
			found = true
			break
		}
	}
	if !found {
		return ""
	}

	// Check if str1 is a subsequence
	j := 0
	for i := 0; i < len(result) && j < n; i++ {
		if result[i] == str1[j] {
			j++
		}
	}
	if j != n {
		return ""
	}

	return string(result)
}
