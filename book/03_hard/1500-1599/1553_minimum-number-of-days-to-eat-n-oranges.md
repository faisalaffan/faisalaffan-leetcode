# 1553 — Minimum Number Of Days To Eat N Oranges

## Deskripsi

**Soal:** [1553. Minimum Number Of Days To Eat N Oranges](https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1553: Minimum Number of Days to Eat N Oranges
// https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/
// Difficulty: Hard
//
// Memoized recursion:
// - If n is divisible by 2, we can eat n/2 oranges in one day (after eating n%2 oranges one by one).
// - If n is divisible by 3, we can eat 2*n/3 oranges in one day (after eating n%3 oranges one by one).
// - Otherwise, eat 1 orange.
// - Use memoization to avoid recomputation.
// - Since n can be up to 2*10^9, we use a map for memoization.

import (
	"fmt"
)

func main() {
	// Example: 10 -> 4; 6 -> 3
	fmt.Println(minDays(10))
	fmt.Println(minDays(6))

	// Additional tests
	fmt.Println(minDays(1))
	fmt.Println(minDays(2))
	fmt.Println(minDays(3))
	fmt.Println(minDays(50))
	fmt.Println(minDays(100))
}

func minDays(n int) int {
  // Membuat map untuk pencarian O(1): key → value
	memo := make(map[int]int)
	return dp(n, memo)
}

func dp(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}

	// Option 1: eat one at a time to make n divisible by 2, then eat half
	option1 := n%2 + 1 + dp(n/2, memo)
	// Option 2: eat one at a time to make n divisible by 3, then eat 2/3
	option2 := n%3 + 1 + dp(n/3, memo)

	result := min(option1, option2)
	memo[n] = result
	return result
}
```
