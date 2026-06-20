# 0136 — Single Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func SingleNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #136: Single Number
// https://leetcode.com/problems/single-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func SingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func main() {
	fmt.Println(SingleNumber([]int{2, 2, 1}))
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(SingleNumber([]int{1}))
}
```
