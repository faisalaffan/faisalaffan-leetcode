package main

// LeetCode #2857: Count Pairs of Points With Distance k
// https://leetcode.com/problems/count-pairs-of-points-with-distance-k/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func CountPairsOfPointsWithDistanceK(coordinates [][]int, k int) int {
	cache := make(map[[2]int]int)
	count := 0

	for _, coord := range coordinates {
		x, y := coord[0], coord[1]
		// XOR property: (x1 ^ x2) + (y1 ^ y2) = k
		// For each possible a, b such that a + b = k:
		for a := 0; a <= k; a++ {
			b := k - a
			px := x ^ a
			py := y ^ b
			count += cache[[2]int{px, py}]
		}
		cache[[2]int{x, y}]++
	}

	return count
}

func main() {
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{1, 2}, {4, 2}, {1, 3}, {5, 2}}, 2))
	fmt.Println(CountPairsOfPointsWithDistanceK([][]int{{0, 0}, {1, 1}, {2, 2}}, 2))
}
