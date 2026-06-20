# 1750 — Minimum Length Of String After Deleting Similar Ends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumLength(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1750: Minimum Length of String After Deleting Similar Ends
// https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minimumLength(s string) int {
	left, right := 0, len(s)-1

  // Two-pointer: gerakkan kiri atau kanan
	for left < right && s[left] == s[right] {
		ch := s[left]
		// Delete from left
		for left <= right && s[left] == ch {
			left++
		}
		// Delete from right
		for left <= right && s[right] == ch {
			right--
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(minimumLength("ca"))            // Expected: 2
	fmt.Println(minimumLength("cabaabac"))      // Expected: 0
	fmt.Println(minimumLength("aabccabba"))     // Expected: 3
}
```
