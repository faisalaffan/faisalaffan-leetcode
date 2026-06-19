package main

// LeetCode #2053: Kth Distinct String in an Array
// https://leetcode.com/problems/kth-distinct-string-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KthDistinctStringInAnArray([]string{"d", "b", "c", "b", "c", "a"}, 2)) // "a"
	fmt.Println(KthDistinctStringInAnArray([]string{"aaa", "aa", "a"}, 1))              // "aaa"
	fmt.Println(KthDistinctStringInAnArray([]string{"a", "b", "a"}, 3))                 // ""
}

// Time: O(n), Space: O(n)
func KthDistinctStringInAnArray(arr []string, k int) string {
	freq := make(map[string]int)
	for _, s := range arr {
		freq[s]++
	}

	idx := 1
	for _, s := range arr {
		if freq[s] == 1 {
			if idx == k {
				return s
			}
			idx++
		}
	}
	return ""
}
