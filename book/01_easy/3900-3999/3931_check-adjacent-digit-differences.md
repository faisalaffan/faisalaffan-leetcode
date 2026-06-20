# 3931 — Check Adjacent Digit Differences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckAdjacentDigitDifferences(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3931: Check Adjacent Digit Differences
// https://leetcode.com/problems/check-adjacent-digit-differences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckAdjacentDigitDifferences("132"))
	fmt.Println(CheckAdjacentDigitDifferences("129"))
}

// Time: O(n)
// Space: O(1)
func CheckAdjacentDigitDifferences(s string) bool {
  // Linear scan O(n)
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		if diff < 0 {
			diff = -diff
		}
		if diff > 2 {
			return false
		}
	}
	return true
}
```
