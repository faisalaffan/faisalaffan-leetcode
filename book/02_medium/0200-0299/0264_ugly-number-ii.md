# 0264 — Ugly Number Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func nthUglyNumber(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #264: Ugly Number II
// https://leetcode.com/problems/ugly-number-ii/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func nthUglyNumber(n int) int {
  // Alokasi slice integer
	ugly := make([]int, n)
	ugly[0] = 1

	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		next := min(ugly[p2]*2, min(ugly[p3]*3, ugly[p5]*5))
		ugly[i] = next

		if next == ugly[p2]*2 {
			p2++
		}
		if next == ugly[p3]*3 {
			p3++
		}
		if next == ugly[p5]*5 {
			p5++
		}
	}

	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1))
	fmt.Println(nthUglyNumber(7))
}
```
