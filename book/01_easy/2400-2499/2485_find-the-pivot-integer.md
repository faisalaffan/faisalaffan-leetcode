# 2485 — Find The Pivot Integer

## Deskripsi

**Soal:** [2485. Find The Pivot Integer](https://leetcode.com/problems/find-the-pivot-integer/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2485: Find the Pivot Integer
// https://leetcode.com/problems/find-the-pivot-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindThePivotInteger(8)) // 6
	fmt.Println(FindThePivotInteger(1)) // 1
	fmt.Println(FindThePivotInteger(4)) // -1
}

func FindThePivotInteger(n int) int {
	total := n * (n + 1) / 2
	sum := 0
	for x := 1; x <= n; x++ {
		sum += x
		if sum == total-sum+x {
			return x
		}
	}
	return -1
}
```
