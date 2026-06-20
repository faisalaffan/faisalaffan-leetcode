# 0828 — Count Unique Characters Of All Substrings Of A Given String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func uniqueLetterString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #828: Count Unique Characters of All Substrings of a Given String
// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/
// Difficulty: Hard
// Approach: Contribution per character. For each s[i], count substrings where s[i] is the
// first occurrence of that character within the substring. Use prev/next occurrence arrays.

import "fmt"

func uniqueLetterString(s string) int {
	n := len(s)
  // Alokasi slice
	prev := make([]int, n)
  // Alokasi slice
	next := make([]int, n)
  // Alokasi slice
	last := make([]int, 26)

  // Range loop
	for i := range last {
		last[i] = -1
	}
	for i := 0; i < n; i++ {
		c := int(s[i] - 'A')
		prev[i] = last[c]
		last[c] = i
	}
  // Range loop
	for i := range last {
		last[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		c := int(s[i] - 'A')
		next[i] = last[c]
		last[c] = i
	}

	ans := 0
	for i := 0; i < n; i++ {
		left := i - prev[i]
		right := next[i] - i
		ans += left * right
	}
	return ans
}

func main() {
	fmt.Println(uniqueLetterString("ABC")) // Expected: 10
	fmt.Println(uniqueLetterString("ABA")) // Expected: 8
	fmt.Println(uniqueLetterString("LEETCODE")) // Additional test
}
```
