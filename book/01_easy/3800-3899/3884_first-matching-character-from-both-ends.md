# 3884 — First Matching Character From Both Ends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FirstMatchingCharacterFromBothEnds(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3884: First Matching Character From Both Ends
// https://leetcode.com/problems/first-matching-character-from-both-ends/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FirstMatchingCharacterFromBothEnds("abcacbd"))
	fmt.Println(FirstMatchingCharacterFromBothEnds("abc"))
	fmt.Println(FirstMatchingCharacterFromBothEnds("abcdab"))
}

// Time: O(n)
// Space: O(1)
func FirstMatchingCharacterFromBothEnds(s string) int {
	n := len(s)
	for i := 0; i <= n/2; i++ {
		if s[i] == s[n-1-i] {
			return i
		}
	}
	return -1
}
```
