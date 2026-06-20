# 1695 — Maximum Erasure Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumUniqueSubarray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1695: Maximum Erasure Value
// https://leetcode.com/problems/maximum-erasure-value/
// Difficulty: Medium
// Time: O(n), Space: O(n)

import "fmt"

func maximumUniqueSubarray(nums []int) int {
  // HashMap: O(1) lookup
	lastPos := make(map[int]int)
	maxSum := 0
	currentSum := 0
	left := 0

	for right, num := range nums {
		if pos, ok := lastPos[num]; ok && pos >= left {
			// Remove elements from left to pos
			for left <= pos {
				currentSum -= nums[left]
				left++
			}
		}
		currentSum += num
		lastPos[num] = right
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))   // Expected: 17
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) // Expected: 8
	fmt.Println(maximumUniqueSubarray([]int{1})) // Expected: 1
}
```
