# 3215 — Count Triplets With Even Xor Set Bits Ii

## Deskripsi

**Soal:** [3215. Count Triplets With Even Xor Set Bits Ii](https://leetcode.com/problems/count-triplets-with-even-xor-set-bits-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func tripletCount(a []int, b []int, c []int) int64`

## Solusi Go

```go
package main

// LeetCode #3215: Count Triplets with Even XOR Set Bits II
// https://leetcode.com/problems/count-triplets-with-even-xor-set-bits-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func tripletCount(a []int, b []int, c []int) int64 {
	evenA, oddA := countBits(a)
	evenB, oddB := countBits(b)
	evenC, oddC := countBits(c)
	return int64(evenA)*int64(oddB)*int64(oddC) +
		int64(oddA)*int64(evenB)*int64(oddC) +
		int64(oddA)*int64(oddB)*int64(evenC) +
		int64(evenA)*int64(evenB)*int64(evenC)
}

func countBits(nums []int) (int, int) {
	even := 0
	for _, v := range nums {
		if bits.OnesCount(uint(v))%2 == 0 {
			even++
		}
	}
	return even, len(nums) - even
}

func main() {
	fmt.Println(tripletCount([]int{1, 2}, []int{3, 4}, []int{5, 6})) // Expected: depends on bit counts
	fmt.Println(tripletCount([]int{0, 0}, []int{0, 0}, []int{0, 0})) // Expected: 8
}
```
