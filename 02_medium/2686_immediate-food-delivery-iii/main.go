package main

// LeetCode #2686: Immediate Food Delivery III
// https://leetcode.com/problems/immediate-food-delivery-iii/
// Difficulty: Medium [Paid] (SQL problem)
// Time: O(n) | Space: O(1)

import "fmt"

func immediateFoodDelivery(delivery []int, schedule []int) int {
	// Count how many scheduled deliveries can be made immediately
	// (Simplified Go equivalent of the SQL problem)
	available := make(map[int]bool)
	for _, d := range delivery {
		available[d] = true
	}
	count := 0
	for _, s := range schedule {
		if available[s] {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", immediateFoodDelivery([]int{1, 2, 3}, []int{1, 3, 5}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", immediateFoodDelivery([]int{1, 2}, []int{3, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", immediateFoodDelivery([]int{}, []int{1}))
	// Expected: 0
}
