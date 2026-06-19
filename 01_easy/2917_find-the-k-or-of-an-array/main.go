package main

// LeetCode #2917: Find the K-or of an Array
// https://leetcode.com/problems/find-the-k-or-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findKOr
	fmt.Println(FindTheKOrOfAnArray([]int{7, 12, 9, 8, 9, 15}, 4)) // 9
	fmt.Println(FindTheKOrOfAnArray([]int{2, 12, 1, 11, 4, 5}, 6)) // 0
	fmt.Println(FindTheKOrOfAnArray([]int{10, 8, 5, 9, 11, 6, 8}, 1)) // 15
}

// Time: O(n * 32) | Space: O(1)
// LeetCode submission name: findKOr
func FindTheKOrOfAnArray(nums []int, k int) int {
	result := 0
	for bit := 0; bit < 32; bit++ {
		count := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				count++
			}
		}
		if count >= k {
			result |= (1 << bit)
		}
	}
	return result
}
