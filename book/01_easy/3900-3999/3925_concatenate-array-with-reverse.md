# 3925 — Concatenate Array With Reverse

## Deskripsi

**Soal:** [3925. Concatenate Array With Reverse](https://leetcode.com/problems/concatenate-array-with-reverse/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3925: Concatenate Array With Reverse
// https://leetcode.com/problems/concatenate-array-with-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenateArrayWithReverse([]int{1, 2, 3}))
	fmt.Println(ConcatenateArrayWithReverse([]int{1}))
}

// Time: O(n)
// Space: O(n)
func ConcatenateArrayWithReverse(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = nums[n-1-i]
	}
	return ans
}
```
