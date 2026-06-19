package main

// LeetCode #2028: Find Missing Observations
// https://leetcode.com/problems/find-missing-observations/
// Difficulty: Medium
// Time: O(n + m) | Space: O(m)

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	totalSum := mean * (m + n)
	knownSum := 0
	for _, v := range rolls {
		knownSum += v
	}
	missingSum := totalSum - knownSum

	if missingSum < n || missingSum > 6*n {
		return []int{}
	}

	result := make([]int, n)
	for i := range result {
		val := missingSum / (n - i)
		if val > 6 {
			val = 6
		}
		result[i] = val
		missingSum -= val
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", missingRolls([]int{3, 2, 4, 3}, 4, 2))
	// Expected: [6,6]

	// Test case 2
	fmt.Println("Test 2:", missingRolls([]int{1, 5, 6}, 3, 4))
	// Expected: [2,3,2,2]

	// Test case 3
	fmt.Println("Test 3:", missingRolls([]int{1, 2, 3, 4}, 6, 4))
	// Expected: [] (impossible)
}
