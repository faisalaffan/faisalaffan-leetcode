# 2429 — Minimize Xor

## Deskripsi

**Soal:** [2429. Minimize Xor](https://leetcode.com/problems/minimize-xor/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(30)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

> **Ide Kunci:** set highest bits of num1 first, then lowest unset bits.

## Solusi Go

```go
package main

// LeetCode #2429: Minimize XOR
// https://leetcode.com/problems/minimize-xor/
// Difficulty: Medium
// Time: O(30) | Space: O(1)
// Find x with same set bits as num2 that minimizes x XOR num1.
// Approach: set highest bits of num1 first, then lowest unset bits.

import "fmt"

func main() {
	fmt.Println(minimizeXor(3, 5))  // 3  (3=11, 5=101, set bits: 2)
	fmt.Println(minimizeXor(1, 12)) // 3  (1=1, 12=1100, set bits: 2)
}

func minimizeXor(num1 int, num2 int) int {
	targetBits := bitsCount(num2)
	x := 0
	// Set bits from num1's highest bits first
	for b := 30; b >= 0 && targetBits > 0; b-- {
		if num1>>b&1 == 1 {
			x |= (1 << b)
			targetBits--
		}
	}
	// Set remaining bits from lowest unset positions
	for b := 0; b <= 30 && targetBits > 0; b++ {
		if x>>b&1 == 0 {
			x |= (1 << b)
			targetBits--
		}
	}
	return x
}

func bitsCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
```
