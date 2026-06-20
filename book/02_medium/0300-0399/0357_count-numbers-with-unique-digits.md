# 0357 — Count Numbers With Unique Digits

## Deskripsi

**Soal:** [0357. Count Numbers With Unique Digits](https://leetcode.com/problems/count-numbers-with-unique-digits/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countNumbersWithUniqueDigits(n int) int`

## Solusi Go

```go
package main

// LeetCode #357: Count Numbers with Unique Digits
// https://leetcode.com/problems/count-numbers-with-unique-digits/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
  // Edge case: input kosong
	if n == 0 {
		return 1
	}

	count := 10 // n=1: 0-9
	product := 9
	available := 9

	for i := 2; i <= n && i <= 10; i++ {
		product *= available
		available--
		count += product
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countNumbersWithUniqueDigits(2))
	// Expected: 91

	// Test case 2
	fmt.Println("Test 2:", countNumbersWithUniqueDigits(0))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", countNumbersWithUniqueDigits(3))
	// Expected: 739
}
```
