# 0829 — Consecutive Numbers Sum

## Deskripsi

**Soal:** [0829. Consecutive Numbers Sum](https://leetcode.com/problems/consecutive-numbers-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func consecutiveNumbersSum(n int) int`

> **Ide Kunci:** For each possible length m, check if N = m*k + m*(m-1)/2.

## Solusi Go

```go
package main

// LeetCode #829: Consecutive Numbers Sum
// https://leetcode.com/problems/consecutive-numbers-sum/
// Difficulty: Hard
// Approach: For each possible length m, check if N = m*k + m*(m-1)/2.
// Which simplifies to: (N - m*(m-1)/2) % m == 0 and > 0.

import "fmt"

func consecutiveNumbersSum(n int) int {
	count := 0
	for m := 1; m*(m-1)/2 < n; m++ {
		rem := n - m*(m-1)/2
		if rem%m == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(consecutiveNumbersSum(5))  // Expected: 2
	fmt.Println(consecutiveNumbersSum(9))  // Expected: 3
	fmt.Println(consecutiveNumbersSum(15)) // Additional test: 15 = 15, 7+8, 4+5+6, 1+2+3+4+5 => 4
}
```
