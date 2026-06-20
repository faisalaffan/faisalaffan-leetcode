# 2259 — Remove Digit From Number To Maximize Result

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func RemoveDigitFromNumberToMaximizeResult(number string, digit byte) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2259: Remove Digit From Number to Maximize Result
// https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("123", '3'))  // "12"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("1231", '1')) // "231"
	fmt.Println(RemoveDigitFromNumberToMaximizeResult("551", '5'))  // "51"
}

func RemoveDigitFromNumberToMaximizeResult(number string, digit byte) string {
	best := ""
  // Linear scan O(n)
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			candidate := number[:i] + number[i+1:]
			if candidate > best {
				best = candidate
			}
		}
	}
	return best
}
```
