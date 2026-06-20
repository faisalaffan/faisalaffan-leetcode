# 3536 — Maximum Product Of Two Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumProductOfTwoDigits(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3536: Maximum Product of Two Digits
// https://leetcode.com/problems/maximum-product-of-two-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoDigits(34))
	fmt.Println(MaximumProductOfTwoDigits(10))
	fmt.Println(MaximumProductOfTwoDigits(99))
}

// MaximumProductOfTwoDigits returns the maximum product of any two digits in n.
// Time: O(log n). Space: O(1).
func MaximumProductOfTwoDigits(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	maxProd := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			prod := digits[i] * digits[j]
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	return maxProd
}
```
