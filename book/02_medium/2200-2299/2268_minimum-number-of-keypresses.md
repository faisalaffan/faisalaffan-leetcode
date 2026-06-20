# 2268 — Minimum Number Of Keypresses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumKeypresses(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n + 26 log 26)  |  **Ruang:** O(26)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2268: Minimum Number of Keypresses
// https://leetcode.com/problems/minimum-number-of-keypresses/
// Difficulty: Medium [Paid]
// Time: O(n + 26 log 26) | Space: O(26)

import (
	"fmt"
	"sort"
)

func minimumKeypresses(s string) int {
  // Alokasi slice
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

  // Custom sort
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	presses := 0
	for i, c := range count {
		if c == 0 {
			break
		}
		presses += c * (i/9 + 1)
	}
	return presses
}

func main() {
	// Test case 1
	fmt.Println(minimumKeypresses("apple"))
	// Expected: 5

	// Test case 2
	fmt.Println(minimumKeypresses("abcdefghijkl"))
	// Expected: 15

	// Test case 3
	fmt.Println(minimumKeypresses("aaaaaaa"))
	// Expected: 7
}
```
