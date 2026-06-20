# 2550 — Count Collisions Of Monkeys On A Polygon

## Deskripsi

**Soal:** [2550. Count Collisions Of Monkeys On A Polygon](https://leetcode.com/problems/count-collisions-of-monkeys-on-a-polygon/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func monkeyMove(n int) int`

## Solusi Go

```go
package main

// LeetCode #2550: Count Collisions of Monkeys on a Polygon
// https://leetcode.com/problems/count-collisions-of-monkeys-on-a-polygon/
// Difficulty: Medium
// Time: O(log n) | Space: O(1)

import "fmt"

func monkeyMove(n int) int {
	// Total ways: 2^n. Only 2 ways avoid collisions: all clockwise or all anticlockwise
	// Answer: (2^n - 2 + mod) % mod
	const mod = 1_000_000_007

	pow2 := 1
	base := 2
	exp := n
	for exp > 0 {
		if exp&1 == 1 {
			pow2 = pow2 * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}

	ans := (pow2 - 2 + mod) % mod
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", monkeyMove(3))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", monkeyMove(4))
	// Expected: 14

	// Test case 3
	fmt.Println("Test 3:", monkeyMove(100))
}
```
