# 1800 — Maximum Ascending Subarray Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaxAscendingSum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1800: Maximum Ascending Subarray Sum
// https://leetcode.com/problems/maximum-ascending-subarray-sum/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaxAscendingSum(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	fmt.Println(MaxAscendingSum([]int{10, 20, 30, 40, 50}))
	fmt.Println(MaxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
```
