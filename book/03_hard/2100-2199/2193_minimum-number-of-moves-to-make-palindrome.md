# 2193 — Minimum Number Of Moves To Make Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minMovesToMakePalindrome(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2193: Minimum Number of Moves to Make Palindrome
// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/
// Difficulty: Hard
//
// Greedy two-pointer simulation: For each left character, find the matching
// character closest to the right. Bring it to the rightmost position via
// adjacent swaps. For center characters (single occurrence), swap it one step
// to the right and continue (it will migrate to the center).
// O(n^2) time, O(n) space. n <= 2000 per constraints.

import (
	"fmt"
)

func main() {
	// "aabb" -> 2
	fmt.Println(minMovesToMakePalindrome("aabb"))
	// "letelt" -> 2
	fmt.Println(minMovesToMakePalindrome("letelt"))
	// "a" -> 0
	fmt.Println(minMovesToMakePalindrome("a"))
	// "abba" -> 0
	fmt.Println(minMovesToMakePalindrome("abba"))
	// "abcba" -> 0
	fmt.Println(minMovesToMakePalindrome("abcba"))
}

func minMovesToMakePalindrome(s string) int {
	b := []byte(s)
	n := len(b)
	moves := 0
	l, r := 0, n-1

  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		// Find matching character for b[l] from the right side.
		match := r
		for match > l && b[match] != b[l] {
			match--
		}

		if match == l {
			// Center character: appears only once among remaining elements.
			// Swap it one step to the right (it will migrate to center).
			b[l], b[l+1] = b[l+1], b[l]
			moves++
			continue
		}

		// Bring the matched character to position r via adjacent swaps.
		for i := match; i < r; i++ {
			b[i], b[i+1] = b[i+1], b[i]
			moves++
		}

		l++
		r--
	}

	return moves
}
```
