# 3782 — Last Remaining Integer After Alternating Deletion Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func lastRemaining(n int) int64
```

> **💡 Hint:** Each step removes two elements (one from each end)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3782: Last Remaining Integer After Alternating
// Deletion Operations
// https://leetcode.com/problems/last-remaining-integer-after-
// alternating-deletion-operations/
// Difficulty: Hard
//
// Start with list [1, 2, ..., n]. Delete elements alternately
// from left and right ends until one remains. Return survivor.
//
// Approach: Each step removes two elements (one from each end)
// alternately. The survivor position follows: f(n) = n/2 + 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(lastRemaining(5))
	// Example 2
	fmt.Println(lastRemaining(7))
	// Edge: single element
	fmt.Println(lastRemaining(1))
	// Edge: two elements
	fmt.Println(lastRemaining(2))
	// Edge: large
	fmt.Println(lastRemaining(100))
}

func lastRemaining(n int) int64 {
	if n <= 0 {
		return 0
	}
	// Each full cycle of left-right deletes 2 elements.
	// The survivor is at position floor(n/2) + 1 (1-indexed).
	return int64(n/2 + 1)
}
```
