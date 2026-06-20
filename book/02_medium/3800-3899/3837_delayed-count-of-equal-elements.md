# 3837 — Delayed Count Of Equal Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func DelayedCountOfEqualElements(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3837: Delayed Count of Equal Elements
// https://leetcode.com/problems/delayed-count-of-equal-elements/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Track frequency of each value and count pairs where values are equal.

import "fmt"

func DelayedCountOfEqualElements(nums []int) int64 {
  // HashMap: O(1) lookup
	freq := make(map[int]int64)
	var ans int64

	for _, v := range nums {
		ans += freq[v]
		freq[v]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 1, 2, 1})) // Expected: 4

	// Example 2
	fmt.Println(DelayedCountOfEqualElements([]int{1, 1, 1, 1})) // Expected: 6

	// Example 3
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 3})) // Expected: 0
}
```
