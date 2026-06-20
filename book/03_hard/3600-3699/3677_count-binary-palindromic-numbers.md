# 3677 — Count Binary Palindromic Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countBinaryPalindromes(n int64) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3677: Count Binary Palindromic Numbers
// https://leetcode.com/problems/count-binary-palindromic-numbers/
// Difficulty: Hard
//
// Count numbers 0 <= k <= n whose binary representation is a palindrome.
//
// Approach: Count palindromes with fewer bits combinatorially,
// then generate same-bit-length candidates by mirroring prefix.

import "fmt"
import "math/bits"
import "strconv"

func main() {
	// Example 1
	fmt.Println(countBinaryPalindromes(9))
	// Example 2
	fmt.Println(countBinaryPalindromes(0))
	// Example 3
	fmt.Println(countBinaryPalindromes(15))
	// Edge: power of 2
	fmt.Println(countBinaryPalindromes(8))
}

func countBinaryPalindromes(n int64) int {
	if n < 0 {
		return 0
	}
  // Edge case: input kosong
	if n == 0 {
		return 1
	}

	bits := int64(bits.Len64(uint64(n))) // number of bits in n

	// Count all binary palindromes with fewer bits
	var total int64
	for b := int64(1); b < bits; b++ {
		total += countPalindromesWithBits(int(b))
	}

	// Count binary palindromes with exactly `bits` bits that are <= n
	total += countPalindromesSameBits(n, int(bits))

	return int(total)
}

// Count binary palindromes with exactly b bits (leading bit is 1)
func countPalindromesWithBits(b int) int64 {
	if b <= 0 {
		return 0
	}
	if b == 1 {
		return 1 // "1"
	}
	// For a palindrome with b bits, first and last bit are 1 (fixed)
	// For the remaining b-2 bits, we need to fill (b-2)/2 positions
	halfLen := (b - 2 + 1) / 2
	// Each position can be 0 or 1
	return 1 << uint(halfLen)
}

// Count binary palindromes with exactly `bits` bits that are <= n
func countPalindromesSameBits(n int64, b int) int64 {
	s := strconv.FormatInt(n, 2)
	// For odd length, middle bit can be anything
	// For even length, mirror exactly
	halfLen := (b + 1) / 2

	prefix := s[:halfLen]
	prefixVal, _ := strconv.ParseInt(prefix, 2, 64)

	var count int64
	// Try all prefixes from 0 to prefixVal
	for p := int64(0); p <= prefixVal; p++ {
		// Construct palindrome from prefix
		pal := constructPalindrome(p, b)
		if pal <= n {
			count++
		}
	}
	return count
}

func constructPalindrome(prefix int64, bits int) int64 {
	prefixStr := strconv.FormatInt(prefix, 2)
	// Pad prefix to correct length
	for len(prefixStr) < (bits+1)/2 {
		prefixStr = "0" + prefixStr
	}
	// Ensure first bit is 1 for the correct bit length
	if len(prefixStr) > (bits+1)/2 {
		prefixStr = prefixStr[len(prefixStr)-(bits+1)/2:]
	}
	// Mirror
	runes := []rune(prefixStr)
	// For even bits: mirror fully; for odd: mirror except middle
	start := 0
	if bits%2 == 1 {
		start = len(runes) - 2
	} else {
		start = len(runes) - 1
	}
	for i := start; i >= 0; i-- {
		prefixStr += string(runes[i])
	}
	val, _ := strconv.ParseInt(prefixStr, 2, 64)
	return val
}
```
