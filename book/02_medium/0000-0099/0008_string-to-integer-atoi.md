# 0008 — String To Integer Atoi

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func myAtoi(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #8: String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i, n := 0, len(s)

	// Skip leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	// Handle sign
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func main() {
	// Test case 1
	fmt.Println(myAtoi("42")) // 42

	// Test case 2
	fmt.Println(myAtoi("   -042")) // -42

	// Test case 3
	fmt.Println(myAtoi("1337c0d3")) // 1337

	// Test case 4
	fmt.Println(myAtoi("0-1")) // 0

	// Test case 5
	fmt.Println(myAtoi("words and 987")) // 0
}

// Time: O(n) | Space: O(1)
```
