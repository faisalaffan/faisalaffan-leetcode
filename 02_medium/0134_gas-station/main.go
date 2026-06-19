package main

// LeetCode #134: Gas Station
// https://leetcode.com/problems/gas-station/
// Difficulty: Medium

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	start, tank := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	return start
}

func main() {
	// Test case 1
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // 3

	// Test case 2
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})) // -1

	// Test case 3
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) // 4
}

// Time: O(n) | Space: O(1)
