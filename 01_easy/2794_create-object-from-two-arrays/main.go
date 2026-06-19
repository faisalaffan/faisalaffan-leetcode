package main

// LeetCode #2794: Create Object from Two Arrays
// https://leetcode.com/problems/create-object-from-two-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Creates a map from keys and values arrays.

import "fmt"

func main() {
	fmt.Println(CreateObjectFromTwoArrays([]string{"a", "b", "c"}, []int{1, 2, 3}))
	fmt.Println(CreateObjectFromTwoArrays([]string{"x"}, []int{10}))
}

func CreateObjectFromTwoArrays(keys []string, values []int) map[string]int {
	result := make(map[string]int, len(keys))
	for i, k := range keys {
		if i < len(values) {
			result[k] = values[i]
		}
	}
	return result
}
