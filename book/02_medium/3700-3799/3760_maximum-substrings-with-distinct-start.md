# 3760 — Maximum Substrings With Distinct Start

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maximumSubstringsWithDistinctStart(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3760: Maximum Substrings With Distinct Start
// https://leetcode.com/problems/maximum-substrings-with-distinct-start/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSubstringsWithDistinctStart(s string) int {
	seen := [26]bool{}
	ans := 0
	for _, ch := range s {
		idx := ch - 'a'
		if !seen[idx] {
			seen[idx] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumSubstringsWithDistinctStart("abacaba"))
	fmt.Println(maximumSubstringsWithDistinctStart("aaaa"))
	fmt.Println(maximumSubstringsWithDistinctStart("abc"))
}
```
