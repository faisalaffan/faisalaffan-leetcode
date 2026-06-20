# 0487 — Max Consecutive Ones Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaxConsecutiveOnesIi(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #487: Max Consecutive Ones II
// https://leetcode.com/problems/max-consecutive-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0}))
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0, 1}))
}

func MaxConsecutiveOnesIi(nums []int) int {
	maxLen := 0
	prevLen, curLen := 0, 0

	for _, num := range nums {
		if num == 1 {
			curLen++
		} else {
			prevLen = curLen
			curLen = 0
		}
		if prevLen+curLen+1 > maxLen {
			maxLen = prevLen + curLen + 1
		}
	}

	if maxLen > len(nums) {
		return len(nums)
	}
	return maxLen
}
```
