package main

// LeetCode #11: Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
// Difficulty: Medium

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxWater {
			maxWater = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	// Test case 1
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49

	// Test case 2
	fmt.Println(maxArea([]int{1, 1})) // 1

	// Test case 3
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4})) // 16
}

// Time: O(n) | Space: O(1)
