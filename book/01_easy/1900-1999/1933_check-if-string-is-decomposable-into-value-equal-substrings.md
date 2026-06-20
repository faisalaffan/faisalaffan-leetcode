# 1933 — Check If String Is Decomposable Into Value Equal Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1933: Check if String Is Decomposable Into Value-Equal Substrings
// https://leetcode.com/problems/check-if-string-is-decomposable-into-value-equal-substrings/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("000111000"))   // false
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("00011111222"))  // true
	fmt.Println(CheckIfStringIsDecomposableIntoValueEqualSubstrings("011100022233")) // false
}

// Time: O(n), Space: O(1)
func CheckIfStringIsDecomposableIntoValueEqualSubstrings(s string) bool {
	hasGroupOfTwo := false
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		count := j - i
		if count%3 == 1 {
			return false
		}
		if count%3 == 2 {
			if hasGroupOfTwo {
				return false
			}
			hasGroupOfTwo = true
		}
		i = j
	}
	return hasGroupOfTwo
}
```
