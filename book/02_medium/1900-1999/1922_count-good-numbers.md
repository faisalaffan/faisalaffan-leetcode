# 1922 — Count Good Numbers

## Deskripsi

**Soal:** [1922. Count Good Numbers](https://leetcode.com/problems/count-good-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #1922: Count Good Numbers
// https://leetcode.com/problems/count-good-numbers/
// Difficulty: Medium

import "fmt"

const mod1922 = 1000000007

func main() {
	fmt.Println(CountGoodNumbers(1))
	fmt.Println(CountGoodNumbers(4))
	fmt.Println(CountGoodNumbers(50))
}

// Time: O(log n), Space: O(1)
func CountGoodNumbers(n int64) int {
	evenPositions := (n + 1) / 2 // positions 0, 2, 4, ... (0-indexed)
	oddPositions := n / 2        // positions 1, 3, 5, ...

	// Even positions: 5 choices (0,2,4,6,8)
	// Odd positions: 4 choices (2,3,5,7)
	return int(powMod(5, evenPositions) * powMod(4, oddPositions) % mod1922)
}

func powMod(base int64, exp int64) int64 {
	result := int64(1)
	b := base % mod1922
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = (result * b) % mod1922
		}
		b = (b * b) % mod1922
		e >>= 1
	}
	return result
}
```
