package main

// LeetCode #42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/
// Difficulty: Hard

import "fmt"

// trap calculates the total amount of water that can be trapped between bars.
// Uses the two-pointer approach.
//
// Complexity: O(n) time, O(1) space
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}

func main() {
	// Test case from LeetCode
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Test 1:", trap(height)) // 6

	// Additional test cases
	fmt.Println("Test 2: [4,2,0,3,2,5] ->", trap([]int{4, 2, 0, 3, 2, 5})) // 9
	fmt.Println("Test 3: [1,2,3] ->", trap([]int{1, 2, 3}))                // 0
	fmt.Println("Test 4: [] ->", trap([]int{}))                            // 0
}
