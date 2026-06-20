package main

// LeetCode #3615: Longest Palindromic Path in Graph
// https://leetcode.com/problems/longest-palindromic-path-in-graph/
// Difficulty: Hard
//
// Given a graph with n nodes (n <= 14) and labels, find the longest palindrome
// that can be formed by visiting unique nodes along a valid path.
//
// Approach: Bitmask DP. For each node, try expanding paths palindromically,
// using DP[mask][last] to track reachable states.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxLen(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, "abba"))
	// Example 2
	fmt.Println(maxLen(3, [][]int{{0, 1}, {1, 2}}, "aba"))
	// Edge: single node
	fmt.Println(maxLen(1, [][]int{}, "a"))
	// Edge: no palindrome possible
	fmt.Println(maxLen(2, [][]int{{0, 1}}, "ab"))
}

func maxLen(n int, edges [][]int, label string) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], v)
	}

	result := 1

	// dp[mask][last][first] Not feasible for n=14.
	// Instead: start from each pair as center, expand outward.

	// Odd length: single center node
	for i := 0; i < n; i++ {
		result = max(result, 1)
	}

	// Even length: pair center (edge between i and j with same label)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if label[i] == label[j] && hasEdge(adj, i, j) {
				// Start expanding from center pair (i, j)
				result = max(result, expand(adj, label, i, j, 1<<i|1<<j, 2))
			}
		}
	}

	// Also try odd length centers and expand
	for i := 0; i < n; i++ {
		result = max(result, expandOdd(adj, label, i, 1<<i, 1))
	}

	return result
}

func expand(adj [][]int, label string, left, right int, mask int, length int) int {
	n := len(label)
	best := length
	// Try to add matching endpoints
	for _, u := range adj[left] {
		if mask&(1<<u) != 0 {
			continue
		}
		for _, v := range adj[right] {
			if mask&(1<<v) != 0 || u == v {
				continue
			}
			if label[u] == label[v] {
				newMask := mask | (1 << u) | (1 << v)
				best = max(best, expand(adj, label, u, v, newMask, length+2))
			}
		}
		// Single expansion from one side (odd length possibility)
	}
	return best
}

func expandOdd(adj [][]int, label string, center int, mask int, length int) int {
	n := len(label)
	best := length
	for _, u := range adj[center] {
		if mask&(1<<u) != 0 {
			continue
		}
		// Try to expand to u, then find pair for u
		for _, v := range adj[center] {
			if u == v || mask&(1<<v) != 0 {
				continue
			}
			if label[u] == label[v] {
				newMask := mask | (1 << u) | (1 << v)
				best = max(best, expand(adj, label, u, v, newMask, length+2))
			}
		}
		// Or just extend odd-length path
		best = max(best, 1+expandSingle(adj, label, u, mask|(1<<u), 1))
	}
	return best
}

func expandSingle(adj [][]int, label string, node int, mask int, length int) int {
	best := length
	for _, v := range adj[node] {
		if mask&(1<<v) != 0 {
			continue
		}
		best = max(best, expandSingle(adj, label, v, mask|(1<<v), length+1))
	}
	return best
}

func hasEdge(adj [][]int, u, v int) bool {
	for _, x := range adj[u] {
		if x == v {
			return true
		}
	}
	return false
}
