# 0487 — Max Consecutive Ones Ii

## Deskripsi

**Soal:** [0487. Max Consecutive Ones Ii](https://leetcode.com/problems/max-consecutive-ones-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #487: Max Consecutive Ones II
// https://leetcode.com/problems/max-consecutive-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0}))
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0, 1}))
}

func MaxConsecutiveOnesIi(nums []int) int {
	maxLen := 0
	prevLen, curLen := 0, 0

	for _, num := range nums {
		if num == 1 {
			curLen++
		} else {
			prevLen = curLen
			curLen = 0
		}
		if prevLen+curLen+1 > maxLen {
			maxLen = prevLen + curLen + 1
		}
	}

	if maxLen > len(nums) {
		return len(nums)
	}
	return maxLen
}
```
