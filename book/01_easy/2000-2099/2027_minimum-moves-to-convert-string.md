# 2027 — Minimum Moves To Convert String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumMovesToConvertString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2027: Minimum Moves to Convert String
// https://leetcode.com/problems/minimum-moves-to-convert-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToConvertString("XXX"))       // 1
	fmt.Println(MinimumMovesToConvertString("XXOX"))      // 2
	fmt.Println(MinimumMovesToConvertString("OOOO"))      // 0
}

// Time: O(n), Space: O(1)
func MinimumMovesToConvertString(s string) int {
	moves := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			moves++
			i += 3
		} else {
			i++
		}
	}
	return moves
}
```
