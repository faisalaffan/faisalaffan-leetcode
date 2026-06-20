# 0409 — Longest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestPalindrome(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #409: Longest Palindrome
// https://leetcode.com/problems/longest-palindrome/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LongestPalindrome(s string) int {
	count := [128]int{}
	for _, c := range s {
		count[c]++
	}
	length, odd := 0, 0
	for _, c := range count {
		length += (c / 2) * 2
		if c%2 == 1 {
			odd = 1
		}
	}
	return length + odd
}

func main() {
	fmt.Println(LongestPalindrome("abccccdd"))
	fmt.Println(LongestPalindrome("a"))
	fmt.Println(LongestPalindrome("bb"))
}
```
