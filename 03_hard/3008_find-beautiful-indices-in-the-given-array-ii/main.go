package main

// LeetCode #3008: Find Beautiful Indices in the Given Array II
// https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/
// Difficulty: Hard
//
// Given a string s and two patterns a and b, find all indices i such that:
//   1. s[i:i+len(a)] == a (i is a starting index of a in s)
//   2. There exists some index j such that s[j:j+len(b)] == b and |i - j| <= k
//
// Approach: KMP + binary search
//   Use KMP to find all starting positions of 'a' and 'b' in s.
//   For each a-position, binary search in b-positions to find one within distance k.

import (
	"fmt"
	"sort"
)

func beautifulIndices(s, a, b string, k int) []int {
	posA := kmpSearch(s, a)
	posB := kmpSearch(s, b)

	ans := make([]int, 0)
	for _, i := range posA {
		// Binary search for the first b-position >= i-k
		idx := sort.SearchInts(posB, i-k)
		if idx < len(posB) && posB[idx] <= i+k {
			ans = append(ans, i)
		}
	}
	return ans
}

// kmpSearch returns all starting indices where pattern occurs in text.
func kmpSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}
	m := len(pattern)

	// Build prefix function (LPS array)
	pi := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}

	// Search
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
	// 'my' at indices 18 and 33, 'squirrel' at 4 and 21
	// 18: |18-21|=3 <= 15 -> beautiful
	// 33: |33-21|=12 <= 15 -> beautiful
	fmt.Println("Test 1:", beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15))

	// Simple test
	fmt.Println("Test 2:", beautifulIndices("abcd", "a", "a", 3))

	// Overlapping patterns
	fmt.Println("Test 3:", beautifulIndices("aaaaa", "aa", "aa", 1))

	// No match
	fmt.Println("Test 4:", beautifulIndices("abc", "d", "e", 1))

	// Long distance
	fmt.Println("Test 5:", beautifulIndices("abcxyz", "a", "x", 5))
	fmt.Println("Test 6:", beautifulIndices("abcxyz", "a", "x", 2))

	// Edge: pattern not found
	fmt.Println("Test 7:", beautifulIndices("hello", "x", "y", 1))
}
