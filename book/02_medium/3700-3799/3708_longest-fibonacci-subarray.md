# 3708 — Longest Fibonacci Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestFibonacciSubarray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3708: Longest Fibonacci Subarray
// https://leetcode.com/problems/longest-fibonacci-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestFibonacciSubarray(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	ans := 2
	cnt := 2
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-2] {
			cnt++
		} else {
			cnt = 2
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(longestFibonacciSubarray([]int{1, 1, 2, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 2, 3, 4, 5, 6}))
}
```
