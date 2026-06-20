# 1099 — Two Sum Less Than K

## Deskripsi

**Soal:** [1099. Two Sum Less Than K](https://leetcode.com/problems/two-sum-less-than-k/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1099: Two Sum Less Than K
// https://leetcode.com/problems/two-sum-less-than-k/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)) // 58
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15))                   // -1
}

// LeetCode submission: twoSumLessThanK
func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	ans := -1
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum < k {
			if sum > ans {
				ans = sum
			}
			i++
		} else {
			j--
		}
	}
	return ans
}
```
