# 0201 — Bitwise And Of Numbers Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func rangeBitwiseAnd(left int, right int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(log max(left, right)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #201: Bitwise AND of Numbers Range
// https://leetcode.com/problems/bitwise-and-of-numbers-range/
// Difficulty: Medium
// Time: O(log max(left, right)), Space: O(1)

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
```
