# 3223 — Minimum Length Of String After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumLength(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(26) = O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3223: Minimum Length of String After Operations
// https://leetcode.com/problems/minimum-length-of-string-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func minimumLength(s string) int {
  // Alokasi slice
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	ans := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumLength("abaacbcbb")) // Expected: 5
	fmt.Println(minimumLength("aa"))         // Expected: 2
}
```
