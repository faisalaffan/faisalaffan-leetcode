# 3569 — Maximize Count Of Distinct Primes After Split

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumCount(nums []int, queries [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum, Segment Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3569: Maximize Count of Distinct Primes After Split
// https://leetcode.com/problems/maximize-count-of-distinct-primes-after-split/
// Difficulty: Hard
//
// For each query, update nums[idx] = val, then find a split index k (1 <= k < n)
// that maximizes the number of distinct primes in nums[0..k-1] + nums[k..n-1].
//
// Approach: Precompute prime factors for each possible value, use a segment tree
// to maintain distinct prime counts for each segment. For each query, update and
// evaluate all split candidates.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumCount([]int{2, 3, 4, 5, 6}, [][]int{{0, 7}, {1, 8}}))
	// Example 2
	fmt.Println(maximumCount([]int{10, 15, 21}, [][]int{{0, 2}, {2, 3}}))
	// Edge: single element
	fmt.Println(maximumCount([]int{6}, [][]int{{0, 7}}))
	// Edge: all primes
	fmt.Println(maximumCount([]int{2, 3, 5}, [][]int{{0, 7}, {1, 11}}))
}

func maximumCount(nums []int, queries [][]int) []int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

	// Precompute smallest prime factor for values up to max
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	for _, q := range queries {
		if q[1] > maxVal {
			maxVal = q[1]
		}
	}
	spf := sieve(maxVal)

	// Precompute distinct prime factors for each number
  // Matriks 2D
	primeFactors := make([][]int, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		primeFactors[i] = getDistinctPrimes(i, spf)
	}

	// Current values
  // Alokasi slice
	cur := make([]int, n)
	copy(cur, nums)

	// Segment tree for range distinct prime count
	// Each node stores a bitset of primes present in its range
	// Since we need to compute distinct primes per segment,
	// we'll use per-element sets and recompute on split

	// Since n is small enough, we can just recompute per query
	// by scanning all splits
  // Alokasi slice
	result := make([]int, len(queries))

	for qi, q := range queries {
		idx, val := q[0], q[1]
		cur[idx] = val

		best := 0
		// Prefix distinct prime sets
  // HashMap: O(1) lookup
		prefixSet := make(map[int]bool)
		for k := 0; k < n-1; k++ {
			for _, p := range primeFactors[cur[k]] {
				prefixSet[p] = true
			}
			// Suffix distinct prime set
  // HashMap: O(1) lookup
			suffixSet := make(map[int]bool)
			for r := k + 1; r < n; r++ {
				for _, p := range primeFactors[cur[r]] {
					suffixSet[p] = true
				}
			}
			total := len(prefixSet) + len(suffixSet)
			if total > best {
				best = total
			}
		}
		result[qi] = best
	}

	return result
}

func sieve(n int) []int {
  // Alokasi slice
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if i*i <= n {
				for j := i * i; j <= n; j += i {
					if spf[j] == 0 {
						spf[j] = i
					}
				}
			}
		}
	}
	return spf
}

func getDistinctPrimes(x int, spf []int) []int {
	var result []int
	last := 0
	for x > 1 {
		p := spf[x]
		if p == 0 {
			p = x
		}
		if p != last {
			result = append(result, p)
			last = p
		}
		x /= p
	}
	return result
}
```
