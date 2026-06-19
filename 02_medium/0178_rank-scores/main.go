package main

// LeetCode #178: Rank Scores
// https://leetcode.com/problems/rank-scores/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func rankScores(scores []int) []int {
	if len(scores) == 0 {
		return nil
	}

	type pair struct {
		score int
		idx   int
	}

	pairs := make([]pair, len(scores))
	for i, s := range scores {
		pairs[i] = pair{s, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	ranks := make([]int, len(scores))
	rank := 1
	for i, p := range pairs {
		if i > 0 && p.score < pairs[i-1].score {
			rank = i + 1
		}
		ranks[p.idx] = rank
	}

	return ranks
}

func main() {
	fmt.Println(rankScores([]int{100, 90, 90, 80}))
	fmt.Println(rankScores([]int{50, 60, 70}))
	fmt.Println(rankScores([]int{90, 80, 80, 70, 60, 60}))
}
