# 0263 — Ugly Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func IsUgly(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #263: Ugly Number
// https://leetcode.com/problems/ugly-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range []int{2, 3, 5} {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

func main() {
	fmt.Println(IsUgly(6))
	fmt.Println(IsUgly(1))
	fmt.Println(IsUgly(14))
}
```
