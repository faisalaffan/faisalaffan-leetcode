# 2900 — Longest Unequal Adjacent Groups Subsequence I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestUnequalAdjacentGroupsSubsequenceI(n int, words []string, groups []int) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

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
