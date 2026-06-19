package main

// LeetCode #2625: Flatten Deeply Nested Array
// https://leetcode.com/problems/flatten-deeply-nested-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func flatten(arr []any, depth int) []any {
	result := []any{}

	var dfs func(item any, currDepth int)
	dfs = func(item any, currDepth int) {
		switch v := item.(type) {
		case []any:
			if currDepth < depth {
				for _, sub := range v {
					dfs(sub, currDepth+1)
				}
			} else {
				result = append(result, v)
			}
		default:
			result = append(result, v)
		}
	}

	for _, item := range arr {
		dfs(item, 0)
	}
	return result
}

func main() {
	// Test case 1: flatten to depth 1
	arr1 := []any{1, []any{2, []any{3, 4}}, 5}
	fmt.Println("Test 1:", flatten(arr1, 1))
	// Expected: [1, 2, [3, 4], 5]

	// Test case 2: flatten to depth 2
	fmt.Println("Test 2:", flatten(arr1, 2))
	// Expected: [1, 2, 3, 4, 5]

	// Test case 3: depth 0 - no flattening
	fmt.Println("Test 3:", flatten(arr1, 0))
	// Expected: [1, [2, [3, 4]], 5]
}
