# 1004 — Max Consecutive Ones Iii

## Deskripsi

**Soal:** [1004. Max Consecutive Ones Iii](https://leetcode.com/problems/max-consecutive-ones-iii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Sliding Window (jendela geser)

> **Ide Kunci:** Sliding window (expand right, shrink left when zeros > k)

## Solusi Go

```go
package main

// LeetCode #1004: Max Consecutive Ones III
// https://leetcode.com/problems/max-consecutive-ones-iii/
// Difficulty: Medium
//
// Approach: Sliding window (expand right, shrink left when zeros > k)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
	fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0)) // 0
}

func longestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
```
