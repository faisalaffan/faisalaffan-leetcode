# 3461 — Check If Digits Are Equal In String After Operations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2). Space: O(n).  |  **Ruang:** O(n).


## 💻 Solusi Go

```go
package main

// LeetCode #3461: Check If Digits Are Equal in String After Operations I
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1234"))
	fmt.Println(CheckIfDigitsAreEqualInStringAfterOperationsI("1111"))
}

// CheckIfDigitsAreEqualInStringAfterOperationsI repeatedly replaces adjacent digit pairs with (sum % 10) until 2 digits remain, then checks equality.
// Time: O(n^2). Space: O(n).
func CheckIfDigitsAreEqualInStringAfterOperationsI(s string) bool {
  // Alokasi slice
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	for len(digits) > 2 {
  // Alokasi slice
		next := make([]int, len(digits)-1)
  // Linear scan O(n)
		for i := 0; i < len(digits)-1; i++ {
			next[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = next
	}
	return digits[0] == digits[1]
}
```
