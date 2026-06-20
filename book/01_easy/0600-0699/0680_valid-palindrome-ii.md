# 0680 — Valid Palindrome Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func validPalindrome(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #680: Valid Palindrome II
// https://leetcode.com/problems/valid-palindrome-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(validPalindrome("aba"))    // true
	fmt.Println(validPalindrome("abca"))   // true
	fmt.Println(validPalindrome("abc"))    // false
	fmt.Println(validPalindrome("deeee"))  // true
}

// validPalindrome checks if the string can be a palindrome after deleting at most one character.
// Time: O(n). Space: O(1).
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if s[l] != s[r] {
			return isPal(s, l+1, r) || isPal(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPal(s string, l, r int) bool {
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```
