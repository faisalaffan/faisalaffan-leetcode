package main

// LeetCode #2225: Find Players With Zero or One Losses
// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	losses := make(map[int]int)
	players := make(map[int]bool)

	for _, m := range matches {
		winner, loser := m[0], m[1]
		players[winner] = true
		players[loser] = true
		losses[loser]++
	}

	winners := []int{}
	oneLoss := []int{}
	for p := range players {
		l := losses[p]
		if l == 0 {
			winners = append(winners, p)
		} else if l == 1 {
			oneLoss = append(oneLoss, p)
		}
	}

	sort.Ints(winners)
	sort.Ints(oneLoss)

	return [][]int{winners, oneLoss}
}

func main() {
	// Test case 1
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	// Expected: [[1,2,10],[4,5,7,8]]

	// Test case 2
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
	// Expected: [[1,2,5,6],[]]
}
