# 1735 — Count Ways To Make Array With Product

## Deskripsi

**Soal:** [1735. Count Ways To Make Array With Product](https://leetcode.com/problems/count-ways-to-make-array-with-product/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func powMod(a, e int) int`

> **Ide Kunci:** Prime factorization + combinatorics (stars and bars).

## Solusi Go

```go
package main

// LeetCode #1735: Count Ways to Make Array With Product
// https://leetcode.com/problems/count-ways-to-make-array-with-product/
// Difficulty: Hard
//
// Approach: Prime factorization + combinatorics (stars and bars).
// For each query [k, n]:
//   1. Factorize n into prime factors with counts.
//   2. For each prime with count c, distribute c identical items into k distinct
//      positions => C(c + k - 1, k - 1) ways.
//   3. Multiply results for all primes (primes are independent).

import (
	"fmt"
)

const MOD = 1_000_000_007

func powMod(a, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		e >>= 1
	}
	return res
}

func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	num, den := 1, 1
	for i := 0; i < r; i++ {
		num = (num * (n - i)) % MOD
		den = (den * (i + 1)) % MOD
	}
	return (num * powMod(den, MOD-2)) % MOD
}

func factorize(n int) map[int]int {
  // Membuat map untuk pencarian O(1): key → value
	factors := make(map[int]int)
	for p := 2; p*p <= n; p++ {
		for n%p == 0 {
			factors[p]++
			n /= p
		}
	}
	if n > 1 {
		factors[n]++
	}
	return factors
}

func waysToFillArray(queries [][]int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for idx, q := range queries {
		k, n := q[0], q[1]
		factors := factorize(n)
		res := 1
		for _, cnt := range factors {
			res = (res * nCr(cnt+k-1, k-1)) % MOD
		}
		ans[idx] = res
	}
	return ans
}

func main() {
	// Example test cases
	queries := [][]int{{2, 6}, {5, 1}, {73, 660}}
	result := waysToFillArray(queries)
	fmt.Println("queries=[[2,6],[5,1],[73,660]]", "→", result)
	// Expected: [4, 1, 50734910]

	// Additional tests
	fmt.Println("queries=[[1,1]] →", waysToFillArray([][]int{{1, 1}}))
	fmt.Println("queries=[[2,2]] →", waysToFillArray([][]int{{2, 2}}))
	fmt.Println("queries=[[3,4]] →", waysToFillArray([][]int{{3, 4}}))
}
```
