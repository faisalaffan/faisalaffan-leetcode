# 1554 — Strings Differ By One Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func DifferByOne(dict []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N*M^2) where N = len(dict), M = string length  |  **Ruang:** O(N*M)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1554: Strings Differ by One Character
// https://leetcode.com/problems/strings-differ-by-one-character/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DifferByOne([]string{"abcd", "acbd", "aacd"}))
	fmt.Println(DifferByOne([]string{"ab", "cd", "yz"}))
	fmt.Println(DifferByOne([]string{"abcd", "cccc", "abxd", "abzd"}))
}

func DifferByOne(dict []string) bool {
	// Time: O(N*M^2) where N = len(dict), M = string length
	// Space: O(N*M)
	// Use rolling hash: for each position, check if any two strings
	// become identical when that position is skipped.

	n := len(dict)
	if n < 2 {
		return false
	}
	m := len(dict[0])

	for skipIdx := 0; skipIdx < m; skipIdx++ {
  // HashMap: O(1) lookup
		seen := make(map[string]bool)
		for i := 0; i < n; i++ {
			// Create string without char at skipIdx
			key := dict[i][:skipIdx] + dict[i][skipIdx+1:]
			if seen[key] {
				return true
			}
			seen[key] = true
		}
	}

	return false
}
```
