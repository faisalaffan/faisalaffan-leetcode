# 3827 — Count Monobit Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountMonobitIntegers(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3827: Count Monobit Integers
// https://leetcode.com/problems/count-monobit-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountMonobitIntegers(1))
	fmt.Println(CountMonobitIntegers(4))
	fmt.Println(CountMonobitIntegers(0))
}

// Time: O(log n)
// Space: O(1)
func CountMonobitIntegers(n int) int {
	ans := 1 // 0 is monobit (all zeros)
	x := 1   // 2^1 - 1 = 1 (all ones)
	for x <= n {
		ans++
		x = x*2 + 1
	}
	return ans
}
```
