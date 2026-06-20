# 1137 — N Th Tribonacci Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func tribonacci(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1137: N-th Tribonacci Number
// https://leetcode.com/problems/n-th-tribonacci-number/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(tribonacci(4))  // 4
	fmt.Println(tribonacci(25)) // 1389537
	fmt.Println(tribonacci(0))  // 0
}

// LeetCode submission: tribonacci
func tribonacci(n int) int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
```
