# 0172 — Factorial Trailing Zeroes

## Deskripsi

**Soal:** [0172. Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func trailingZeroes(n int) int`

## Solusi Go

```go
package main

// LeetCode #172: Factorial Trailing Zeroes
// https://leetcode.com/problems/factorial-trailing-zeroes/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}
```
