# 2900 — Longest Unequal Adjacent Groups Subsequence I

## Deskripsi

**Soal:** [2900. Longest Unequal Adjacent Groups Subsequence I](https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2900: Longest Unequal Adjacent Groups Subsequence I
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: getWordsInLongestSubsequence
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(3, []string{"e", "a", "b"}, []int{0, 0, 1})) // ["e","b"]
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(4, []string{"a", "b", "c", "d"}, []int{1, 0, 1, 0})) // ["a","b","c","d"]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: getWordsInLongestSubsequence
func LongestUnequalAdjacentGroupsSubsequenceI(n int, words []string, groups []int) []string {
	result := []string{words[0]}
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}
	return result
}
```
