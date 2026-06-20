# 2779 — Maximum Beauty Of An Array After Applying Operation

## Deskripsi

**Soal:** [2779. Maximum Beauty Of An Array After Applying Operation](https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MaximumBeautyOfAnArrayAfterApplyingOperation(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2779: Maximum Beauty of an Array After Applying Operation
// https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func MaximumBeautyOfAnArrayAfterApplyingOperation(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	best := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

func main() {
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{4, 6, 1, 2}, 2))
	fmt.Println(MaximumBeautyOfAnArrayAfterApplyingOperation([]int{1, 1, 1, 1}, 10))
}
```
