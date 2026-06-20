# 0400 — Nth Digit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findNthDigit(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #400: Nth Digit
// https://leetcode.com/problems/nth-digit/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	// 1-9: 9 digits,  10-99: 90*2 digits,  100-999: 900*3 digits
	length := 1
	count := 9
	start := 1

	for n > length*count {
		n -= length * count
		length++
		count *= 10
		start *= 10
	}

	// Find the actual number
	num := start + (n-1)/length
	// Find the digit within the number
	digitIdx := (n - 1) % length
	digitStr := strconv.Itoa(num)
	return int(digitStr[digitIdx] - '0')
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findNthDigit(3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", findNthDigit(11))
	// Expected: 0 (from 10)

	// Test case 3
	fmt.Println("Test 3:", findNthDigit(190))
	// Expected: 1 (from 100)
}
```
