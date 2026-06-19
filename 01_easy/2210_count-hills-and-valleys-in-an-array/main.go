package main

// LeetCode #2210: Count Hills and Valleys in an Array
// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountHillsAndValleysInAnArray([]int{2, 4, 1, 1, 6, 5})) // 3
	fmt.Println(CountHillsAndValleysInAnArray([]int{6, 6, 5, 5, 4, 1})) // 0
}

// Time: O(n), Space: O(1)
func CountHillsAndValleysInAnArray(nums []int) int {
	// Build flattened array (remove consecutive duplicates)
	flat := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			flat = append(flat, nums[i])
		}
	}

	count := 0
	for i := 1; i < len(flat)-1; i++ {
		if (flat[i] > flat[i-1] && flat[i] > flat[i+1]) || // hill
			(flat[i] < flat[i-1] && flat[i] < flat[i+1]) { // valley
			count++
		}
	}
	return count
}
