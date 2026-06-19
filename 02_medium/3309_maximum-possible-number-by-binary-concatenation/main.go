package main

// LeetCode #3309: Maximum Possible Number by Binary Concatenation
// https://leetcode.com/problems/maximum-possible-number-by-binary-concatenation/
// Difficulty: Medium
// Time: O(1) Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxGoodNumber([]int{1, 2, 3}))    // 30
	fmt.Println(maxGoodNumber([]int{2, 8, 16}))   // 1296
	fmt.Println(maxGoodNumber([]int{1, 1, 1}))    // 7
}

func maxGoodNumber(nums []int) int {
	// Try all 6 permutations
	perms := [][]int{
		{nums[0], nums[1], nums[2]},
		{nums[0], nums[2], nums[1]},
		{nums[1], nums[0], nums[2]},
		{nums[1], nums[2], nums[0]},
		{nums[2], nums[0], nums[1]},
		{nums[2], nums[1], nums[0]},
	}

	maxVal := 0
	for _, p := range perms {
		val := 0
		for _, x := range p {
			bits := 0
			temp := x
			for temp > 0 {
				bits++
				temp >>= 1
			}
			if x == 0 {
				bits = 1
			}
			val = (val << bits) | x
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
