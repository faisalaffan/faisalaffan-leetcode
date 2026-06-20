# 0521 — Longest Uncommon Subsequence I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestUncommonSubsequenceI(a, b string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #521: Longest Uncommon Subsequence I
// https://leetcode.com/problems/longest-uncommon-subsequence-i/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func LongestUncommonSubsequenceI(a, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(LongestUncommonSubsequenceI("aba", "cdc"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "bbb"))
	fmt.Println(LongestUncommonSubsequenceI("aaa", "aaa"))
}
```
