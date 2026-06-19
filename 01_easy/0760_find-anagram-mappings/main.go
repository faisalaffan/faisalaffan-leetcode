package main

// LeetCode #760: Find Anagram Mappings
// https://leetcode.com/problems/find-anagram-mappings/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28})) // [1,4,3,2,0]
	fmt.Println(anagramMappings([]int{1, 2}, []int{2, 1}))                              // [1,0]
}

// anagramMappings returns a mapping array P where P[i] is the index of A[i] in B.
// Time: O(n). Space: O(n).
func anagramMappings(nums1 []int, nums2 []int) []int {
	pos := make(map[int]int)
	for i, v := range nums2 {
		pos[v] = i
	}
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = pos[v]
	}
	return result
}
