# 1911 — Maximum Alternating Subsequence Sum

## Deskripsi

**Soal:** [1911. Maximum Alternating Subsequence Sum](https://leetcode.com/problems/maximum-alternating-subsequence-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1911: Maximum Alternating Subsequence Sum
// https://leetcode.com/problems/maximum-alternating-subsequence-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxAlternatingSum([]int{4, 2, 5, 3}))
	fmt.Println(MaxAlternatingSum([]int{5, 6, 7, 8}))
	fmt.Println(MaxAlternatingSum([]int{6, 2, 1, 2, 4, 5}))
}

// Time: O(n), Space: O(1)
func MaxAlternatingSum(nums []int) int64 {
	even := int64(nums[0]) // max alternating sum ending with even index (added)
	odd := int64(0)         // max alternating sum ending with odd index (subtracted)

	for i := 1; i < len(nums); i++ {
		newEven := max64(even, max64(odd+int64(nums[i]), int64(nums[i])))
		odd = max64(odd, even-int64(nums[i]))
		even = newEven
	}
	return even
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```
