# 0906 — Super Palindromes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func isPalindrome(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #906: Super Palindromes
// https://leetcode.com/problems/super-palindromes/
// Difficulty: Hard
// Generate palindromes up to sqrt(R), square them, check if square
// is palindrome and within [L, R].

import (
	"fmt"
	"math"
	"strconv"
)

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func superpalindromesInRange(left string, right string) int {
	L, _ := strconv.ParseInt(left, 10, 64)
	R, _ := strconv.ParseInt(right, 10, 64)
	limit := int(math.Sqrt(float64(R))) + 1
	count := 0

	// Generate odd-length palindromes
	for seed := 1; ; seed++ {
		s := strconv.Itoa(seed)
		// odd length: seed + reverse(seed[:len-1])
		runes := []rune(s)
		n := len(runes)
		palRunes := make([]rune, 2*n-1)
		for i := 0; i < n; i++ {
			palRunes[i] = runes[i]
		}
		for i := 0; i < n-1; i++ {
			palRunes[n+i] = runes[n-2-i]
		}
		palStr := string(palRunes)
		pal, _ := strconv.ParseInt(palStr, 10, 64)
		if pal > int64(limit) {
			break
		}
		sq := pal * pal
		if sq >= L && sq <= R && isPalindrome(strconv.FormatInt(sq, 10)) {
			count++
		}
	}

	// Generate even-length palindromes
	for seed := 1; ; seed++ {
		s := strconv.Itoa(seed)
		// even length: seed + reverse(seed)
		runes := []rune(s)
		n := len(runes)
		palRunes := make([]rune, 2*n)
		for i := 0; i < n; i++ {
			palRunes[i] = runes[i]
		}
		for i := 0; i < n; i++ {
			palRunes[n+i] = runes[n-1-i]
		}
		palStr := string(palRunes)
		pal, _ := strconv.ParseInt(palStr, 10, 64)
		if pal > int64(limit) {
			break
		}
		sq := pal * pal
		if sq >= L && sq <= R && isPalindrome(strconv.FormatInt(sq, 10)) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(superpalindromesInRange("4", "1000"))  // Expected: 4
	fmt.Println(superpalindromesInRange("1", "2"))     // Expected: 1
	fmt.Println(superpalindromesInRange("1", "1"))     // Expected: 1
	fmt.Println(superpalindromesInRange("100", "10000")) // Expected: 2 (121, 484, 676 -> 4^2=16, 11^2=121, 22^2=484, 26^2=676, 101^2=10201 >10000?)
}
```
