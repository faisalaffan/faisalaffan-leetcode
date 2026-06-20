# 0045 — Jump Game Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func jump(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #45: Jump Game II
// https://leetcode.com/problems/jump-game-ii/
// Difficulty: Medium

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

func main() {
	// Test case 1
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2

	// Test case 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 2

	// Test case 3
	fmt.Println(jump([]int{0})) // 0
}

// Time: O(n) | Space: O(1)
```
