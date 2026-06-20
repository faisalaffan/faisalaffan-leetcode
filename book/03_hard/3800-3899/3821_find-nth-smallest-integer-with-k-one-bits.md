# 3821 — Find Nth Smallest Integer With K One Bits

## Deskripsi

**Soal:** [3821. Find Nth Smallest Integer With K One Bits](https://leetcode.com/problems/find-nth-smallest-integer-with-k-one-bits/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Use combinatorial ranking. Generate numbers with k bits

## Solusi Go

```go
package main

// LeetCode #3821: Find Nth Smallest Integer With K One Bits
// https://leetcode.com/problems/find-nth-smallest-integer-with-k-one-bits/
// Difficulty: Hard
//
// Find the n-th smallest positive integer that has exactly k set
// bits in its binary representation.
//
// Approach: Use combinatorial ranking. Generate numbers with k bits
// in increasing order using next combination (Gosper's hack).

import "fmt"

func main() {
	// Example 1
	fmt.Println(nthSmallest(3, 2))
	// Example 2
	fmt.Println(nthSmallest(5, 3))
	// Edge: n=1
	fmt.Println(nthSmallest(1, 1))
	// Edge: k=0
	fmt.Println(nthSmallest(1, 0))
}

func nthSmallest(n int64, k int) int64 {
	if k == 0 {
		if n == 1 {
			return 0
		}
		return -1
	}

	num := int64((1 << uint(k)) - 1)
	for i := int64(1); i < n; i++ {
		// Next number with k bits set (Gosper's hack)
		c := num & -num
		r := num + c
		num = (((r ^ num) >> 2) / c) | r
	}
	return num
}
```
