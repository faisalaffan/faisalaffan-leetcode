package main

// LeetCode #2791: Count Paths That Can Form a Palindrome in a Tree
// https://leetcode.com/problems/count-paths-that-can-form-a-palindrome-in-a-tree/
// Difficulty: Hard
//
// DFS + bitmask. Characters are on edges (s[i] = edge char from parent[i] to i).
// XOR mask from root to node. Path(u,v) XOR = mask[u] ^ mask[v]. A palindrome
// requires at most 1 bit set in the path XOR.
// O(N * 26) time, O(N) space.

import "fmt"

func countPalindromePaths(parent []int, s string) int {
	n := len(parent)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	maskCount := make(map[int]int)
	maskCount[0] = 1 // empty path from root to itself
	result := 0

	var dfs func(node int, mask int)
	dfs = func(node int, mask int) {
		for _, child := range children[node] {
			edgeMask := 1 << (s[child] - 'a')
			childMask := mask ^ edgeMask

			// Paths with even parity (XOR = 0)
			result += maskCount[childMask]

			// Paths with exactly one odd character (XOR has 1 bit set)
			for b := 0; b < 26; b++ {
				needed := childMask ^ (1 << b)
				result += maskCount[needed]
			}

			maskCount[childMask]++
			dfs(child, childMask)
			maskCount[childMask]--
		}
	}

	dfs(0, 0)
	return result
}

func main() {
	// Example 1: parent=[-1,0,0,1,1,2], s="acaabc" => 8
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 1, 1, 2}, "acaabc"))
	// Example 2: parent=[-1,0,0,0,0], s="aaaaa" => 10
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 0, 0}, "aaaaa"))
	// Single node (no edges)
	fmt.Println(countPalindromePaths([]int{-1}, "a"))
	// Two nodes
	fmt.Println(countPalindromePaths([]int{-1, 0}, "aa"))
	fmt.Println(countPalindromePaths([]int{-1, 0}, "ab"))
	// User's example
	fmt.Println(countPalindromePaths([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}
