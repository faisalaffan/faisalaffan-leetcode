# 2240 — Number Of Ways To Buy Pens And Pencils

## Deskripsi

**Soal:** [2240. Number Of Ways To Buy Pens And Pencils](https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(total / cost1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64`

## Solusi Go

```go
package main

// LeetCode #2240: Number of Ways to Buy Pens and Pencils
// https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/
// Difficulty: Medium
// Time: O(total / cost1) | Space: O(1)

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ways int64 = 0
	for pens := 0; pens*cost1 <= total; pens++ {
		remaining := total - pens*cost1
		ways += int64(remaining/cost2) + 1
		if cost2 == 0 {
			break
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
	// Expected: 9

	// Test case 2
	fmt.Println(waysToBuyPensPencils(5, 10, 10))
	// Expected: 1

	// Test case 3
	fmt.Println(waysToBuyPensPencils(100, 1, 1))
	// Expected: 5151
}
```
