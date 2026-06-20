# 3317 — Find The Number Of Possible Ways For An Event

## Deskripsi

**Soal:** [3317. Find The Number Of Possible Ways For An Event](https://leetcode.com/problems/find-the-number-of-possible-ways-for-an-event/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Use Stirling numbers of the second kind S(n,k) for partitioning

## Solusi Go

```go
package main

// LeetCode #3317: Find the Number of Possible Ways for an Event
// https://leetcode.com/problems/find-the-number-of-possible-ways-for-an-event/
// Difficulty: Hard
//
// Given n performers, x stages, and score range [1, y], count the number of
// distinct events possible. Each performer is assigned to a stage (may be
// empty). Each non-empty stage gets a score. Two events differ if any performer
// is on a different stage or any band gets a different score.
//
// Approach: Use Stirling numbers of the second kind S(n,k) for partitioning
// n performers into k non-empty groups. Multiply by P(x,k) = x!/(x-k)! for
// choosing ordered stages, and y^k for score assignments.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfWays(3, 3, 4))
	// Example 2
	fmt.Println(numberOfWays(2, 3, 4))
	// Example 3
	fmt.Println(numberOfWays(1, 2, 3))
	// Edge: single performer, single stage
	fmt.Println(numberOfWays(1, 1, 5))
	// Edge: n > x
	fmt.Println(numberOfWays(5, 3, 2))
}

const mod = 1000000007

func numberOfWays(n int, x int, y int) int {
	// Precompute factorials
  // Membuat slice untuk menyimpan hasil
	fact := make([]int64, x+1)
	fact[0] = 1
	for i := 1; i <= x; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}

	// Precompute inverse factorials
  // Membuat slice untuk menyimpan hasil
	invFact := make([]int64, x+1)
	invFact[x] = powMod(fact[x], mod-2)
	for i := x - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % mod
	}

	// Precompute powers of y
  // Membuat slice untuk menyimpan hasil
	powY := make([]int64, x+1)
	powY[0] = 1
	for k := 1; k <= x; k++ {
		powY[k] = powY[k-1] * int64(y) % mod
	}

	// Stirling numbers of the second kind S(n, k) using DP
  // Membuat slice untuk menyimpan hasil
	stirling := make([]int64, x+1)
	stirling[0] = 1
	for i := 1; i <= n; i++ {
		kMax := i
		if kMax > x {
			kMax = x
		}
		for k := kMax; k >= 1; k-- {
			stirling[k] = (int64(k)*stirling[k] + stirling[k-1]) % mod
		}
		stirling[0] = 0
	}

	// Sum over k
	var ans int64
	limit := n
	if limit > x {
		limit = x
	}
	for k := 1; k <= limit; k++ {
		// P(x, k) = x! / (x-k)!
		perm := fact[x] * invFact[x-k] % mod
		term := perm * stirling[k] % mod
		term = term * powY[k] % mod
		ans = (ans + term) % mod
	}

	return int(ans)
}

func powMod(a int64, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}
```
