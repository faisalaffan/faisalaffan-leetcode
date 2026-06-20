# 2240 — Number Of Ways To Buy Pens And Pencils

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(total / cost1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2240: Number of Ways to Buy Pens and Pencils
// https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/
// Difficulty: Medium
// Time: O(total / cost1) | Space: O(1)

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ways int64 = 0
	for pens := 0; pens*cost1 <= total; pens++ {
		remaining := total - pens*cost1
		ways += int64(remaining/cost2) + 1
		if cost2 == 0 {
			break
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
	// Expected: 9

	// Test case 2
	fmt.Println(waysToBuyPensPencils(5, 10, 10))
	// Expected: 1

	// Test case 3
	fmt.Println(waysToBuyPensPencils(100, 1, 1))
	// Expected: 5151
}
```
