# 0522 — Longest Uncommon Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindLUSlength(strs []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2 * L) where L is max length  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #522: Longest Uncommon Subsequence II
// https://leetcode.com/problems/longest-uncommon-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2 * L) where L is max length
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(FindLUSlength([]string{"aaa", "aaa", "aa"}))
}

func FindLUSlength(strs []string) int {
	maxLen := -1

  // Linear scan O(n)
	for i := 0; i < len(strs); i++ {
		isUnique := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSubseq(strs[i], strs[j]) {
				isUnique = false
				break
			}
		}
		if isUnique && len(strs[i]) > maxLen {
			maxLen = len(strs[i])
		}
	}

	return maxLen
}

func isSubseq(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
```
