# 1985 — Find The Kth Largest Integer In The Array

## Deskripsi

**Soal:** [1985. Find The Kth Largest Integer In The Array](https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(1) (ignoring sort space)  
**Kompleksitas Ruang:** O(1) (ignoring sort space)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1985: Find the Kth Largest Integer in the Array
// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"3", "6", "7", "10"}, 4))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"2", "21", "12", "1"}, 3))
	fmt.Println(FindTheKthLargestIntegerInTheArray([]string{"0", "0"}, 2))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func FindTheKthLargestIntegerInTheArray(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) > len(nums[j])
		}
		return nums[i] > nums[j]
	})
	return nums[k-1]
}
```
