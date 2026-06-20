# 0371 — Sum Of Two Integers

## Deskripsi

**Soal:** [0371. Sum Of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func getSum(a int, b int) int`

## Solusi Go

```go
package main

// LeetCode #371: Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getSum(1, 2))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", getSum(2, 3))
	// Expected: 5

	// Test case 3: Negative numbers
	fmt.Println("Test 3:", getSum(-1, 1))
	// Expected: 0
}
```
