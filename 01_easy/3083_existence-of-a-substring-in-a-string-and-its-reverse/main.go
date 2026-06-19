package main

// LeetCode #3083: Existence of a Substring in a String and Its Reverse
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isSubstringPresent
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("leetcode")) // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcba"))   // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcd"))    // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isSubstringPresent
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool {
	// Build set of all substrings of length 2
	substrings := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		substrings[s[i:i+2]] = true
	}

	// Check reverse for any of those substrings
	for i := len(s) - 1; i > 0; i-- {
		if substrings[string(s[i])+string(s[i-1])] {
			return true
		}
	}
	return false
}
