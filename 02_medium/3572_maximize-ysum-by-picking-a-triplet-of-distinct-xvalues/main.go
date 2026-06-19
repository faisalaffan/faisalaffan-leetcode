package main

// LeetCode #3572: Maximize Y-Sum by Picking a Triplet of Distinct X-Values
// https://leetcode.com/problems/maximize-ysum-by-picking-a-triplet-of-distinct-xvalues/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	points := [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 5}}
	fmt.Println("Test 1:", MaximizeYsumByPickingATripletOfDistinctXvalues(points))
	// Test case 2
	points2 := [][]int{{1, 5}, {2, 5}, {3, 5}}
	fmt.Println("Test 2:", MaximizeYsumByPickingATripletOfDistinctXvalues(points2))
	// Test case 3
	points3 := [][]int{{1, 10}, {2, 1}, {3, 1}}
	fmt.Println("Test 3:", MaximizeYsumByPickingATripletOfDistinctXvalues(points3))
}

func MaximizeYsumByPickingATripletOfDistinctXvalues(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	// Sort by y descending
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] > points[j][1]
	})

	used := make(map[int]bool)
	sum := 0
	count := 0
	for _, p := range points {
		if !used[p[0]] {
			used[p[0]] = true
			sum += p[1]
			count++
			if count == 3 {
				return sum
			}
		}
	}
	return sum
}
