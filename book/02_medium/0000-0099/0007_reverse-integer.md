# 0007 — Reverse Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func reverse(x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log₁₀(n))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #7: Reverse Integer
// https://leetcode.com/problems/reverse-integer/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(reverse(123)) // 321

	// Test case 2
	fmt.Println(reverse(-123)) // -321

	// Test case 3
	fmt.Println(reverse(1534236469)) // 0 (overflow)
}

// Time: O(log₁₀(n)) | Space: O(1)
```
