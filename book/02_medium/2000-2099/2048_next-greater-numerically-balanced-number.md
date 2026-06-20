# 2048 — Next Greater Numerically Balanced Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func nextBeautifulNumber(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n is the number until we find it  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2048: Next Greater Numerically Balanced Number
// https://leetcode.com/problems/next-greater-numerically-balanced-number/
// Difficulty: Medium
// Time: O(n) where n is the number until we find it | Space: O(1)

import "fmt"

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if isBalanced(i) {
			return i
		}
	}
}

func isBalanced(n int) bool {
  // Alokasi slice integer
	digits := make([]int, 10)
	for n > 0 {
		d := n % 10
		digits[d]++
		n /= 10
	}
	for d := 1; d <= 9; d++ {
		if digits[d] > 0 && digits[d] != d {
			return false
		}
	}
	return digits[0] == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", nextBeautifulNumber(1))
	// Expected: 22

	// Test case 2
	fmt.Println("Test 2:", nextBeautifulNumber(1000))
	// Expected: 1333

	// Test case 3
	fmt.Println("Test 3:", nextBeautifulNumber(3000))
	// Expected: 3133
}
```
