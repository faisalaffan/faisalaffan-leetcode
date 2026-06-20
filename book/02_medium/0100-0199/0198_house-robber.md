# 0198 — House Robber

## Deskripsi

**Soal:** [0198. House Robber](https://leetcode.com/problems/house-robber/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func rob(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #198: House Robber
// https://leetcode.com/problems/house-robber/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rob(nums []int) int {
	prev, curr := 0, 0

	for _, num := range nums {
		prev, curr = curr, max(curr, prev+num)
	}

	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{0}))
}
```
