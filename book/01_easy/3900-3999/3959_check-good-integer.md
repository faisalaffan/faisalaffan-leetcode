# 3959 — Check Good Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckGoodInteger(n int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3959: Check Good Integer
// https://leetcode.com/problems/check-good-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckGoodInteger(1000))
	fmt.Println(CheckGoodInteger(19))
}

// Time: O(log n)
// Space: O(1)
func CheckGoodInteger(n int) bool {
	total := 0
	for n > 0 {
		d := n % 10
		total += d * (d - 1)
		n /= 10
	}
	return total >= 50
}
```
