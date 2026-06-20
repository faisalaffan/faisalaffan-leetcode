# 1968 — Array With Elements Not Equal To Average Of Neighbors

## Deskripsi

**Soal:** [1968. Array With Elements Not Equal To Average Of Neighbors](https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1968: Array With Elements Not Equal to Average of Neighbors
// https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RearrangeArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RearrangeArray([]int{6, 2, 0, 9, 7}))
}

// Time: O(n log n), Space: O(n)
func RearrangeArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = nums[left]
			left++
		} else {
			result[i] = nums[right]
			right--
		}
	}
	return result
}
```
