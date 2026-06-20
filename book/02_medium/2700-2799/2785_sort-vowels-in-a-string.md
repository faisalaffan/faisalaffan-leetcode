# 2785 — Sort Vowels In A String

## Deskripsi

**Soal:** [2785. Sort Vowels In A String](https://leetcode.com/problems/sort-vowels-in-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func SortVowelsInAString(s string) string`

## Solusi Go

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

  // Membuat slice untuk menyimpan hasil
	vowels := make([]byte, 0)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			vowels = append(vowels, s[i])
		}
	}

	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	result := []byte(s)
	vi := 0
  // Loop standar: indeks 0 sampai n-1
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
