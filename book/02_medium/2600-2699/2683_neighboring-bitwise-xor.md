# 2683 — Neighboring Bitwise Xor

## Deskripsi

**Soal:** [2683. Neighboring Bitwise Xor](https://leetcode.com/problems/neighboring-bitwise-xor/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func doesValidArrayExist(derived []int) bool`

## Solusi Go

```go
package main

// LeetCode #2683: Neighboring Bitwise XOR
// https://leetcode.com/problems/neighboring-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, v := range derived {
		xor ^= v
	}
	return xor == 0
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", doesValidArrayExist([]int{1, 1, 0}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", doesValidArrayExist([]int{1, 1}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", doesValidArrayExist([]int{1, 0}))
	// Expected: false
}
```
