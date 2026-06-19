package main

// LeetCode #2410: Maximum Matching of Players With Trainers
// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/
// Difficulty: Medium
// Time: O(n log n + m log m) | Space: O(1)
// Sort both, greedy match: assign smallest sufficient trainer to each player.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8})) // 2
	fmt.Println(matchPlayersAndTrainers([]int{1, 1, 1}, []int{10}))         // 1
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			i++
		}
		j++
	}
	return i
}
