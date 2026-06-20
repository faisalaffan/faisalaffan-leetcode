# 2177 — Find Three Consecutive Integers That Sum To A Given Number

## Deskripsi

**Soal:** [2177. Find Three Consecutive Integers That Sum To A Given Number](https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func sumOfThree(num int64) []int64`

## Solusi Go

```go
package main

// LeetCode #2177: Find Three Consecutive Integers That Sum to a Given Number
// https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	x := num / 3
	return []int64{x - 1, x, x + 1}
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sumOfThree(33))
	// Expected: [10, 11, 12]

	// Test case 2
	fmt.Println("Test 2:", sumOfThree(4))
	// Expected: []

	// Test case 3
	fmt.Println("Test 3:", sumOfThree(0))
	// Expected: [-1, 0, 1]
}
```
