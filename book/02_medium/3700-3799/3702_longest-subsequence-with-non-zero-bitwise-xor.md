# 3702 — Longest Subsequence With Non Zero Bitwise Xor

## Deskripsi

**Soal:** [3702. Longest Subsequence With Non Zero Bitwise Xor](https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func longestSubsequenceWithNonZeroBitwiseXor(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #3702: Longest Subsequence With Non-Zero Bitwise XOR
// https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequenceWithNonZeroBitwiseXor(nums []int) int {
	xorSum := 0
	allZero := true
	for _, num := range nums {
		xorSum ^= num
		if num != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}
	if xorSum != 0 {
		return len(nums)
	}
	return len(nums) - 1
}

func main() {
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{1, 2, 3}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{2, 3, 4}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{0, 0, 0}))
}
```
