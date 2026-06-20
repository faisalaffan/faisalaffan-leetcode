# 1624 — Largest Substring Between Two Equal Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaxLengthBetweenEqualCharacters(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(1) (since only 26 letters)  |  **Ruang:** O(1) (since only 26 letters)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1624: Largest Substring Between Two Equal Characters
// https://leetcode.com/problems/largest-substring-between-two-equal-characters/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1) (since only 26 letters)
func MaxLengthBetweenEqualCharacters(s string) int {
  // HashMap: O(1) lookup
	firstIndex := make(map[rune]int)
	maxLen := -1
	for i, ch := range s {
		if idx, exists := firstIndex[ch]; exists {
			if i-idx-1 > maxLen {
				maxLen = i - idx - 1
			}
		} else {
			firstIndex[ch] = i
		}
	}
	return maxLen
}

func main() {
	fmt.Println(MaxLengthBetweenEqualCharacters("aa"))
	fmt.Println(MaxLengthBetweenEqualCharacters("abca"))
	fmt.Println(MaxLengthBetweenEqualCharacters("cbzxy"))
}
```
