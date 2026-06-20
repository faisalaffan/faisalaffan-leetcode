# 3646 — Next Special Palindrome Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah bilangan bulat (integer). Tugasmu adalah menentukan apakah bilangan tersebut adalah **palindrome** — dibaca sama dari depan maupun dari belakang.

Contoh: `121` → palindrome. `-121` → bukan (tanda minus!). `10` → bukan.

**Cara berpikir:** Balik setengah digit secara matematika menggunakan modulo (`%`) dan pembagian (`/`). Tidak perlu konversi ke string.

**Fungsi Solusi:** `func nextSpecialPalindrome(n int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3646: Next Special Palindrome Number
// https://leetcode.com/problems/next-special-palindrome-number/
// Difficulty: Hard
//
// Find the smallest palindrome number strictly greater than n
// that can be formed using only even digits (0, 2, 4, 6, 8).
// A "special palindrome" uses only even digits.
//
// Approach: Generate palindrome candidates from even digits.
// For each possible length starting from len(n)+1, try all
// left halves from even digits and construct palindrome.

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Example 1
	fmt.Println(nextSpecialPalindrome(123))
	// Example 2
	fmt.Println(nextSpecialPalindrome(88))
	// Edge: n = 0
	fmt.Println(nextSpecialPalindrome(0))
	// Edge: already special
	fmt.Println(nextSpecialPalindrome(242))
}

func nextSpecialPalindrome(n int64) int64 {
	if n < 0 {
		return 0
	}

	s := strconv.FormatInt(n, 10)
	m := len(s)

	// Try lengths from m to m+2
	for length := m; length <= m+2; length++ {
		halfLen := (length + 1) / 2
		start := int64(0)
		if halfLen > 0 {
			start = int64(math.Pow10(halfLen - 1))
		}

		limit := int64(math.Pow10(halfLen))

		for left := start; left < limit; left++ {
			leftStr := strconv.FormatInt(left, 10)
			if len(leftStr) < halfLen {
				continue
			}

			// Check all digits are even
			allEven := true
			for _, ch := range leftStr {
				if (ch-'0')%2 != 0 {
					allEven = false
					break
				}
			}
			if !allEven {
				continue
			}

			// Build palindrome
			runes := []byte(leftStr)
			if length%2 == 0 {
				for i := halfLen - 1; i >= 0; i-- {
					runes = append(runes, runes[i])
				}
			} else {
				for i := halfLen - 2; i >= 0; i-- {
					runes = append(runes, runes[i])
				}
			}
			palStr := string(runes)

			pal, _ := strconv.ParseInt(palStr, 10, 64)
			if pal > n {
				return pal
			}
		}
	}

	return 0
}
```
