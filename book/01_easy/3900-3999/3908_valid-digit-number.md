# 3908 — Valid Digit Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidDigitNumber(n int, x int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3908: Valid Digit Number
// https://leetcode.com/problems/valid-digit-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ValidDigitNumber(101, 0))
	fmt.Println(ValidDigitNumber(232, 2))
	fmt.Println(ValidDigitNumber(5, 1))
}

// Time: O(log n)
// Space: O(1)
func ValidDigitNumber(n int, x int) bool {
	// Check if n starts with digit x
	first := n
	for first >= 10 {
		first /= 10
	}
	if first == x {
		return false
	}

	// Check if n contains digit x
	m := n
	for m > 0 {
		if m%10 == x {
			return true
		}
		m /= 10
	}
	return false
}
```
