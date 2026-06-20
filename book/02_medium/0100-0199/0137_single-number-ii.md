# 0137 — Single Number Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func singleNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #137: Single Number II
// https://leetcode.com/problems/single-number-ii/
// Difficulty: Medium

import "fmt"

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}
	return ones
}

func main() {
	// Test case 1
	fmt.Println(singleNumber([]int{2, 2, 3, 2})) // 3

	// Test case 2
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99})) // 99

	// Test case 3
	fmt.Println(singleNumber([]int{30000, 500, 100, 30000, 100, 30000, 100})) // 500
}

// Time: O(n) | Space: O(1)
```
