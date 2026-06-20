# 1133 — Largest Unique Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestUniqueNumber(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1133: Largest Unique Number
// https://leetcode.com/problems/largest-unique-number/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})) // 8
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))                // -1
}

// LeetCode submission: largestUniqueNumber
func largestUniqueNumber(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	ans := -1
	for k, v := range count {
		if v == 1 && k > ans {
			ans = k
		}
	}
	return ans
}
```
