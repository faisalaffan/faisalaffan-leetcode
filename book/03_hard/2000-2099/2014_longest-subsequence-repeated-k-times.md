# 2014 — Longest Subsequence Repeated K Times

## Deskripsi

**Soal:** [2014. Longest Subsequence Repeated K Times](https://leetcode.com/problems/longest-subsequence-repeated-k-times/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

**Fungsi Solusi:** `func longestSubsequenceRepeatedK(s string, k int) string`

> **Ide Kunci:** BFS generate candidate strings in order of length.

## Solusi Go

```go
package main

// LeetCode #2014: Longest Subsequence Repeated k Times
// https://leetcode.com/problems/longest-subsequence-repeated-k-times/
// Difficulty: Hard
// Approach: BFS generate candidate strings in order of length.
// Count character frequencies, max_uses = freq / k.
// Generate all possible strings up to n/k length.
// For each candidate, check if repeated k times is a subsequence of s.
// Keep the longest.

import "fmt"

func longestSubsequenceRepeatedK(s string, k int) string {
	// Count frequencies
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Max uses for each character
  // Membuat slice untuk menyimpan hasil
	maxUses := make([]int, 26)
	for i := 0; i < 26; i++ {
		maxUses[i] = freq[i] / k
	}

	// Check if str is a subsequence of s
	isSubseq := func(str string) bool {
		j := 0
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}
		return j == len(str)
	}

	// Check if t repeated k times is a subsequence of s
	check := func(t string) bool {
		if len(t) == 0 {
			return false
		}
		concat := ""
		for i := 0; i < k; i++ {
			concat += t
		}
		return isSubseq(concat)
	}

	// BFS to generate candidates
	queue := []string{""}
	best := ""

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for c := 0; c < 26; c++ {
			if maxUses[c] == 0 {
				continue
			}
			next := cur + string(rune('a'+c))
			if !check(next) {
				continue
			}
			queue = append(queue, next)
			if len(next) > len(best) || (len(next) == len(best) && next > best) {
				best = next
			}
		}
	}

	return best
}

func main() {
	// Example: "letsleetcode", k=2 -> "let"
	fmt.Println(longestSubsequenceRepeatedK("letsleetcode", 2))

	// Additional tests
	fmt.Println(longestSubsequenceRepeatedK("aabbaabbaabb", 3))
	fmt.Println(longestSubsequenceRepeatedK("abcd", 2))
}
```
