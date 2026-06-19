package main

// LeetCode #2820: Election Results
// https://leetcode.com/problems/election-results/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Vote struct {
	Voter   string
	Candidate string
}

func ElectionResults(votes []Vote) string {
	counts := make(map[string]int)
	for _, v := range votes {
		counts[v.Candidate]++
	}

	type kv struct {
		Candidate string
		Count     int
	}
	sorted := make([]kv, 0, len(counts))
	for k, v := range counts {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Count != sorted[j].Count {
			return sorted[i].Count > sorted[j].Count
		}
		return sorted[i].Candidate < sorted[j].Candidate
	})

	if len(sorted) > 0 {
		return sorted[0].Candidate
	}
	return ""
}

func main() {
	votes := []Vote{
		{"A", "Alice"}, {"B", "Bob"}, {"C", "Alice"},
	}
	fmt.Println(ElectionResults(votes))

	votes2 := []Vote{
		{"X", "Cand1"}, {"Y", "Cand1"}, {"Z", "Cand2"},
	}
	fmt.Println(ElectionResults(votes2))
}
