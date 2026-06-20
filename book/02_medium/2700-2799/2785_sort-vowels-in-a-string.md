# 2785 — Sort Vowels In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SortVowelsInAString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2785: Sort Vowels in a String
// https://leetcode.com/problems/sort-vowels-in-a-string/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func SortVowelsInAString(s string) string {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	vowels := make([]byte, 0)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

  // Custom sort
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	result := []byte(s)
	vi := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = vowels[vi]
			vi++
		}
	}

	return string(result)
}

func main() {
	fmt.Println(SortVowelsInAString("lEetcOde"))
	fmt.Println(SortVowelsInAString("lYmpH"))
}
```
