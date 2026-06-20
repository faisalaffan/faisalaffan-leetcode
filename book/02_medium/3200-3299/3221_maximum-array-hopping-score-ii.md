# 3221 — Maximum Array Hopping Score Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxScore(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3221: Maximum Array Hopping Score II
// https://leetcode.com/problems/maximum-array-hopping-score-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxScore(nums []int) int64 {
	ans := int64(0)
	mx := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
		}
		ans += int64(mx)
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5})) // Expected: 14
	fmt.Println(maxScore([]int{5, 4, 3, 2, 1})) // Expected: 4
}
```
