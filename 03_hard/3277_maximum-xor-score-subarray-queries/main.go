package main

// LeetCode #3277: Maximum XOR Score Subarray Queries
// https://leetcode.com/problems/maximum-xor-score-subarray-queries/
// Difficulty: Hard
//
// Given an array nums and queries [l, r], for each query find the maximum
// XOR of any subarray within nums[l..r] (inclusive).
//
// Approach:
//  1. Compute prefix XOR array pref where pref[i] = XOR of nums[0..i-1].
//  2. The XOR of subarray (i, j) = pref[i] ^ pref[j+1].
//  3. Precompute dp[l][r] = max XOR subarray in [l, r] using a trie-based DP.
//  4. Answer queries in O(1).

import (
	"fmt"
)

func main() {
	// Example 1
	nums := []int{0, 7, 3, 2, 1}
	queries := [][]int{{0, 3}, {1, 3}, {2, 4}}
	fmt.Println(maximumXORScoreSubarrayQueries(nums, queries))
	// Example 2
	nums2 := []int{1, 2, 3, 4}
	queries2 := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	fmt.Println(maximumXORScoreSubarrayQueries(nums2, queries2))
	// Example 3
	nums3 := []int{5, 8, 13}
	queries3 := [][]int{{0, 2}}
	fmt.Println(maximumXORScoreSubarrayQueries(nums3, queries3))
	// Example 4: single element
	nums4 := []int{10}
	queries4 := [][]int{{0, 0}}
	fmt.Println(maximumXORScoreSubarrayQueries(nums4, queries4))
}

// Trie node for XOR maximization (binary trie for up to 20 bits).
type xorTrieNode struct {
	child [2]*xorTrieNode
}

func insertXorTrie(root *xorTrieNode, val int) {
	node := root
	for b := 20; b >= 0; b-- {
		bit := (val >> b) & 1
		if node.child[bit] == nil {
			node.child[bit] = &xorTrieNode{}
		}
		node = node.child[bit]
	}
}

func queryMaxXor(root *xorTrieNode, val int) int {
	node := root
	ans := 0
	for b := 20; b >= 0; b-- {
		bit := (val >> b) & 1
		want := 1 - bit
		if node.child[want] != nil {
			ans |= (1 << b)
			node = node.child[want]
		} else {
			node = node.child[bit]
		}
	}
	return ans
}

func maximumXORScoreSubarrayQueries(nums []int, queries [][]int) []int {
	n := len(nums)

	// Prefix XOR: pref[0] = 0, pref[i] = nums[0] ^ ... ^ nums[i-1].
	pref := make([]int, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] ^ v
	}

	// dp[l][r] = max XOR of any subarray within [l, r].
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// For each right endpoint r, build a trie of pref[l..r+1]
	// and compute max subarray XOR for each possible left endpoint.
	for r := 0; r < n; r++ {
		trie := &xorTrieNode{}
		// Insert pref[r+1] to represent the empty subarray ending at r.
		insertXorTrie(trie, pref[r+1])
		bestEnding := 0
		for l := r; l >= 0; l-- {
			// Insert pref[l] into the trie.
			insertXorTrie(trie, pref[l])
			// Query max XOR of pref[l] with any pref in (l, r+1].
			// This gives the max XOR subarray starting at l and ending in [l, r].
			curXor := queryMaxXor(trie, pref[l])
			if curXor > bestEnding {
				bestEnding = curXor
			}
			// dp[l][r] = max(subarrays ending at r with start >= l, subarrays entirely within [l, r-1]).
			dp[l][r] = bestEnding
			if r > 0 && dp[l][r-1] > dp[l][r] {
				dp[l][r] = dp[l][r-1]
			}
		}
	}

	// Answer each query.
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = dp[q[0]][q[1]]
	}
	return ans
}
