package main

// LeetCode #2214: Minimum Health to Beat Game
// https://leetcode.com/problems/minimum-health-to-beat-game/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func minimumHealth(damage []int, armor int) int64 {
	var total int64 = 0
	maxDmg := 0
	for _, d := range damage {
		total += int64(d)
		if d > maxDmg {
			maxDmg = d
		}
	}
	// We can use armor to reduce the largest damage
	saved := armor
	if maxDmg < saved {
		saved = maxDmg
	}
	return total - int64(saved) + 1
}

func main() {
	// Test case 1
	fmt.Println(minimumHealth([]int{2, 7, 4, 3}, 4))
	// Expected: 13

	// Test case 2
	fmt.Println(minimumHealth([]int{3, 3, 3}, 0))
	// Expected: 10

	// Test case 3
	fmt.Println(minimumHealth([]int{1, 2, 3, 4}, 5))
	// Expected: 7
}
