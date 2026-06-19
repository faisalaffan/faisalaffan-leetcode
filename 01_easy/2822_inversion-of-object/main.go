package main

// LeetCode #2822: Inversion of Object
// https://leetcode.com/problems/inversion-of-object/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Swaps keys and values of a map.

import "fmt"

func main() {
	input := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(InversionOfObject(input))
}

func InversionOfObject(obj map[string]int) map[int]string {
	result := make(map[int]string, len(obj))
	for k, v := range obj {
		result[v] = k
	}
	return result
}
