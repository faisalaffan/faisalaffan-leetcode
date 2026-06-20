# 1328 — Break A Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func breakPalindrome(palindrome string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = length of palindrome string  
**Kompleksitas Ruang:** O(n) for the byte array

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1328: Break a Palindrome
// https://leetcode.com/problems/break-a-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(breakPalindrome("abccba")) // "aaccba"

	// Test case 2
	fmt.Println(breakPalindrome("a")) // ""

	// Test case 3
	fmt.Println(breakPalindrome("aa")) // "ab"

	// Test case 4 - all 'a's
	fmt.Println(breakPalindrome("aaa")) // "aab"
}

// Time: O(n) where n = length of palindrome string
// Space: O(n) for the byte array
func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n <= 1 {
		return ""
	}

	bytes := []byte(palindrome)
	// Try to change first non-'a' to 'a' (only in first half to maintain smallest lexicographically)
	for i := 0; i < n/2; i++ {
		if bytes[i] != 'a' {
			bytes[i] = 'a'
			return string(bytes)
		}
	}

	// All characters in first half are 'a', change last character to 'b'
	bytes[n-1] = 'b'
	return string(bytes)
}
```
