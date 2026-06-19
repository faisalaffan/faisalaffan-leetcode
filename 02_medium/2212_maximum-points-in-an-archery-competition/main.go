package main

// LeetCode #2212: Maximum Points in an Archery Competition
// https://leetcode.com/problems/maximum-points-in-an-archery-competition/
// Difficulty: Medium
// Time: O(n * 2^n) | Space: O(n)

import "fmt"

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	bestScore := 0
	var bestMask int

	for mask := 1; mask < (1 << 12); mask++ {
		arrows := 0
		score := 0
		for i := 0; i < 12; i++ {
			if mask&(1<<i) != 0 {
				arrows += aliceArrows[i] + 1
				score += i
			}
		}
		if arrows <= numArrows && score > bestScore {
			bestScore = score
			bestMask = mask
		}
	}

	result := make([]int, 12)
	used := 0
	for i := 0; i < 12; i++ {
		if bestMask&(1<<i) != 0 {
			result[i] = aliceArrows[i] + 1
			used += result[i]
		}
	}
	// Put remaining arrows in first section (index 0)
	if used < numArrows {
		result[0] += numArrows - used
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))
	// Expected: [0,0,0,0,0,0,0,0,1,1,1,0] or similar valid

	// Test case 2
	fmt.Println(maximumBobPoints(3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}))
	// Expected: [0,1,1,0,0,0,0,0,0,0,0,0] or similar valid
}
