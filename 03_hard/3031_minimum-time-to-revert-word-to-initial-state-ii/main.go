package main

// LeetCode #3031: Minimum Time to Revert Word to Initial State II
// https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/
// Difficulty: Hard
//
// Approach: Z-algorithm (linear-time pattern matching)
// After each operation we remove the first k characters and append k arbitrary
// characters. The word returns to its initial state if, at some time t,
// the suffix starting at position t*k matches the prefix of the word.
// Use the Z-array to find the longest common prefix between word and
// each suffix starting at position pos = t*k.

import "fmt"

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min2(r-i+1, z[i-l])
		}
		for i+z[i] < n && word[z[i]] == word[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	for t := 1; t*k < n; t++ {
		pos := t * k
		if z[pos] >= n-pos {
			return t
		}
	}
	return (n + k - 1) / k
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example: "abaca", k=3 -> 3
	fmt.Println(minimumTimeToInitialState("abaca", 3))
	// Example: "abacaba", k=3 -> 2
	fmt.Println(minimumTimeToInitialState("abacaba", 3))
	// Example: "abacaba", k=2 -> 4
	fmt.Println(minimumTimeToInitialState("abacaba", 2))
}
