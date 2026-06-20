# 3876 — Construct Uniform Parity Array Ii

## Deskripsi

**Soal:** [3876. Construct Uniform Parity Array Ii](https://leetcode.com/problems/construct-uniform-parity-array-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ConstructUniformParityArrayIi(nums1 []int) bool`

> **Ide Kunci:** If all same parity, return true. Else find smallest odd number.

## Solusi Go

```go
package main

// LeetCode #3876: Construct Uniform Parity Array II
// https://leetcode.com/problems/construct-uniform-parity-array-ii/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: If all same parity, return true. Else find smallest odd number.
// If any even number smaller than smallest odd, impossible (can't flip parity).

import "fmt"

func ConstructUniformParityArrayIi(nums1 []int) bool {
	hasEven, hasOdd := false, false
	minOdd := int(1e9 + 1)
	for _, v := range nums1 {
		if v%2 == 0 {
			hasEven = true
		} else {
			hasOdd = true
			if v < minOdd {
				minOdd = v
			}
		}
	}
	if !hasEven || !hasOdd {
		return true
	}
	// Both parities present. Check if any even < smallest odd.
	for _, v := range nums1 {
		if v%2 == 0 && v < minOdd {
			return false
		}
	}
	return true
}

func main() {
	// Example 1
	fmt.Println(ConstructUniformParityArrayIi([]int{1, 4, 7})) // Expected: true

	// Example 2
	fmt.Println(ConstructUniformParityArrayIi([]int{2, 3})) // Expected: false

	// Example 3
	fmt.Println(ConstructUniformParityArrayIi([]int{4, 6})) // Expected: true
}
```
