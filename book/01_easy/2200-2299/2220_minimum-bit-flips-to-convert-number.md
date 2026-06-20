# 2220 — Minimum Bit Flips To Convert Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumBitFlipsToConvertNumber(start int, goal int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2220: Minimum Bit Flips to Convert Number
// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumBitFlipsToConvertNumber(10, 7))  // 3
	fmt.Println(MinimumBitFlipsToConvertNumber(3, 4))   // 3
}

// Time: O(1), Space: O(1)
func MinimumBitFlipsToConvertNumber(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor > 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}
```
