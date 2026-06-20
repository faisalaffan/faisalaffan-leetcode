# 2044 — Count Number Of Maximum Bitwise Or Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countMaxOrSubsets(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking

**Waktu:** O(2^n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2044: Count Number of Maximum Bitwise-OR Subsets
// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import "fmt"

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	for _, v := range nums {
		maxOr |= v
	}

	count := 0
	var backtrack func(idx int, curOr int)
	backtrack = func(idx int, curOr int) {
		if idx == len(nums) {
			if curOr == maxOr {
				count++
			}
			return
		}
		// Skip current
		backtrack(idx+1, curOr)
		// Take current
		backtrack(idx+1, curOr|nums[idx])
	}

	backtrack(0, 0)
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countMaxOrSubsets([]int{3, 1}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", countMaxOrSubsets([]int{2, 2, 2}))
	// Expected: 7

	// Test case 3
	fmt.Println("Test 3:", countMaxOrSubsets([]int{3, 2, 1, 5}))
	// Expected: 6
}
```
