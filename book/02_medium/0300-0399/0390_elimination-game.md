# 0390 — Elimination Game

## Deskripsi

**Soal:** [0390. Elimination Game](https://leetcode.com/problems/elimination-game/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func lastRemaining(n int) int`

## Solusi Go

```go
package main

// LeetCode #390: Elimination Game
// https://leetcode.com/problems/elimination-game/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func lastRemaining(n int) int {
	head := 1
	remaining := n
	step := 1
	leftToRight := true

	for remaining > 1 {
		if leftToRight || remaining%2 == 1 {
			head += step
		}
		remaining /= 2
		step *= 2
		leftToRight = !leftToRight
	}
	return head
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", lastRemaining(9))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", lastRemaining(1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", lastRemaining(100))
	// Expected: 54
}
```
