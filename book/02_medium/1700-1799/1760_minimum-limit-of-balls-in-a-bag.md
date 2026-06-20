# 1760 — Minimum Limit Of Balls In A Bag

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumSize(nums []int, maxOperations int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log M) where M = max(nums), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1760: Minimum Limit of Balls in a Bag
// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
// Difficulty: Medium
// Time: O(n log M) where M = max(nums), Space: O(1)

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 0
	for _, v := range nums {
		if v > right {
			right = v
		}
	}

  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2
		if canDivide(nums, maxOperations, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canDivide(nums []int, maxOps, limit int) bool {
	ops := 0
	for _, v := range nums {
		if v > limit {
			ops += (v - 1) / limit
			if ops > maxOps {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(minimumSize([]int{9}, 2))               // Expected: 3
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))      // Expected: 2
	fmt.Println(minimumSize([]int{7, 17}, 2))            // Expected: 7
}
```
