package main

// LeetCode #3218: Minimum Cost for Cutting Cake I
// https://leetcode.com/problems/minimum-cost-for-cutting-cake-i/
// Difficulty: Medium
// Time: O(h log h + v log v) | Space: O(log h + log v)

import (
	"fmt"
	"sort"
)

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	sort.Slice(horizontalCut, func(i, j int) bool { return horizontalCut[i] > horizontalCut[j] })
	sort.Slice(verticalCut, func(i, j int) bool { return verticalCut[i] > verticalCut[j] })

	hPieces, vPieces := 1, 1
	i, j := 0, 0
	cost := 0

	for i < len(horizontalCut) && j < len(verticalCut) {
		if horizontalCut[i] >= verticalCut[j] {
			cost += horizontalCut[i] * vPieces
			hPieces++
			i++
		} else {
			cost += verticalCut[j] * hPieces
			vPieces++
			j++
		}
	}

	for i < len(horizontalCut) {
		cost += horizontalCut[i] * vPieces
		i++
	}
	for j < len(verticalCut) {
		cost += verticalCut[j] * hPieces
		j++
	}
	return cost
}

func main() {
	fmt.Println(minimumCost(3, 2, []int{1, 3}, []int{5})) // Expected: 13
	fmt.Println(minimumCost(2, 2, []int{7}, []int{4}))     // Expected: 15
}
