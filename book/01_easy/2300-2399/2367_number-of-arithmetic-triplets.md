# 2367 — Number Of Arithmetic Triplets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfArithmeticTriplets(nums []int, diff int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2367: Number of Arithmetic Triplets
// https://leetcode.com/problems/number-of-arithmetic-triplets/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(NumberOfArithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3)) // 2
	fmt.Println(NumberOfArithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))  // 2
}

func NumberOfArithmeticTriplets(nums []int, diff int) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}
	count := 0
	for _, n := range nums {
		if seen[n+diff] && seen[n+2*diff] {
			count++
		}
	}
	return count
}
```
