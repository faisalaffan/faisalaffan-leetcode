package main

// LeetCode #2350: Shortest Impossible Sequence of Rolls
// https://leetcode.com/problems/shortest-impossible-sequence-of-rolls/
// Difficulty: Hard
//
// Approach: Greedy. We want the shortest sequence of rolls (each 1..k) that
// is NOT a subsequence of the given rolls array. The key insight: in each
// "round", collect all k distinct values 1..k. Count how many complete rounds
// we can make. Answer = rounds + 1.
//
// Example: rolls=[4,2,1,2,3,3,2,4,1], k=4
// Round 1: collect {4,2,1,3} (positions 0,1,2,4)
// Round 2: collect {3,2,4,1} (positions 5,6,7,8)
// rounds=2, answer=3

import "fmt"

func main() {
	// Example 1: [4,2,1,2,3,3,2,4,1], 4 => 3
	fmt.Println(shortestSequence([]int{4, 2, 1, 2, 3, 3, 2, 4, 1}, 4))
	// Example 2: [1,1,2,3], 4 => 2
	fmt.Println(shortestSequence([]int{1, 1, 2, 3}, 4))
	// Example 3: [1,2,3,4], 5 => 2
	fmt.Println(shortestSequence([]int{1, 2, 3, 4}, 5))
	// Edge: first roll already incomplete
	fmt.Println(shortestSequence([]int{1, 1, 1, 1}, 3))
	// Edge: single round
	fmt.Println(shortestSequence([]int{1, 2, 3}, 3))
}

func shortestSequence(rolls []int, k int) int {
	seen := make([]bool, k+1)
	rounds := 0
	count := 0

	for _, r := range rolls {
		if !seen[r] {
			seen[r] = true
			count++
			if count == k {
				rounds++
				count = 0
				// Reset seen
				seen = make([]bool, k+1)
			}
		}
	}
	return rounds + 1
}
