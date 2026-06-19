package main

// LeetCode #3019: Number of Changing Keys
// https://leetcode.com/problems/number-of-changing-keys/
// Difficulty: Easy

import "fmt"
import "unicode"

func main() {
	// LeetCode name: countKeyChanges
	fmt.Println(NumberOfChangingKeys("aAbBcC")) // 2
	fmt.Println(NumberOfChangingKeys("AaAaAaaA")) // 0
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: countKeyChanges
func NumberOfChangingKeys(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i-1])) {
			count++
		}
	}
	return count
}
