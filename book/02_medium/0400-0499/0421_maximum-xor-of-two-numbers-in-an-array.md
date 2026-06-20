# 0421 — Maximum Xor Of Two Numbers In An Array

## Deskripsi

**Soal:** [0421. Maximum Xor Of Two Numbers In An Array](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func findMaximumXOR(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #421: Maximum XOR of Two Numbers in an Array
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findMaximumXOR(nums []int) int {
	maxXor := 0
	mask := 0

	// Try each bit from MSB to LSB
	for i := 31; i >= 0; i-- {
		mask |= 1 << i
  // Membuat map untuk pencarian O(1): key → value
		prefixSet := make(map[int]bool)
		for _, num := range nums {
			prefixSet[num&mask] = true
		}

		candidate := maxXor | (1 << i)
		for prefix := range prefixSet {
			if prefixSet[prefix^candidate] {
				maxXor = candidate
				break
			}
		}
	}
	return maxXor
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	// Expected: 28

	// Test case 2
	fmt.Println("Test 2:", findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
	// Expected: 127

	// Test case 3
	fmt.Println("Test 3:", findMaximumXOR([]int{0}))
	// Expected: 0
}
```
