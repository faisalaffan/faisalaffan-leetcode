package main

// LeetCode #2078: Two Furthest Houses With Different Colors
// https://leetcode.com/problems/two-furthest-houses-with-different-colors/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 1, 1, 6, 1, 1, 1})) // 3
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{1, 2, 3, 4, 5}))       // 4
	fmt.Println(TwoFurthestHousesWithDifferentColors([]int{0, 1}))                // 1
}

// Time: O(n), Space: O(1)
func TwoFurthestHousesWithDifferentColors(colors []int) int {
	n := len(colors)
	maxDist := 0

	// Check from leftmost with rightmost
	if colors[0] != colors[n-1] {
		return n - 1
	}

	// If ends are same, find furthest different color from either end
	for i := 1; i < n-1; i++ {
		if colors[i] != colors[0] {
			dist := n - 1 - i
			if dist > maxDist {
				maxDist = dist
			}
			dist = i
			if dist > maxDist {
				maxDist = dist
			}
			break
		}
	}
	return maxDist
}
