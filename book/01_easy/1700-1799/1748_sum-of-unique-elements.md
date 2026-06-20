# 1748 — Sum Of Unique Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SumOfUnique(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1748: Sum of Unique Elements
// https://leetcode.com/problems/sum-of-unique-elements/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func SumOfUnique(nums []int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sum := 0
	for num, count := range freq {
		if count == 1 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(SumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(SumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(SumOfUnique([]int{1, 2, 3, 4, 5}))
}
```
