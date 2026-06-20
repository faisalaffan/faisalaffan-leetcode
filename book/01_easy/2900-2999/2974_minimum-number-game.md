# 2974 — Minimum Number Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MinimumNumberGame(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2974: Minimum Number Game
// https://leetcode.com/problems/minimum-number-game/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: numberGame
	fmt.Println(MinimumNumberGame([]int{5, 4, 2, 3})) // [3, 2, 5, 4]
	fmt.Println(MinimumNumberGame([]int{2, 5}))       // [5, 2]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: numberGame
func MinimumNumberGame(nums []int) []int {
  // Sort O(n log n)
	sort.Ints(nums)
  // Alokasi slice
	result := make([]int, len(nums))
  // Linear scan O(n)
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[i+1]
		result[i+1] = nums[i]
	}
	return result
}
```
