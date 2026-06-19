package main

// LeetCode #574: Winning Candidate
// https://leetcode.com/problems/winning-candidate/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// votes: {candidate_id}
	votes := []int{1, 2, 2, 3, 3, 3}
	// candidate names: map candidate_id -> name
	candidateNames := map[int]int{1: 1001, 2: 1002, 3: 1003}
	fmt.Println(FindWinningCandidate(votes, candidateNames))
}

func FindWinningCandidate(votes []int, candidateNames map[int]int) int {
	counts := make(map[int]int)
	for _, v := range votes {
		counts[v]++
	}

	maxVotes := 0
	winner := -1
	for id, count := range counts {
		if count > maxVotes {
			maxVotes = count
			winner = candidateNames[id]
		}
	}

	return winner
}
