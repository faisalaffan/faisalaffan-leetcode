# 2262 — Total Appeal Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func appealSum(s string) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2262: Total Appeal of A String
// https://leetcode.com/problems/total-appeal-of-a-string/
// Difficulty: Hard
//
// The appeal of a string is the number of distinct characters in it.
// Given a string s, return the total appeal of all its substrings.

import (
	"fmt"
)

// appealSum returns the sum of appeal over all substrings of s.
func appealSum(s string) int64 {
	// For each character, we count how many substrings include it
	// as the FIRST occurrence from the left (to avoid double counting).
	//
	// For position i with character c, the number of substrings where
	// c contributes its distinctness is:
	//   (i - lastSeen[c]) * (n - i)
	// where lastSeen[c] is the previous occurrence of c (-1 if none).
	// This counts substrings that START after last occurrence and END at or after i.

	n := len(s)
  // Membuat map (HashMap) — pencarian O(1)
	lastSeen := make(map[byte]int)
	var result int64

	for i := 0; i < n; i++ {
		c := s[i]
		prev, ok := lastSeen[c]
		if !ok {
			prev = -1
		}
		// substrings starting in (prev, i] and ending >= i
		// left choices: i - prev
		// right choices: n - i
		result += int64(i-prev) * int64(n-i)
		lastSeen[c] = i
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(appealSum("abbca")) // Expected: 28

	// Example 2
	fmt.Println(appealSum("code")) // Expected: 20
}
```
