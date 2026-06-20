# 2909 — Minimum Sum Of Mountain Triplets Ii

## Deskripsi

**Soal:** [2909. Minimum Sum Of Mountain Triplets Ii](https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2909: Minimum Sum of Mountain Triplets II
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumSum([]int{8, 6, 1, 5, 3}))
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
	fmt.Println(minimumSum([]int{6, 5, 4, 3, 4, 5}))
}

func minimumSum(nums []int) int {
	n := len(nums)
	const inf = 1 << 30
  // Membuat slice untuk menyimpan hasil
	right := make([]int, n+1)
	right[n] = inf
	for i := n - 1; i >= 0; i-- {
		if right[i+1] < nums[i] {
			right[i] = right[i+1]
		} else {
			right[i] = nums[i]
		}
	}
	ans, left := inf, inf
	for i, x := range nums {
		if left < x && right[i+1] < x {
			sum := left + x + right[i+1]
			if sum < ans {
				ans = sum
			}
		}
		if x < left {
			left = x
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}
```
