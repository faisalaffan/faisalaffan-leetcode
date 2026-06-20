# 0397 — Integer Replacement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func integerReplacement(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #397: Integer Replacement
// https://leetcode.com/problems/integer-replacement/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func integerReplacement(n int) int {
	count := 0
	for n > 1 {
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || n&3 == 1 {
			n--
		} else {
			n++
		}
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", integerReplacement(8))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", integerReplacement(7))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", integerReplacement(4))
	// Expected: 2
}
```
