# 0343 — Integer Break

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func integerBreak(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #343: Integer Break
// https://leetcode.com/problems/integer-break/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}
	product *= n
	return product
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", integerBreak(2))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", integerBreak(10))
	// Expected: 36

	// Test case 3
	fmt.Println("Test 3:", integerBreak(5))
	// Expected: 6
}
```
