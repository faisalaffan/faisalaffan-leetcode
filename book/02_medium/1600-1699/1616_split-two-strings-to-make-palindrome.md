# 1616 — Split Two Strings To Make Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckPalindromeFormation(a string, b string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1616: Split Two Strings to Make Palindrome
// https://leetcode.com/problems/split-two-strings-to-make-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CheckPalindromeFormation("x", "y"))
	fmt.Println(CheckPalindromeFormation("abdef", "fecab"))
	fmt.Println(CheckPalindromeFormation("ulacfd", "jizalu"))
}

func CheckPalindromeFormation(a string, b string) bool {
	// Time: O(N), Space: O(1)
	return canForm(a, b) || canForm(b, a)
}

func canForm(a, b string) bool {
	left := 0
	right := len(a) - 1

	// Find the first mismatch from the outside
  // Two-pointer: gerakkan kiri atau kanan
	for left < right && a[left] == b[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}

	// Check if a[left..right] is palindrome
	return isPalindrome(a, left, right) || isPalindrome(b, left, right)
}

func isPalindrome(s string, left, right int) bool {
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```
