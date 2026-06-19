package main

// LeetCode #2374: Node With Highest Edge Score
// https://leetcode.com/problems/node-with-highest-edge-score/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Each node i points to edges[i]. Score of j = sum of i where edges[i] == j.

import "fmt"

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5})) // 7
	fmt.Println(edgeScore([]int{2, 0, 0, 2}))               // 0
}

func edgeScore(edges []int) int {
	n := len(edges)
	score := make([]int, n)
	for i, to := range edges {
		score[to] += i
	}

	maxScore := -1
	ans := -1
	for i, s := range score {
		if s > maxScore {
			maxScore = s
			ans = i
		}
	}
	return ans
}
