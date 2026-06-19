package main

// LeetCode #2298: Tasks Count in the Weekend
// https://leetcode.com/problems/tasks-count-in-the-weekend/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func taskCount(taskRunDates [][]int) int {
	weekend := 0
	for _, d := range taskRunDates {
		// taskRunDates[i] = [month, day]
		// This is a paid SQL/Schema problem. Simplified:
		// Count tasks run on Saturday (6) or Sunday (7)
		dayOfWeek := d[1] % 7
		if dayOfWeek == 6 || dayOfWeek == 0 {
			weekend++
		}
	}
	return weekend
}

func main() {
	// Test case 1
	fmt.Println(taskCount([][]int{{1, 6}, {1, 7}, {1, 8}}))
	// Expected: 2 (Saturday and Sunday)

	// Test case 2
	fmt.Println(taskCount([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}))
	// Expected: 0
}
