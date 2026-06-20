# 2278 — Percentage Of Letter In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PercentageOfLetterInString(s string, letter byte) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2278: Percentage of Letter in String
// https://leetcode.com/problems/percentage-of-letter-in-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PercentageOfLetterInString("foobar", 'o')) // 33
	fmt.Println(PercentageOfLetterInString("jjjj", 'k'))   // 0
	fmt.Println(PercentageOfLetterInString("sgawtb", 's')) // 16
}

func PercentageOfLetterInString(s string, letter byte) int {
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return count * 100 / len(s)
}
```
