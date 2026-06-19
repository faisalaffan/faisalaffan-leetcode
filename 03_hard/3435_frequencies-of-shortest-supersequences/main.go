package main

// LeetCode #3435: Frequencies of Shortest Supersequences
// https://leetcode.com/problems/frequencies-of-shortest-supersequences/
// Difficulty: Hard
//
// Each letter appears 1 or 2 times. Enumerate masks (appears-once letters).
// Kahn's algorithm to check acyclicity.

import "fmt"

func main() {
	fmt.Println(FrequenciesOfShortestSupersequences([]string{"ab", "bc"}))
}

func FrequenciesOfShortestSupersequences(words []string) [][]int {
	graph := make([][]int, 26)
	indeg := make([]int, 26)
	has := make([]bool, 26)

	for _, w := range words {
		u := int(w[0] - 'a')
		v := int(w[1] - 'a')
		graph[u] = append(graph[u], v)
		indeg[v]++
		has[u] = true
		has[v] = true
	}

	// Count letters that appear (have any edges)
	var letters []int
	for i := 0; i < 26; i++ {
		if has[i] {
			letters = append(letters, i)
		}
	}
	L := len(letters)

	var result [][]int

	// Try each mask: bit = 1 means this letter appears only once (cannot form duplicate pair)
	for mask := 0; mask < (1 << L); mask++ {
		// Build in-degree for this mask
		tmpIndeg := make([]int, 26)
		copy(tmpIndeg, indeg)

		// For letters that appear once, they can't be destination of edges from same letter
		// Actually: If a letter appears only once, it can't be both source and target of
		// same letter pair. Mask bit = 1 means this letter appears once (1 occurrence).
		for i, u := range letters {
			if mask>>i&1 == 1 {
				// This letter appears once. Remove all outgoing edges from u to u (self-loops)
				// and adjust indegrees.
				// A letter appearing once means it cannot be in both positions of a pair.
				// So we remove edges where it would form a 2-occurrence pattern.
				for j := 0; j < len(graph[u]); j++ {
					v := graph[u][j]
					if v == u {
						tmpIndeg[v]--
					}
				}
			}
		}

		// Kahn's algorithm to check acyclicity
		q := make([]int, 0, 26)
		cnt := 0
		for _, u := range letters {
			if tmpIndeg[u] == 0 {
				q = append(q, u)
				cnt++
			}
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if v == u && mask>>(indexOf(letters, u))&1 == 1 {
					continue
				}
				tmpIndeg[v]--
				if tmpIndeg[v] == 0 {
					q = append(q, v)
					cnt++
				}
			}
		}

		if cnt == L {
			// Acyclic
			freq := make([]int, 26)
			for i, u := range letters {
				if mask>>i&1 == 1 {
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

func indexOf(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
