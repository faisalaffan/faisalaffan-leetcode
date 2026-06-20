# 1523 — Count Odd Numbers In An Interval Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countOdds(low int, high int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1523: Count Odd Numbers in an Interval Range
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
// Difficulty: Easy
//
// LeetCode submission: func countOdds(low int, high int) int

import "fmt"

func main() {
	fmt.Println(CountOddNumbersInAnIntervalRange(3, 7)) // 3
	fmt.Println(CountOddNumbersInAnIntervalRange(8, 10)) // 1
	fmt.Println(CountOddNumbersInAnIntervalRange(0, 0)) // 0
}

// Time: O(1), Space: O(1)
func CountOddNumbersInAnIntervalRange(low int, high int) int {
	return (high + 1) / 2 - low / 2
}
```
