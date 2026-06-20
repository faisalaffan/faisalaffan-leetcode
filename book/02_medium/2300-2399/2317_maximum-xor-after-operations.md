# 2317 — Maximum Xor After Operations

## Deskripsi

**Soal:** [2317. Maximum Xor After Operations](https://leetcode.com/problems/maximum-xor-after-operations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximumXOR(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2317: Maximum XOR After Operations
// https://leetcode.com/problems/maximum-xor-after-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumXOR(nums []int) int {
	result := 0
	for _, v := range nums {
		result |= v
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
	// Expected: 7

	// Test case 2
	fmt.Println(maximumXOR([]int{1, 2, 3, 4, 5, 6, 7}))
	// Expected: 7

	// Test case 3
	fmt.Println(maximumXOR([]int{0}))
	// Expected: 0
}
```
