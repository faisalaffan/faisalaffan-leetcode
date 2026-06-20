# 0693 — Binary Number With Alternating Bits

## Deskripsi

**Soal:** [0693. Binary Number With Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #693: Binary Number with Alternating Bits
// https://leetcode.com/problems/binary-number-with-alternating-bits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))  // true (101)
	fmt.Println(hasAlternatingBits(7))  // false (111)
	fmt.Println(hasAlternatingBits(11)) // false (1011)
	fmt.Println(hasAlternatingBits(10)) // true (1010)
}

// hasAlternatingBits checks if the binary representation of n has alternating bits.
// Time: O(log n). Space: O(1).
func hasAlternatingBits(n int) bool {
	// XOR with n>>1 gives all 1s if alternating
	x := n ^ (n >> 1)
	// Check if x is all 1s (i.e., x & (x+1) == 0)
	return x&(x+1) == 0
}
```
