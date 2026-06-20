# 0365 — Water And Jug Problem

## Deskripsi

**Soal:** [0365. Water And Jug Problem](https://leetcode.com/problems/water-and-jug-problem/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log min(x,y))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func canMeasureWater(x int, y int, target int) bool`

## Solusi Go

```go
package main

// LeetCode #365: Water and Jug Problem
// https://leetcode.com/problems/water-and-jug-problem/
// Difficulty: Medium
// Time: O(log min(x,y)) | Space: O(1)

import "fmt"

func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	return target%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", canMeasureWater(3, 5, 4))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", canMeasureWater(2, 6, 5))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", canMeasureWater(1, 2, 3))
	// Expected: true
}
```
