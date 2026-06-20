# 3498 — Reverse Degree Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ReverseDegreeOfAString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3498: Reverse Degree of a String
// https://leetcode.com/problems/reverse-degree-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseDegreeOfAString("abc"))
	fmt.Println(ReverseDegreeOfAString("zaba"))
}

// ReverseDegreeOfAString computes the sum of (position_in_reversed_alphabet * (i+1)) for each character.
// Reverse: a=26, b=25, ..., z=1.
// Time: O(n). Space: O(1).
func ReverseDegreeOfAString(s string) int {
	sum := 0
	for i, ch := range s {
		revPos := 26 - int(ch-'a')
		sum += revPos * (i + 1)
	}
	return sum
}
```
