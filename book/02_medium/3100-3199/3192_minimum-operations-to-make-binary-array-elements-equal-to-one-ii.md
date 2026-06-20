# 3192 — Minimum Operations To Make Binary Array Elements Equal To One Ii

## Deskripsi

**Soal:** [3192. Minimum Operations To Make Binary Array Elements Equal To One Ii](https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minOperations(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #3192: Minimum Operations to Make Binary Array Elements Equal to One II
// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	ans := 0
	flip := 0

	for i := 0; i < n; i++ {
		cur := nums[i] ^ flip
		if cur == 0 {
			ans++
			flip ^= 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 0, 1})) // Expected: 3
	fmt.Println(minOperations([]int{1, 0, 1, 0}))     // Expected: 2
}
```
