# 1842 — Next Palindrome Using Same Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func nextPalindrome(num string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1842: Next Palindrome Using Same Digits
// https://leetcode.com/problems/next-palindrome-using-same-digits/
// Difficulty: Hard [Paid]
//
// Given a palindromic number as a string, find the next palindrome using the
// same digits. Since the right half is a mirror of the left half, we only
// need to find the next permutation of the left half, then mirror it back.
//
// If no larger palindrome can be formed, return an empty string.

import "fmt"

func main() {
	// Example 1: "1221" -> "2112"
	fmt.Println(nextPalindrome("1221"))
	// Example 2: "12321" -> "13231" (odd length, center stays)
	fmt.Println(nextPalindrome("12321"))
	// Example 3: "1" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("1"))
	// Example 4: "11" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("11"))
	// Example 5: "121" -> "211" (next left half of "12" is "21", mirror to "211")
	fmt.Println(nextPalindrome("121"))
	// Example 6: "32123" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("32123"))
	// Example 7: "21312" -> "23132"
	fmt.Println(nextPalindrome("21312"))
	// Example 8: "4554" -> "5445"
	fmt.Println(nextPalindrome("4554"))
	// Example 9: "1234321" -> "1243421"
	fmt.Println(nextPalindrome("1234321"))
	// Example 10: "9999" -> "" (no next palindrome)
	fmt.Println(nextPalindrome("9999"))
}

func nextPalindrome(num string) string {
	n := len(num)
	if n <= 1 {
		return ""
	}

	half := n / 2
	// Work with the left half (including middle character for odd length)
	s := []byte(num[:half])
	// If odd length, we may or may not need to include middle
	// Actually, for palindrome, only the left half determines the number.
	// For odd length, center character stays the same, so we work with half.

	// Find next greater permutation of the left half
	// Find the rightmost position where s[i] < s[i+1]
	i := half - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return ""
	}

	// Find the rightmost element greater than s[i]
	j := half - 1
	for s[j] <= s[i] {
		j--
	}

	// Swap
	s[i], s[j] = s[j], s[i]

	// Reverse suffix starting at i+1
	for a, b := i+1, half-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}

	// Build the palindrome
	result := make([]byte, n)
	copy(result, s)
	for idx := 0; idx < half; idx++ {
		result[n-1-idx] = s[idx]
	}
	// For odd length, the middle character is unchanged
	if n%2 == 1 {
		result[half] = num[half]
	}

	return string(result)
}
```
