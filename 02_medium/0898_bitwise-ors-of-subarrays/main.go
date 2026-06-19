package main

// LeetCode #898: Bitwise ORs of Subarrays
// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BitwiseOrsOfSubarrays([]int{0}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 1, 2}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 2, 4}))
}

// Time: O(n * log(max)) | Space: O(n)
func BitwiseOrsOfSubarrays(arr []int) int {
	set := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[i]|arr[j] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			set[arr[j]] = true
		}
	}

	return len(set)
}
