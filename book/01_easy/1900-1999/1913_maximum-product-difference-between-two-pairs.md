# 1913 — Maximum Product Difference Between Two Pairs

## Deskripsi

**Soal:** [1913. Maximum Product Difference Between Two Pairs](https://leetcode.com/problems/maximum-product-difference-between-two-pairs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1913: Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{5, 6, 2, 7, 4}))      // 34
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{4, 2, 5, 9, 7, 4, 8})) // 64
}

// Time: O(n log n), Space: O(1) ignoring sort
func MaximumProductDifferenceBetweenTwoPairs(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
```
