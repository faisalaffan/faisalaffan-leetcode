package main

// LeetCode #447: Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	total := 0

	for i := 0; i < len(points); i++ {
		distCount := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			dist := dx*dx + dy*dy
			distCount[dist]++
		}
		for _, count := range distCount {
			total += count * (count - 1)
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numberOfBoomerangs([][]int{{0, 0}}))
	// Expected: 0
}
