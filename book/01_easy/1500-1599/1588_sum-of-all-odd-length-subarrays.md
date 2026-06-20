# 1588 — Sum Of All Odd Length Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sumOddLengthSubarrays(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1588: Sum of All Odd Length Subarrays
// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/
// Difficulty: Easy
//
// LeetCode submission: func sumOddLengthSubarrays(arr []int) int

import "fmt"

func main() {
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 4, 2, 5, 3})) // 58
	fmt.Println(SumOfAllOddLengthSubarrays([]int{1, 2}))          // 3
	fmt.Println(SumOfAllOddLengthSubarrays([]int{10, 11, 12}))    // 66
}

// Time: O(n), Space: O(1)
func SumOfAllOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0
	for i, v := range arr {
		contribution := ((i+1)*(n-i) + 1) / 2
		sum += v * contribution
	}
	return sum
}
```
