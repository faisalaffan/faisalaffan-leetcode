# 3704 — Count No Zero Pairs That Sum To N

## Deskripsi

**Soal:** [3704. Count No Zero Pairs That Sum To N](https://leetcode.com/problems/count-no-zero-pairs-that-sum-to-n/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Digit DP. Process n digit by digit. At each position,

## Solusi Go

```go
package main

// LeetCode #3704: Count No-Zero Pairs That Sum to N
// https://leetcode.com/problems/count-no-zero-pairs-that-sum-to-n/
// Difficulty: Hard
//
// A "no-zero" positive integer has no digit 0 in its decimal
// representation. Count pairs (a, b) where a and b are no-zero
// positive integers and a + b = n.
//
// Approach: Digit DP. Process n digit by digit. At each position,
// count ways to split the digit sum between a and b, ensuring
// neither a nor b has a zero digit.

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1
	fmt.Println(countNoZeroPairs(11))
	// Example 2
	fmt.Println(countNoZeroPairs(20))
	// Edge: small n
	fmt.Println(countNoZeroPairs(2))
}

func countNoZeroPairs(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)

  // Membuat slice untuk menyimpan hasil
	memo := make([][2][2]int64, m)
  // Iterasi seluruh elemen
	for i := range memo {
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	digits := make([]int, m)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	var dfs func(pos int, borrow int, started int) int64
	dfs = func(pos int, borrow int, started int) int64 {
		if pos < 0 {
			if borrow == 0 {
				return 1
			}
			return 0
		}
		if memo[pos][borrow][started] != -1 {
			return memo[pos][borrow][started]
		}

		d := digits[pos] - borrow
		var total int64

		if d >= 0 {
			// Both digits are non-zero
			if started == 1 {
				ways := twoSumWays(d)
				total = dfs(pos-1, 0, 1) * int64(ways)
				total += dfs(pos-1, 1, 1) * int64(twoSumWays(d+10))
			}
		} else {
			total = dfs(pos-1, 1, started) * int64(twoSumWays(d+10))
		}

		if pos < m-1 {
			// One of the numbers has leading zeros
			// Only count if this position has value (not all zeros)
		}

		memo[pos][borrow][started] = total
		return total
	}

	return dfs(m-1, 0, 1)
}

func twoSumWays(target int) int {
	if target < 2 || target > 18 {
		return 0
	}
	// Count of digit pairs (x,y) with x,y in [1,9], x+y = target
	cnt := 0
	for x := 1; x <= 9; x++ {
		y := target - x
		if y >= 1 && y <= 9 {
			cnt++
		}
	}
	return cnt
}
```
