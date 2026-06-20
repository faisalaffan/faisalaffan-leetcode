# 2505 — Bitwise Or Of All Subsequence Sums

## Deskripsi

**Soal:** [2505. Bitwise Or Of All Subsequence Sums](https://leetcode.com/problems/bitwise-or-of-all-subsequence-sums/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 20)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Prefix Sum (jumlah kumulatif)

## Solusi Go

```go
package main

// LeetCode #2505: Bitwise OR of All Subsequence Sums
// https://leetcode.com/problems/bitwise-or-of-all-subsequence-sums/
// Difficulty: Medium
// Time: O(n * 20) | Space: O(1)
// A bit is achievable if any number has that bit, or can be formed by combination.
// Result = OR of all prefix sums? Actually: any sum that can be formed = OR of
// all elements and their combinations. Answer = OR of all elements (since we can
// always form any single element sum via a subsequence of size 1).

import "fmt"

func main() {
	fmt.Println(subsequenceSumOr([]int{2, 1, 4})) // 7 (bits 0,1,2)
	fmt.Println(subsequenceSumOr([]int{2, 3}))    // 7 (sums: 0,2,3,5 -> OR = 7)
}

func subsequenceSumOr(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans |= v
	}
	return ans
}
```
