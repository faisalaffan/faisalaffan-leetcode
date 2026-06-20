# 1048 — Longest String Chain

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestStrChain(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Sorting

**Waktu:** O(n * L^2) where L is max word length  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1048: Longest String Chain
// https://leetcode.com/problems/longest-string-chain/
// Difficulty: Medium
//
// Approach: Sort by length, DP with hash map.
//           For each word, check all possible predecessors by removing one char.
// Time: O(n * L^2) where L is max word length
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"})) // 4
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})) // 5
}

func longestStrChain(words []string) int {
  // Custom sort
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

  // HashMap: O(1) lookup
	dp := make(map[string]int)
	result := 1

	for _, w := range words {
		best := 1
  // Linear scan O(n)
		for i := 0; i < len(w); i++ {
			pred := w[:i] + w[i+1:]
			if val, ok := dp[pred]; ok {
				if val+1 > best {
					best = val + 1
				}
			}
		}
		dp[w] = best
		if best > result {
			result = best
		}
	}

	return result
}
```
