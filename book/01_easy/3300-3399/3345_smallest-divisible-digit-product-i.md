# 3345 — Smallest Divisible Digit Product I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func digitProduct(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(answer * log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3345: Smallest Divisible Digit Product I
// https://leetcode.com/problems/smallest-divisible-digit-product-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestDivisibleDigitProductI(10, 2))
	fmt.Println(SmallestDivisibleDigitProductI(15, 3))
}

// digitProduct returns the product of digits of n.
func digitProduct(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n /= 10
	}
	return product
}

// SmallestDivisibleDigitProductI returns the smallest number >= n whose digit product is divisible by t.
// Time: O(answer * log n). Space: O(1).
func SmallestDivisibleDigitProductI(n int, t int) int {
	for {
		dp := digitProduct(n)
		if dp%t == 0 {
			return n
		}
		n++
	}
}
```
