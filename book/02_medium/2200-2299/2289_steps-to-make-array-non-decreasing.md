# 2289 — Steps To Make Array Non Decreasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func totalSteps(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2289: Steps to Make Array Non-decreasing
// https://leetcode.com/problems/steps-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func totalSteps(nums []int) int {
	stack := []int{} // indices
  // Alokasi slice
	steps := make([]int, len(nums))
	maxSteps := 0

  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		curSteps := 0
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			if steps[stack[len(stack)-1]] > curSteps {
				curSteps = steps[stack[len(stack)-1]]
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			steps[i] = curSteps + 1
		} else {
			steps[i] = 0
		}
		if steps[i] > maxSteps {
			maxSteps = steps[i]
		}
		stack = append(stack, i)
	}
	return maxSteps
}

func main() {
	// Test case 1
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	// Expected: 3

	// Test case 2
	fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	// Expected: 0

	// Test case 3
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
	// Expected: 6
}
```
