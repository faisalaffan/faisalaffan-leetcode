# 1616 — Split Two Strings To Make Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func CheckPalindromeFormation(a string, b string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

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
  // Two-pointer loop
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
  // Two-pointer loop
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
