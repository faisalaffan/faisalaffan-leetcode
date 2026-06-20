# 1480 — Running Sum Of 1D Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func runningSum(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1) excluding output  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #1480: Running Sum of 1d Array
// https://leetcode.com/problems/running-sum-of-1d-array/
// Difficulty: Easy
//
// LeetCode submission: func runningSum(nums []int) []int

import "fmt"

func main() {
	fmt.Println(RunningSumOfOneDArray([]int{1, 2, 3, 4}))    // [1 3 6 10]
	fmt.Println(RunningSumOfOneDArray([]int{1, 1, 1, 1, 1})) // [1 2 3 4 5]
}

// Time: O(n), Space: O(1) excluding output
func RunningSumOfOneDArray(nums []int) []int {
  // Alokasi slice
	res := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		res[i] = sum
	}
	return res
}
```
