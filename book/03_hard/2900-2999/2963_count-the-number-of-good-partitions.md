# 2963 — Count The Number Of Good Partitions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfGoodPartitions(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2963: Count the Number of Good Partitions
// https://leetcode.com/problems/count-the-number-of-good-partitions/
// Difficulty: Hard
//
// For each value, find its first and last occurrence (forming an interval).
// Merge overlapping intervals. Each merged segment can be an independent
// partition boundary. Number of ways to partition k independent segments
// into contiguous groups = 2^(k-1) mod (10^9+7).
//
// A partition is "good" if no value appears in more than one part.

import "fmt"

func numberOfGoodPartitions(nums []int) int {
	const mod = 1_000_000_007

	// Last occurrence of each value
  // HashMap: O(1) lookup
	last := make(map[int]int)
	for i, x := range nums {
		last[x] = i
	}

	// Scan and merge overlapping intervals
	maxEnd := -1
	parts := 0
	for i, x := range nums {
		if last[x] > maxEnd {
			maxEnd = last[x]
		}
		if i == maxEnd {
			parts++
		}
	}

	// 2^(parts-1) mod MOD
	ans := 1
	for i := 1; i < parts; i++ {
		ans = (ans * 2) % mod
	}
	return ans
}

func main() {
	// Example: [1,2,3,4] -> 8 (4 segments, 2^3)
	fmt.Println(numberOfGoodPartitions([]int{1, 2, 3, 4}))

	// All same value
	fmt.Println(numberOfGoodPartitions([]int{1, 1, 1, 1}))

	// Overlapping intervals
	fmt.Println(numberOfGoodPartitions([]int{1, 2, 1, 3}))

	// Two distinct values interleaved
	fmt.Println(numberOfGoodPartitions([]int{1, 2, 1, 2}))
}
```
