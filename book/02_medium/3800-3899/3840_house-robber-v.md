# 3840 — House Robber V

## Deskripsi

**Soal:** [3840. House Robber V](https://leetcode.com/problems/house-robber-v/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func HouseRobberV(nums []int, colors []int) int`

> **Ide Kunci:** DP with two states (rob/notRob). If adjacent houses have same color,

## Solusi Go

```go
package main

// LeetCode #3840: House Robber V
// https://leetcode.com/problems/house-robber-v/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: DP with two states (rob/notRob). If adjacent houses have same color,
// cannot rob both. If different colors, can rob both.

import "fmt"

func HouseRobberV(nums []int, colors []int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	notRob, rob := 0, nums[0]

	for i := 1; i < n; i++ {
		newNotRob := max(notRob, rob)
		var newRob int
		if colors[i] != colors[i-1] {
			newRob = max(notRob, rob) + nums[i]
		} else {
			newRob = notRob + nums[i]
		}
		notRob, rob = newNotRob, newRob
	}

	return max(notRob, rob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(HouseRobberV([]int{1, 4, 3, 5}, []int{1, 1, 2, 2})) // Expected: 9

	// Example 2
	fmt.Println(HouseRobberV([]int{3, 1, 2, 4}, []int{2, 3, 2, 2})) // Expected: 8

	// Example 3
	fmt.Println(HouseRobberV([]int{10, 1, 3, 9}, []int{1, 1, 1, 2})) // Expected: 22
}
```
