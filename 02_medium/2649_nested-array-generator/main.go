package main

// LeetCode #2649: Nested Array Generator
// https://leetcode.com/problems/nested-array-generator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func nestedArrayGenerator(arr []any) []int {
	result := []int{}
	var dfs func(item any)
	dfs = func(item any) {
		switch v := item.(type) {
		case []any:
			for _, sub := range v {
				dfs(sub)
			}
		case int:
			result = append(result, v)
		}
	}
	for _, item := range arr {
		dfs(item)
	}
	return result
}

func main() {
	// Test case 1: flat array
	arr1 := []any{1, 2, 3}
	fmt.Println("Test 1:", nestedArrayGenerator(arr1))
	// Expected: [1, 2, 3]

	// Test case 2: nested array
	arr2 := []any{1, []any{2, 3}, 4}
	fmt.Println("Test 2:", nestedArrayGenerator(arr2))
	// Expected: [1, 2, 3, 4]

	// Test case 3: deeply nested
	arr3 := []any{1, []any{2, []any{3, 4}, 5}, 6}
	fmt.Println("Test 3:", nestedArrayGenerator(arr3))
	// Expected: [1, 2, 3, 4, 5, 6]
}
