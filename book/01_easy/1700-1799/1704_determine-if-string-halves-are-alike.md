# 1704 — Determine If String Halves Are Alike

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func HalvesAreAlike(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1704: Determine if String Halves Are Alike
// https://leetcode.com/problems/determine-if-string-halves-are-alike/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func HalvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	mid := len(s) / 2
	count := 0
	for i := 0; i < mid; i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[i+mid]] {
			count--
		}
	}
	return count == 0
}

func main() {
	fmt.Println(HalvesAreAlike("book"))
	fmt.Println(HalvesAreAlike("textbook"))
}
```
