# 3227 — Vowels Game In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func doesAliceWin(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3227: Vowels Game in a String
// https://leetcode.com/problems/vowels-game-in-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesAliceWin(s string) bool {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // Expected: true
	fmt.Println(doesAliceWin("bbcd"))       // Expected: false
}
```
