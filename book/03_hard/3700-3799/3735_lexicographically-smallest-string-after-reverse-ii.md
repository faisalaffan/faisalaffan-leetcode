# 3735 — Lexicographically Smallest String After Reverse Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func smallestString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3735: Lexicographically Smallest String After Reverse II
// https://leetcode.com/problems/lexicographically-smallest-string-after-reverse-ii/
// Difficulty: Hard [Paid]
//
// Perform exactly one operation: choose k (1 <= k <= n) and either
// reverse the first k characters or reverse the last k characters.
// Return the lexicographically smallest string obtainable.
//
// Approach: Track minimum character positions from left and right.
// For prefix reversal, best k brings a minimal char to position 0.
// For suffix reversal, best k brings a minimal suffix char forward.
// Try all candidates and pick min.

import "fmt"

func main() {
	// Example 1
	fmt.Println(smallestString("acdb"))
	// Example 2
	fmt.Println(smallestString("abba"))
	// Edge: sorted
	fmt.Println(smallestString("abc"))
	// Edge: reverse all
	fmt.Println(smallestString("cba"))
	// Edge: single char
	fmt.Println(smallestString("z"))
}

func smallestString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	best := s

	// Prefix reversal: try best candidates
	// Track minimum char from right to left.
	// When we find a new minimum, that position could give a better result.
	minChar := byte('z' + 1)
	for k := n; k >= 1; k-- {
		ch := s[k-1]
		if ch < minChar {
			minChar = ch
			cand := reverse(s[:k]) + s[k:]
			if cand < best {
				best = cand
			}
		}
	}

	// Suffix reversal: try best candidates.
	// The first affected position is n-k, where we place s[n-1].
	// If s[n-1] < s[n-k], the result improves at that position.
	// Also always try k=n as a candidate.
	lastCh := s[n-1]
	bestK := n
	found := false
	for i := 0; i < n-1; i++ {
		if lastCh < s[i] {
			bestK = n - i
			found = true
			break
		}
	}
	if !found {
		// Try k=n (reverse all) as fallback
		bestK = n
	}

	// Also try k=2 through n for suffix: just check candidates at
	// positions where s[n-1] <= s[n-k] (might still be better due to
	// subsequent characters)
	for k := 2; k <= n; k++ {
		cand := s[:n-k] + reverse(s[n-k:])
		if cand < best {
			best = cand
		}
		// Optimization: if lastCh < s[n-k], this is a "good" reversal
		// and we don't need smaller k (which would affect later positions)
		if lastCh < s[n-k] {
			break
		}
	}
	if found {
		cand := s[:n-bestK] + reverse(s[n-bestK:])
		if cand < best {
			best = cand
		}
	}

	return best
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```
