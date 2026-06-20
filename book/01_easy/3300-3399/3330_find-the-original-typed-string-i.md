# 3330 — Find The Original Typed String I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindTheOriginalTypedStringI(word string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3330: Find the Original Typed String I
// https://leetcode.com/problems/find-the-original-typed-string-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheOriginalTypedStringI("aabbccdd"))
	fmt.Println(FindTheOriginalTypedStringI("aaaa"))
	fmt.Println(FindTheOriginalTypedStringI("abc"))
}

// FindTheOriginalTypedStringI counts possible original strings where adjacent equal characters could be merged.
// Time: O(n). Space: O(1).
func FindTheOriginalTypedStringI(word string) int {
	count := 1
	streak := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			streak++
		} else {
			streak = 1
		}
		if streak >= 2 {
			// If we have at least 2 of the same char consecutively, we can type fewer
			count++
		}
	}
	return count
}
```
