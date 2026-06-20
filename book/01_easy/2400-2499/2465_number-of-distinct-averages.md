# 2465 — Number Of Distinct Averages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfDistinctAverages(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2465: Number of Distinct Averages
// https://leetcode.com/problems/number-of-distinct-averages/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(NumberOfDistinctAverages([]int{4, 1, 4, 0, 3, 5})) // 2
	fmt.Println(NumberOfDistinctAverages([]int{1, 100}))            // 1
}

func NumberOfDistinctAverages(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	seen := map[int]bool{}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		seen[nums[i]+nums[j]] = true
	}
	return len(seen)
}
```
