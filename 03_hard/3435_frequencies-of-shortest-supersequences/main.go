package main

// LeetCode #3435: Frequencies of Shortest Supersequences
// https://leetcode.com/problems/frequencies-of-shortest-supersequences/
// Difficulty: Hard
//
// Each word is a pair (a,b). Build a directed graph a→b.
// We want the shortest string that contains every word as subsequence.
// The optimal length = (letters appearing twice) + (letters appearing once).
// Enumerate which letters appear once (mask); the constraint is the resulting
// graph must be acyclic. Use Kahn's algorithm to check.

import "fmt"

func FrequenciesOfShortestSupersequences(words []string) [][]int {
	// Build adjacency from words
	graph := make([][]int, 26)
	indeg := make([]int, 26)
	present := make([]bool, 26)

	for _, w := range words {
		u := int(w[0] - 'a')
		v := int(w[1] - 'a')
		graph[u] = append(graph[u], v)
		indeg[v]++
		present[u] = true
		present[v] = true
	}

	// Collect present letters
	var letters []int
	for i := 0; i < 26; i++ {
		if present[i] {
			letters = append(letters, i)
		}
	}
	L := len(letters)

	// Map letter to index in letters slice
	pos := make([]int, 26)
	for i, c := range letters {
		pos[c] = i
	}

	var result [][]int
	shortestLen := int(1e9)

	for mask := 0; mask < (1 << L); mask++ {
		// mask bit = 1 → letter appears once (frequency 1)
		// mask bit = 0 → letter appears twice (frequency 2)

		tmpIndeg := make([]int, 26)
		copy(tmpIndeg, indeg)

		// For letters that appear once, self-loops (u→u) must be removed because
		// if a letter appears only once, it cannot serve as both source and target
		// of the same pair.
		valid := true
		for _, u := range letters {
			if mask>>pos[u]&1 == 1 {
				// Letter appears once: remove all outgoing edges to itself
				for j := 0; j < len(graph[u]); j++ {
					v := graph[u][j]
					if v == u {
						tmpIndeg[v]--
					}
				}
			}
		}

		// Kahn's topological sort
		q := make([]int, 0, L)
		visited := 0
		for _, u := range letters {
			if tmpIndeg[u] == 0 {
				q = append(q, u)
				visited++
			}
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				// Skip self-loop if u appears once
				if v == u && mask>>pos[u]&1 == 1 {
					continue
				}
				tmpIndeg[v]--
				if tmpIndeg[v] == 0 {
					q = append(q, v)
					visited++
				}
			}
		}

		if visited != L {
			// Cycle detected — this mask is invalid
			continue
		}

		// Build frequency array
		twos := 0
		for _, u := range letters {
			if mask>>pos[u]&1 == 0 {
				twos++
			}
		}
		// Total length = 2*twos + 1*(L-twos) = L + twos
		totalLen := L + twos
		if totalLen < shortestLen {
			shortestLen = totalLen
			result = nil
		}
		if totalLen == shortestLen {
			freq := make([]int, 26)
			for _, u := range letters {
				if mask>>pos[u]&1 == 1 {
					freq[u] = 1
				} else {
					freq[u] = 2
				}
			}
			result = append(result, freq)
		}
	}
	return result
}

func main() {
	// Test: words=["ab","bc","ac"] => expected [2,2,2] (all three letters appear twice)
	res := FrequenciesOfShortestSupersequences([]string{"ab", "bc", "ac"})
	fmt.Printf("words=[ab,bc,ac] -> %v\n", res)

	// Test: words=["ab","bc"] => all letters appear once or twice, check
	res2 := FrequenciesOfShortestSupersequences([]string{"ab", "bc"})
	fmt.Printf("words=[ab,bc] -> %v\n", res2)

	// Test: words=["aa","ab","ba"]
	res3 := FrequenciesOfShortestSupersequences([]string{"aa", "ab", "ba"})
	fmt.Printf("words=[aa,ab,ba] -> %v\n", res3)
}
