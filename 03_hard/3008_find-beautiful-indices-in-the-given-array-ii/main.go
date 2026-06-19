package main

// LeetCode #3008: Find Beautiful Indices in the Given Array II
// https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/
// Difficulty: Hard
//
// Approach: KMP + Two-pointer
// Use KMP to find all occurrences of pattern 'a' and pattern 'b' in string 's'.
// Then use a two-pointer sweep to find every index i (from posA) that has
// a matching index j (from posB) within distance k: |i - j| <= k.

import (
	"fmt"
	"sort"
)

func beautifulIndices(s, a, b string, k int) []int {
	posA := kmpSearch(s, a)
	posB := kmpSearch(s, b)
	ans := make([]int, 0)
	j := 0
	m := len(posB)
	for _, i := range posA {
		// Advance j to the first posB[j] >= i - k
		idx := sort.SearchInts(posB, i-k)
		if idx < len(posB) && posB[idx] <= i+k {
			ans = append(ans, i)
		}
		// Original two-pointer approach (same result)
		_ = j
		_ = m
	}
	return ans
}

func kmpSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}
	m := len(pattern)
	pi := make([]int, m)
	var j int
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	pos := make([]int, 0)
	j = 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			pos = append(pos, i-m+1)
			j = pi[j-1]
		}
	}
	return pos
}

func main() {
	// Example: "isawsquirrelnearmysquirrelhouseohmy", a="my", b="squirrel", k=15
	// Expected: [16, 33]
	fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15))
	// Simple test
	fmt.Println(beautifulIndices("abcd", "a", "a", 3))
}
