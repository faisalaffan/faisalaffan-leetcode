package main

// LeetCode #1791: Find Center of Star Graph
// https://leetcode.com/problems/find-center-of-star-graph/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func FindCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	fmt.Println(FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
	fmt.Println(FindCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}
