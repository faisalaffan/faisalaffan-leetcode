# 3621 — Number Of Integers With Popcount Depth Equal To K I

## Deskripsi

**Soal:** [3621. Number Of Integers With Popcount Depth Equal To K I](https://leetcode.com/problems/number-of-integers-with-popcount-depth-equal-to-k-i/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Precompute popcount-depth for all values up to 1000

## Solusi Go

```go
package main

// LeetCode #3621: Number of Integers With Popcount-Depth Equal to K I
// https://leetcode.com/problems/number-of-integers-with-popcount-depth-equal-to-k-i/
// Difficulty: Hard
//
// Popcount-depth of integer x is the number of times we need to
// replace x with popcount(x) until x becomes 1. Count numbers
// in [1, n] with popcount-depth exactly k.
//
// Approach: Precompute popcount-depth for all values up to 1000
// (max needed since popcount of any n <= 10^15 is at most 50).
// Then count how many numbers in [1, n] have a given popcount,
// and check which popcount values have the right depth.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfIntegers(10, 1))
	// Example 2
	fmt.Println(numberOfIntegers(100, 2))
	// Edge: n = 1
	fmt.Println(numberOfIntegers(1, 1))
}

func numberOfIntegers(n int64, k int) int64 {
	// Precompute depth for all possible popcount values (1..60)
  // Membuat slice untuk menyimpan hasil
	depth := make([]int, 61)
	for i := 2; i <= 60; i++ {
		depth[i] = depth[popcount(i)] + 1
	}

	// Counts[n][b] = numbers in [0, n-1] with exactly b bits set
	s := fmt.Sprintf("%b", n)
	m := len(s)

  // Membuat slice 2D untuk DP/tabel
	memo := make([][][]int64, m)
  // Iterasi seluruh elemen
	for i := range memo {
		memo[i] = make([][]int64, 2)
		for j := range memo[i] {
			memo[i][j] = make([]int64, 61)
			for b := range memo[i][j] {
				memo[i][j][b] = -1
			}
		}
	}

	var dfs func(pos int, tight int, bits int) int64
	dfs = func(pos int, tight int, bits int) int64 {
		if pos == m {
			if bits == 0 {
				return 0
			}
			return 1
		}
		if memo[pos][tight][bits] != -1 {
			return memo[pos][tight][bits]
		}

		limit := byte('1')
		if tight == 1 {
			limit = s[pos]
		}

		var total int64
		for d := byte('0'); d <= limit; d++ {
			nt := 0
			if tight == 1 && d == limit {
				nt = 1
			}
			nb := bits
			if d == '1' {
				nb++
			}
			total += dfs(pos+1, nt, nb)
		}

		memo[pos][tight][bits] = total
		return total
	}

	// Count numbers with each bit count, then filter by depth
	var ans int64
	for bits := 1; bits <= 60; bits++ {
		if depth[bits] == k-1 {
			ans += dfs(0, 1, 0)
		}
	}
	return ans
}

func popcount(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}
```
