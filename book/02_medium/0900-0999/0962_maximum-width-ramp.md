# 0962 — Maximum Width Ramp

## Deskripsi

**Soal:** [0962. Maximum Width Ramp](https://leetcode.com/problems/maximum-width-ramp/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func maxWidthRamp(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #962: Maximum Width Ramp
// https://leetcode.com/problems/maximum-width-ramp/
// Difficulty: Medium

import "fmt"

// Time: O(n) | Space: O(n)
func maxWidthRamp(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)

	// Build decreasing stack of indices
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	ans := 0
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			width := j - stack[len(stack)-1]
			if width > ans {
				ans = width
			}
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5}))
	fmt.Println(maxWidthRamp([]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}))
}
```
