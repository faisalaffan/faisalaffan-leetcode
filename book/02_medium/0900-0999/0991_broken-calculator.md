# 0991 — Broken Calculator

## Deskripsi

**Soal:** [0991. Broken Calculator](https://leetcode.com/problems/broken-calculator/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log target)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

> **Ide Kunci:** Work backwards from target to startValue

## Solusi Go

```go
package main

// LeetCode #991: Broken Calculator
// https://leetcode.com/problems/broken-calculator/
// Difficulty: Medium
//
// Approach: Work backwards from target to startValue
//   - If target is even, divide by 2 (reverse of multiply by 2)
//   - If target is odd, add 1 (reverse of subtract 1)
// Time: O(log target)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))  // 2
	fmt.Println(brokenCalc(5, 8))  // 2
	fmt.Println(brokenCalc(3, 10)) // 3
}

func brokenCalc(startValue int, target int) int {
	ops := 0
	for target > startValue {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		ops++
	}
	return ops + (startValue - target)
}
```
