package main

// LeetCode #3014: Minimum Number of Pushes to Type Word I
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumPushes
	fmt.Println(MinimumNumberOfPushesToTypeWordI("abcde")) // 5
	fmt.Println(MinimumNumberOfPushesToTypeWordI("xycdefghij")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumPushes
// Each key can hold up to 4 distinct letters. First 8 letters cost 1 push each,
// next 8 cost 2 pushes, etc.
func MinimumNumberOfPushesToTypeWordI(word string) int {
	n := len(word)
	pushes := 0
	// First 8 distinct chars: 1 push each
	// Next 8: 2 pushes each, etc.
	for i := 0; i < n; i++ {
		pushes += (i / 8) + 1
	}
	return pushes
}
