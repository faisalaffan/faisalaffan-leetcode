# 3798 — Largest Even Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func LargestEvenNumber(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3798: Largest Even Number
// https://leetcode.com/problems/largest-even-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LargestEvenNumber("1112"))
	fmt.Println(LargestEvenNumber("221"))
	fmt.Println(LargestEvenNumber("1"))
}

// Time: O(n)
// Space: O(n)
func LargestEvenNumber(s string) string {
	i := len(s)
	for i > 0 && s[i-1] == '1' {
		i--
	}
	return s[:i]
}
```
