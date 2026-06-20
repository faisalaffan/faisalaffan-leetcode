# 2735 — Collecting Chocolates

## Deskripsi

**Soal:** [2735. Collecting Chocolates](https://leetcode.com/problems/collecting-chocolates/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func CollectingChocolates(nums []int, x int) int64`

## Solusi Go

```go
package main

// LeetCode #2735: Collecting Chocolates
// https://leetcode.com/problems/collecting-chocolates/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func CollectingChocolates(nums []int, x int) int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	minCost := make([]int, n)
	copy(minCost, nums)

	var best int64
	for i := 0; i < n; i++ {
		best += int64(nums[i])
	}

	for shift := 1; shift < n; shift++ {
		var total int64 = int64(shift) * int64(x)
		for i := 0; i < n; i++ {
			idx := (i + shift) % n
			if nums[idx] < minCost[i] {
				minCost[i] = nums[idx]
			}
			total += int64(minCost[i])
		}
		if total < best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(CollectingChocolates([]int{20, 1, 15}, 5))
	fmt.Println(CollectingChocolates([]int{1, 2, 3}, 4))
}
```
