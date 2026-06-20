# 0520 — Detect Capital

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func DetectCapital(word string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #520: Detect Capital
// https://leetcode.com/problems/detect-capital/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func DetectCapital(word string) bool {
	upperCount := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			upperCount++
		}
	}
	return upperCount == len(word) || upperCount == 0 || (upperCount == 1 && word[0] >= 'A' && word[0] <= 'Z')
}

func main() {
	fmt.Println(DetectCapital("USA"))
	fmt.Println(DetectCapital("FlaG"))
	fmt.Println(DetectCapital("Google"))
}
```
