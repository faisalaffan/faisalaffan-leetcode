# 1552 — Magnetic Force Between Two Balls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaxDistance(position []int, m int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(N log N + N log MAX), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1552: Magnetic Force Between Two Balls
// https://leetcode.com/problems/magnetic-force-between-two-balls/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 7}, 3))
	fmt.Println(MaxDistance([]int{5, 4, 3, 2, 1, 1000000000}, 2))
	fmt.Println(MaxDistance([]int{1, 2, 3, 4, 5, 6}, 3))
}

func MaxDistance(position []int, m int) int {
	// Time: O(N log N + N log MAX), Space: O(1)
  // Sort O(n log n)
	sort.Ints(position)

	low, high := 1, position[len(position)-1]-position[0]
	result := 0

	for low <= high {
		mid := low + (high-low)/2
		if canPlace(position, m, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}

func canPlace(position []int, m int, minDist int) bool {
	count := 1
	lastPos := position[0]

	for i := 1; i < len(position); i++ {
		if position[i]-lastPos >= minDist {
			count++
			lastPos = position[i]
			if count >= m {
				return true
			}
		}
	}

	return false
}
```
