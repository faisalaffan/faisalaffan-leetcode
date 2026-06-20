# 3719 — Longest Balanced Subarray I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestBalancedSubarrayI(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3719: Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func longestBalancedSubarrayI(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		if n-i <= ans {
			break
		}
  // HashMap: O(1) lookup
		evenVisited := make(map[int]bool)
  // HashMap: O(1) lookup
		oddVisited := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenVisited[nums[j]] {
					evenVisited[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddVisited[nums[j]] {
					oddVisited[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				if j-i+1 > ans {
					ans = j - i + 1
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestBalancedSubarrayI([]int{1, 2, 3, 4}))
	fmt.Println(longestBalancedSubarrayI([]int{2, 4, 6, 8}))
	fmt.Println(longestBalancedSubarrayI([]int{1, 1, 2, 2, 3, 3}))
}
```
