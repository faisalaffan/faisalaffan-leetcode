package main

// LeetCode #3694: Distinct Points Reachable After Substring Removal
// https://leetcode.com/problems/distinct-points-reachable-after-substring-removal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func distinctPointsReachableAfterSubstringRemoval(s string, k int) int {
	n := len(s)
	// prefix sums for x and y
	fx := make([]int, n+1)
	fy := make([]int, n+1)
	for i, ch := range s {
		fx[i+1] = fx[i]
		fy[i+1] = fy[i]
		switch ch {
		case 'U':
			fy[i+1]++
		case 'D':
			fy[i+1]--
		case 'L':
			fx[i+1]--
		case 'R':
			fx[i+1]++
		}
	}

	seen := make(map[[2]int]bool)
	for i := k; i <= n; i++ {
		x := fx[n] - (fx[i] - fx[i-k])
		y := fy[n] - (fy[i] - fy[i-k])
		seen[[2]int{x, y}] = true
	}

	return len(seen)
}

func main() {
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("LUL", 1))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UDLR", 4))
	fmt.Println(distinctPointsReachableAfterSubstringRemoval("UU", 1))
}
