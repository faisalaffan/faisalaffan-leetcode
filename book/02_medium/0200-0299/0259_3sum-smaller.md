# 0259 — 3Sum Smaller

## Deskripsi

**Soal:** [0259. 3Sum Smaller](https://leetcode.com/problems/3sum-smaller/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func threeSumSmaller(nums []int, target int) int`

## Solusi Go

```go
package main

// LeetCode #259: 3Sum Smaller
// https://leetcode.com/problems/3sum-smaller/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(1)

import (
	"fmt"
	"sort"
)

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	count := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
  // Loop two-pointer: kiri vs kanan
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{1, 1, -2}, 1))
	fmt.Println(threeSumSmaller([]int{0, 0, 0}, 0))
}
```
