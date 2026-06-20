# 0171 — Excel Sheet Column Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func TitleToNumber(columnTitle string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #171: Excel Sheet Column Number
// https://leetcode.com/problems/excel-sheet-column-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func TitleToNumber(columnTitle string) int {
	result := 0
  // Linear scan O(n)
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(TitleToNumber("A"))
	fmt.Println(TitleToNumber("AB"))
	fmt.Println(TitleToNumber("ZY"))
}
```
