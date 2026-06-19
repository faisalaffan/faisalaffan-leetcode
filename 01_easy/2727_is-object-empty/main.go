package main

// LeetCode #2727: Is Object Empty
// https://leetcode.com/problems/is-object-empty/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Checks if array/slice is empty.

import "fmt"

func main() {
	fmt.Println(IsObjectEmpty([]int{}))
	fmt.Println(IsObjectEmpty([]int{1, 2}))
}

func IsObjectEmpty(obj interface{}) bool {
	switch v := obj.(type) {
	case []int:
		return len(v) == 0
	case []string:
		return len(v) == 0
	default:
		return false
	}
}
