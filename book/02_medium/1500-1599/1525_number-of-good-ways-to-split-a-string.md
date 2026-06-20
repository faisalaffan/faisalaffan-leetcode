# 1525 — Number Of Good Ways To Split A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumSplits(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N), Space: O(1) (26 chars)  
**Kompleksitas Ruang:** O(1) (26 chars)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1525: Number of Good Ways to Split a String
// https://leetcode.com/problems/number-of-good-ways-to-split-a-string/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumSplits("aacaba"))
	fmt.Println(NumSplits("abcd"))
	fmt.Println(NumSplits("aaaaa"))
}

func NumSplits(s string) int {
	// Time: O(N), Space: O(1) (26 chars)
	n := len(s)
  // Alokasi slice integer
	leftCount := make([]int, 26)
  // Alokasi slice integer
	rightCount := make([]int, 26)
	leftUnique := 0
	rightUnique := 0

	// Initialize right side
	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		if rightCount[idx] == 0 {
			rightUnique++
		}
		rightCount[idx]++
	}

	result := 0
	for i := 0; i < n-1; i++ {
		idx := s[i] - 'a'
		// Move char from right to left
		if leftCount[idx] == 0 {
			leftUnique++
		}
		leftCount[idx]++

		rightCount[idx]--
		if rightCount[idx] == 0 {
			rightUnique--
		}

		if leftUnique == rightUnique {
			result++
		}
	}

	return result
}
```
