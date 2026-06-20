# 0014 — Longest Common Prefix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LongestCommonPrefix(strs []string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n*m)  |  **Ruang:** O(1) where n=len(strs), m=len(shortest string)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #14: Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/
// Difficulty: Easy

import "fmt"

// Time: O(n*m) | Space: O(1) where n=len(strs), m=len(shortest string)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
  // Linear scan O(n)
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"a"}))
}
```
