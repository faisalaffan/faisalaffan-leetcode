package main

// LeetCode #2924: Find Champion II
// https://leetcode.com/problems/find-champion-ii/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(findChampionII(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampionII(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampionII(2, [][]int{{0, 1}}))
}

func findChampionII(n int, edges [][]int) int {
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i, x := range indeg {
		if x == 0 {
			cnt++
			ans = i
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
