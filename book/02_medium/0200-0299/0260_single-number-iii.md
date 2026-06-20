# 0260 — Single Number Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func singleNumber(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #260: Single Number III
// https://leetcode.com/problems/single-number-iii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	diff := xor & -xor

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&diff == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{-1, 0}))
	fmt.Println(singleNumber([]int{0, 1}))
}
```
