# 3083 — Existence Of A Substring In A String And Its Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ExistenceOfASubstringInAStringAndItsReverse(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3083: Existence of a Substring in a String and Its Reverse
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isSubstringPresent
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("leetcode")) // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcba"))   // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcd"))    // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isSubstringPresent
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool {
	// Build set of all substrings of length 2
  // HashMap: O(1) lookup
	substrings := make(map[string]bool)
  // Linear scan O(n)
	for i := 0; i < len(s)-1; i++ {
		substrings[s[i:i+2]] = true
	}

	// Check reverse for any of those substrings
	for i := len(s) - 1; i > 0; i-- {
		if substrings[string(s[i])+string(s[i-1])] {
			return true
		}
	}
	return false
}
```
