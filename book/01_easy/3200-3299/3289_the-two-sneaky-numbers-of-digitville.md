# 3289 — The Two Sneaky Numbers Of Digitville

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TheTwoSneakyNumbersOfDigitville(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3289: The Two Sneaky Numbers of Digitville
// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 1, 1, 0}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{0, 3, 2, 1, 3, 2}))
	fmt.Println(TheTwoSneakyNumbersOfDigitville([]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}))
}

// TheTwoSneakyNumbersOfDigitville returns the two numbers that appear twice in the array.
// Time: O(n). Space: O(n).
func TheTwoSneakyNumbersOfDigitville(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]int)
	result := []int{}
	for _, num := range nums {
		seen[num]++
		if seen[num] == 2 {
			result = append(result, num)
		}
	}
	return result
}
```
