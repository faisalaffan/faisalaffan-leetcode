package main

// LeetCode #1010: Pairs of Songs With Total Durations Divisible by 60
// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// Difficulty: Medium
//
// Approach: Use modulo counting like Two Sum
// Time: O(n)
// Space: O(60) = O(1)

import "fmt"

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})) // 3
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))           // 3
}

func numPairsDivisibleBy60(time []int) int {
	count := make([]int, 60)
	result := 0

	for _, t := range time {
		mod := t % 60
		need := (60 - mod) % 60
		result += count[need]
		count[mod]++
	}

	return result
}
