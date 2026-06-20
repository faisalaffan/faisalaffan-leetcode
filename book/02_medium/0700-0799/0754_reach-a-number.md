# 0754 — Reach A Number

## Deskripsi

**Soal:** [0754. Reach A Number](https://leetcode.com/problems/reach-a-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(target))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #754: Reach a Number
// https://leetcode.com/problems/reach-a-number/
// Difficulty: Medium
// Time: O(sqrt(target))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	sum := 0
	steps := 0

	for sum < target || (sum-target)%2 != 0 {
		steps++
		sum += steps
	}

	return steps
}
```
