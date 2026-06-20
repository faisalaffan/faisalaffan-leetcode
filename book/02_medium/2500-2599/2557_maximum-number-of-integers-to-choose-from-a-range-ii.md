# 2557 — Maximum Number Of Integers To Choose From A Range Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maxCount(banned []int, n int, maxSum int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2557: Maximum Number of Integers to Choose From a Range II
// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maxCount(banned []int, n int, maxSum int) int {
  // Sort O(n log n)
	sort.Ints(banned)
  // HashMap: O(1) lookup
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		if sum+i > maxSum {
			break
		}
		sum += i
		count++
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxCount([]int{1, 6, 5}, 5, 6))
	// Expected: 2

	// Test case 2: large banned set
	fmt.Println("Test 2:", maxCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, 40))
	// Expected: 3 (pick 11,12,13)

	// Test case 3
	fmt.Println("Test 3:", maxCount([]int{}, 5, 15))
	// Expected: 5 (1+2+3+4+5=15)
}
```
