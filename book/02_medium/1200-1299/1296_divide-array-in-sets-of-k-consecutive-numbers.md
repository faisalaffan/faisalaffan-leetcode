# 1296 — Divide Array In Sets Of K Consecutive Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func isPossibleDivide(nums []int, k int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1296: Divide Array in Sets of K Consecutive Numbers
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
// Difficulty: Medium

// Sort array, greedily form groups of size k.
// Use frequency map to track available numbers.

// Time: O(n log n)
// Space: O(n)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

  // Sort O(n log n)
	sort.Ints(nums)
	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			if freq[v+i] == 0 {
				return false
			}
			freq[v+i]--
		}
	}

	return true
}

func main() {
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	fmt.Printf("%t (expected: false)\n", isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 4}, 2))
}
```
