# 3707 — Equal Score Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func EqualScoreSubstrings(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3707: Equal Score Substrings
// https://leetcode.com/problems/equal-score-substrings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(EqualScoreSubstrings("adcb"))
	fmt.Println(EqualScoreSubstrings("bace"))
}

// Time: O(n)
// Space: O(1)
func EqualScoreSubstrings(s string) bool {
	total := 0
	for _, ch := range s {
		total += int(ch-'a') + 1
	}

	prefix := 0
  // Linear scan O(n)
	for i := 0; i < len(s)-1; i++ {
		prefix += int(s[i]-'a') + 1
		if prefix == total-prefix {
			return true
		}
	}
	return false
}
```
