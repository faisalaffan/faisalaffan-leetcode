# 2847 — Smallest Number With Given Digit Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestNumberWithGivenDigitProduct(n int64) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2847: Smallest Number With Given Digit Product
// https://leetcode.com/problems/smallest-number-with-given-digit-product/
// Difficulty: Medium [Paid]
// Time: O(log n) | Space: O(log n)

import "fmt"

func SmallestNumberWithGivenDigitProduct(n int64) string {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}

	digits := make([]byte, 0)
	for i := 9; i >= 2; i-- {
		for n%int64(i) == 0 {
			digits = append([]byte{byte('0' + i)}, digits...)
			n /= int64(i)
		}
	}

	if n > 1 {
		return "-1"
	}

	return string(digits)
}

func main() {
	fmt.Println(SmallestNumberWithGivenDigitProduct(36))
	fmt.Println(SmallestNumberWithGivenDigitProduct(17))
	fmt.Println(SmallestNumberWithGivenDigitProduct(1))
}
```
