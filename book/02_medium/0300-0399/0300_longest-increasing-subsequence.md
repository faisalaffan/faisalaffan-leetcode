# 0300 — Longest Increasing Subsequence

## Deskripsi

**Soal:** [0300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func lengthOfLIS(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #300: Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import "fmt"

func lengthOfLIS(nums []int) int {
	tails := []int{}

	for _, num := range nums {
		left, right := 0, len(tails)
  // Loop two-pointer: kiri vs kanan
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
```
