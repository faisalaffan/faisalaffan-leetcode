# 1798 — Maximum Number Of Consecutive Values You Can Make

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func getMaximumConsecutive(coins []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1798: Maximum Number of Consecutive Values You Can Make
// https://leetcode.com/problems/maximum-number-of-consecutive-values-you-can-make/
// Difficulty: Medium
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

func getMaximumConsecutive(coins []int) int {
  // Sort O(n log n)
	sort.Ints(coins)
	maxReach := 0
	for _, c := range coins {
		if c > maxReach+1 {
			break
		}
		maxReach += c
	}
	return maxReach + 1
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1, 3}))          // Expected: 2
	fmt.Println(getMaximumConsecutive([]int{1, 1, 1, 4}))   // Expected: 8
	fmt.Println(getMaximumConsecutive([]int{1, 4, 10, 3, 1})) // Expected: 20
}
```
