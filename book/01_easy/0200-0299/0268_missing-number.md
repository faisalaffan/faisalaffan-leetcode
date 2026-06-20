# 0268 — Missing Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MissingNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #268: Missing Number
// https://leetcode.com/problems/missing-number/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}

func main() {
	fmt.Println(MissingNumber([]int{3, 0, 1}))
	fmt.Println(MissingNumber([]int{0, 1}))
	fmt.Println(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
```
