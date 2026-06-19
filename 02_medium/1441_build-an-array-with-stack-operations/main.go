package main

// LeetCode #1441: Build an Array With Stack Operations
// https://leetcode.com/problems/build-an-array-with-stack-operations/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(buildArray([]int{1, 3}, 3)) // ["Push","Push","Pop","Push"]

	// Test case 2
	fmt.Println(buildArray([]int{1, 2, 3}, 3)) // ["Push","Push","Push"]

	// Test case 3
	fmt.Println(buildArray([]int{1, 2}, 4)) // ["Push","Push"]
}

// Time: O(n) where n = max number in target
// Space: O(n) for result
func buildArray(target []int, n int) []string {
	result := make([]string, 0, n*2)
	current := 1

	for _, t := range target {
		for current < t {
			result = append(result, "Push", "Pop")
			current++
		}
		result = append(result, "Push")
		current++
	}

	return result
}
