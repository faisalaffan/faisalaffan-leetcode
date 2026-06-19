package main

// LeetCode #2041: Accepted Candidates From the Interviews
// https://leetcode.com/problems/accepted-candidates-from-the-interviews/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Candidate struct {
	ID       int
	Scores   []int
	Accepted bool
}

func acceptedCandidates(candidates []Candidate, minScore int) []int {
	result := []int{}
	for _, c := range candidates {
		count := 0
		for _, s := range c.Scores {
			if s >= minScore {
				count++
			}
		}
		if count >= 2 {
			// Calculate total
			total := 0
			sorted := make([]int, len(c.Scores))
			copy(sorted, c.Scores)
			sort.Ints(sorted)
			for _, s := range sorted[1:] {
				total += s
			}
			result = append(result, c.ID)
			_ = total
		}
	}
	return result
}

func main() {
	// Test case 1
	candidates1 := []Candidate{
		{1, []int{5, 6, 7}, false},
		{2, []int{8, 9, 10}, false},
	}
	fmt.Println("Test 1:", acceptedCandidates(candidates1, 5))
	// Expected: [1, 2]

	// Test case 2
	candidates2 := []Candidate{
		{1, []int{4, 5, 6}, false},
		{2, []int{7, 3, 8}, false},
		{3, []int{2, 3, 4}, false},
	}
	fmt.Println("Test 2:", acceptedCandidates(candidates2, 5))
	// Expected: [1, 2]

	// Test case 3
	fmt.Println("Test 3:", acceptedCandidates(nil, 5))
	// Expected: []
}
