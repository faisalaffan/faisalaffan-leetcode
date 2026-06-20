# 3803 — Count Residue Prefixes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountResiduePrefixes(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1) — at most 26 distinct chars

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3803: Count Residue Prefixes
// https://leetcode.com/problems/count-residue-prefixes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountResiduePrefixes("abc"))
	fmt.Println(CountResiduePrefixes("dd"))
	fmt.Println(CountResiduePrefixes("bob"))
}

// Time: O(n)
// Space: O(1) — at most 26 distinct chars
func CountResiduePrefixes(s string) int {
  // HashMap: O(1) lookup
	seen := make(map[byte]bool)
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		seen[s[i]] = true
		if len(seen) == (i+1)%3 {
			count++
		}
	}
	return count
}
```
