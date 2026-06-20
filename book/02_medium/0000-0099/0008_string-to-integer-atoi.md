# 0008 — String To Integer Atoi

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func myAtoi(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
