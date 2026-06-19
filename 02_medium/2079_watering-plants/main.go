package main

// LeetCode #2079: Watering Plants
// https://leetcode.com/problems/watering-plants/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func wateringPlants(plants []int, capacity int) int {
	steps := 0
	curWater := capacity

	for i := 0; i < len(plants); i++ {
		if curWater < plants[i] {
			// Go back to river (i steps) and come back (i steps)
			steps += i*2 + 1
			curWater = capacity - plants[i]
		} else {
			curWater -= plants[i]
			steps++
		}
	}
	return steps
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", wateringPlants([]int{2, 2, 3, 3}, 5))
	// Expected: 14

	// Test case 2
	fmt.Println("Test 2:", wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4))
	// Expected: 30

	// Test case 3
	fmt.Println("Test 3:", wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8))
	// Expected: 49
}
