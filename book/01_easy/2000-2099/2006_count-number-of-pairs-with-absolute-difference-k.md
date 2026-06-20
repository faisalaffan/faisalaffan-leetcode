# 2006 — Count Number Of Pairs With Absolute Difference K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2006: Count Number of Pairs With Absolute Difference K
// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 2, 2, 1}, 1))   // 4
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{1, 3}, 3))          // 0
	fmt.Println(CountNumberOfPairsWithAbsoluteDifferenceK([]int{3, 2, 1, 5, 4}, 2)) // 3
}

// Time: O(n), Space: O(n)
func CountNumberOfPairsWithAbsoluteDifferenceK(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	count := 0
	for _, v := range nums {
		count += freq[v-k] + freq[v+k]
		freq[v]++
	}
	return count
}
```
