# 1963 — Minimum Number Of Swaps To Make The String Balanced

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinSwapsBalanced(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1963: Minimum Number of Swaps to Make the String Balanced
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinSwapsBalanced("][]["))
	fmt.Println(MinSwapsBalanced("]]][[["))
	fmt.Println(MinSwapsBalanced("[]"))
}

// Time: O(n), Space: O(1)
func MinSwapsBalanced(s string) int {
	unmatched := 0
	maxUnmatched := 0
	for _, c := range s {
		if c == '[' {
			unmatched--
		} else {
			unmatched++
		}
		if unmatched > maxUnmatched {
			maxUnmatched = unmatched
		}
	}
	return (maxUnmatched + 1) / 2
}
```
