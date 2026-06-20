# 2465 — Number Of Distinct Averages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfDistinctAverages(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2465: Number of Distinct Averages
// https://leetcode.com/problems/number-of-distinct-averages/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(NumberOfDistinctAverages([]int{4, 1, 4, 0, 3, 5})) // 2
	fmt.Println(NumberOfDistinctAverages([]int{1, 100}))            // 1
}

func NumberOfDistinctAverages(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	seen := map[int]bool{}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		seen[nums[i]+nums[j]] = true
	}
	return len(seen)
}
```
