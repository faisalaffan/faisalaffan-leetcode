# 3395 — Subsequences With A Unique Middle Mode I

## Deskripsi

**Soal:** [3395. Subsequences With A Unique Middle Mode I](https://leetcode.com/problems/subsequences-with-a-unique-middle-mode-i/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3395: Subsequences with a Unique Middle Mode I
// https://leetcode.com/problems/subsequences-with-a-unique-middle-mode-i/
// Difficulty: Hard
//
// Count subsequences of length 5 where the middle element (index 2) is the
// unique mode. Fix middle index, use prefix/suffix counts, inclusion-exclusion.

import "fmt"

func main() {
	// Example: nums=[1,2,1,2,1] -> 6
	fmt.Println(subsequencesWithMiddleMode([]int{1, 2, 1, 2, 1}))

	// n=5, all distinct
	fmt.Println(subsequencesWithMiddleMode([]int{1, 2, 3, 4, 5}))

	// n=5, all same
	fmt.Println(subsequencesWithMiddleMode([]int{1, 1, 1, 1, 1}))

	// Larger example
	fmt.Println(subsequencesWithMiddleMode([]int{1, 2, 2, 3, 3, 4}))

	// All ones
	fmt.Println(subsequencesWithMiddleMode([]int{1, 1, 1, 1, 1, 1, 1}))
}

const MOD = 1000000007

func subsequencesWithMiddleMode(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}

	// Coordinate compression
  // Membuat map untuk pencarian O(1): key → value
	comp := make(map[int]int)
	for _, v := range nums {
		comp[v] = 1
	}
	m := 0
	for k := range comp {
		comp[k] = m
		m++
	}
  // Membuat slice untuk menyimpan hasil
	arr := make([]int, n)
	for i, v := range nums {
		arr[i] = comp[v]
	}

	// Total count of each value
  // Membuat slice untuk menyimpan hasil
	tot := make([]int, m)
	for _, v := range arr {
		tot[v]++
	}

	// Precompute combinations up to n, choose up to 5
  // Membuat slice 2D untuk DP/tabel
	C := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int, 6)
		C[i][0] = 1
		for j := 1; j <= i && j <= 5; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
		}
	}
	comb := func(a, b int) int {
		if a < b || b < 0 {
			return 0
		}
		return C[a][b]
	}

	ans := 0
  // Membuat slice untuk menyimpan hasil
	cnt := make([]int, m) // prefix count as we sweep

	for i := 0; i < n; i++ {
		x := arr[i]
		cnt[x]++
		remx := tot[x] - cnt[x] // count of x to the right (including i)
		leftOther := i + 1 - cnt[x]
		rightOther := n - i - 1 - (remx - 1) // -1 because i is included in remx

		// Case: x appears >= 3 times in the subsequence (total 5)
		// We need at least 2 more xs besides the middle one.
		// Choose l from left and r from right, l + r >= 2
		// Remaining 2 slots filled with non-x elements
		for l := 0; l <= 2 && l <= cnt[x]-1; l++ {
			for r := 0; r <= 2 && r <= remx-1; r++ {
				if l+r < 2 {
					continue
				}
				if 2-l > leftOther || 2-r > rightOther {
					continue
				}
				ways := comb(cnt[x]-1, l) * comb(remx-1, r) % MOD
				ways = ways * comb(leftOther, 2-l) % MOD
				ways = ways * comb(rightOther, 2-r) % MOD
				ans = (ans + ways) % MOD
			}
		}

		// Case: x appears exactly 2 times in the subsequence.
		// The extra x comes from left or right (not both, since that'd be 3 total).
		// We need to subtract cases where another value y also appears 2+ times.
		// This is complex; for the exact approach, consider:
		// - x appears 2 times: one at middle i, one from left (or right)
		// - Need to ensure no other value y appears 2+ times

		// Subcase: extra x from left
		if cnt[x] >= 2 {
			// cnt[x]-1 ways to pick the left x
			ways := comb(leftOther, 2) * comb(rightOther, 2) % MOD
			ways = ways * (cnt[x] - 1) % MOD
			// Subtract invalid: some y also appears 2+ times
			for y := 0; y < m; y++ {
				if y == x {
					continue
				}
				cntY := cnt[y]
				remY := tot[y] - cnt[y]
				// y appears 2+ times: we need to subtract
				// Case: y appears 2 times in right (both slots on right)
				if remY >= 2 {
					sub := comb(remY, 2) * comb(leftOther, 2) % MOD
					sub = sub * (cnt[x] - 1) % MOD
					ways = (ways - sub + MOD) % MOD
				}
				// Case: y appears 1 on left and 1 on right
				if cntY >= 1 && remY >= 1 {
					// Pick 1 y from left, 1 y from right
					sub := cntY * remY % MOD
					// Remaining: 1 more from left (non-x, non-y), 1 more from right (non-x, non-y)
					leftRest := leftOther - cntY
					rightRest := rightOther - remY
					if leftRest >= 1 && rightRest >= 1 {
						sub = sub * leftRest % MOD
						sub = sub * rightRest % MOD
						sub = sub * (cnt[x] - 1) % MOD
						ways = (ways - sub + MOD) % MOD
					}
				}
				// Case: y appears 2 times on left
				if cntY >= 2 {
					sub := comb(cntY, 2) * comb(rightOther, 2) % MOD
					sub = sub * (cnt[x] - 1) % MOD
					ways = (ways - sub + MOD) % MOD
				}
			}
			ans = (ans + ways) % MOD
		}

		// Subcase: extra x from right
		if remx >= 2 {
			ways := comb(leftOther, 2) * comb(rightOther, 2) % MOD
			ways = ways * (remx - 1) % MOD
			for y := 0; y < m; y++ {
				if y == x {
					continue
				}
				cntY := cnt[y]
				remY := tot[y] - cnt[y]
				if cntY >= 2 {
					sub := comb(cntY, 2) * comb(rightOther, 2) % MOD
					sub = sub * (remx - 1) % MOD
					ways = (ways - sub + MOD) % MOD
				}
				if cntY >= 1 && remY >= 1 {
					sub := cntY * remY % MOD
					leftRest := leftOther - cntY
					rightRest := rightOther - remY
					if leftRest >= 1 && rightRest >= 1 {
						sub = sub * leftRest % MOD
						sub = sub * rightRest % MOD
						sub = sub * (remx - 1) % MOD
						ways = (ways - sub + MOD) % MOD
					}
				}
				if remY >= 2 {
					sub := comb(remY, 2) * comb(leftOther, 2) % MOD
					sub = sub * (remx - 1) % MOD
					ways = (ways - sub + MOD) % MOD
				}
			}
			ans = (ans + ways) % MOD
		}

		// Decrement right count for next iteration
		// (since we're moving past i, what's currently "right" shrinks)
		// Actually cnt is updated at the start of each iteration.
		// The issue is that remx includes i itself.
		// After this iteration, we're moving i forward, so the right counts
		// for the NEXT iteration will not include element i.
		// Actually cnt is already incremented at the start. The right count
		// naturally decreases as i increases.
		// No adjustment needed since we recompute remx each iteration.
	}

	return ans
}
```
