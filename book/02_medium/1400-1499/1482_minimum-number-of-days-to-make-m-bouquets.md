# 1482 — Minimum Number Of Days To Make M Bouquets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minDays(bloomDay []int, m int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search

**Waktu:** O(n * log max(bloomDay)) for binary search  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1482: Minimum Number of Days to Make m Bouquets
// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
// Difficulty: Medium

import "fmt"
import "math"

func main() {
	// Test case 1
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1)) // 3

	// Test case 2
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 2)) // -1

	// Test case 3
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3)) // 12

	// Test case 4
	fmt.Println(minDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2)) // 9
}

// Time: O(n * log max(bloomDay)) for binary search
// Space: O(1)
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	minDay, maxDay := math.MaxInt32, 0
	for _, d := range bloomDay {
		if d < minDay {
			minDay = d
		}
		if d > maxDay {
			maxDay = d
		}
	}

	canMake := func(day int) bool {
		bouquets := 0
		consecutive := 0
		for _, d := range bloomDay {
			if d <= day {
				consecutive++
				if consecutive == k {
					bouquets++
					consecutive = 0
				}
			} else {
				consecutive = 0
			}
		}
		return bouquets >= m
	}

	// Binary search
	left, right := minDay, maxDay
  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2
		if canMake(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
```
