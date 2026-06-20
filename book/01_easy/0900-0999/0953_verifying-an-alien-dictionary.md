# 0953 — Verifying An Alien Dictionary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func isAlienSorted(words []string, order string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #953: Verifying an Alien Dictionary
// https://leetcode.com/problems/verifying-an-alien-dictionary/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")) // false
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz")) // false
}

// isAlienSorted checks if words are sorted in the alien language order.
// Time: O(n * m). Space: O(1).
func isAlienSorted(words []string, order string) bool {
  // HashMap: O(1) lookup
	orderMap := make(map[byte]int)
  // Linear scan O(n)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}
	for i := 1; i < len(words); i++ {
		if !isLess(words[i-1], words[i], orderMap) {
			return false
		}
	}
	return true
}

func isLess(a, b string, order map[byte]int) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if order[a[i]] < order[b[i]] {
			return true
		}
		if order[a[i]] > order[b[i]] {
			return false
		}
	}
	return len(a) <= len(b)
}
```
