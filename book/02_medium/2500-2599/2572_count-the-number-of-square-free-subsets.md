# 2572 — Count The Number Of Square Free Subsets

## Deskripsi

**Soal:** [2572. Count The Number Of Square Free Subsets](https://leetcode.com/problems/count-the-number-of-square-free-subsets/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 2^p)  
**Kompleksitas Ruang:** O(2^p)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func squareFreeSubsets(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2572: Count the Number of Square-Free Subsets
// https://leetcode.com/problems/count-the-number-of-square-free-subsets/
// Difficulty: Medium
// Time: O(n * 2^p) | Space: O(2^p)

import "fmt"

func squareFreeSubsets(nums []int) int {
	const mod = 1_000_000_007

	// Primes up to 30 (since nums[i] <= 30)
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	p := len(primes)

	// Map each number to its prime mask
  // Membuat slice untuk menyimpan hasil
	primeMask := make([]int, 31)
	for i := 1; i <= 30; i++ {
		mask := 0
		x := i
		for j, prime := range primes {
			cnt := 0
			for x%prime == 0 {
				x /= prime
				cnt++
			}
			if cnt > 1 {
				mask = -1 // Not square-free
				break
			}
			if cnt == 1 {
				mask |= (1 << j)
			}
		}
		primeMask[i] = mask
	}

	// Count frequency of each number
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 31)
	for _, v := range nums {
		freq[v]++
	}

	// DP: dp[mask] = number of ways to get this mask
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, 1<<p)
	dp[0] = 1

	for val := 1; val <= 30; val++ {
		if freq[val] == 0 {
			continue
		}
		mask := primeMask[val]
		if mask < 0 {
			continue
		}

		// We have freq[val] copies of this value
		// Each copy can be either taken or not
		// Ways to take any subset of freq[val] copies = 2^freq[val] - 1
		pow := 1
		for i := 0; i < freq[val]; i++ {
			pow = (pow * 2) % mod
		}
		ways := (pow - 1 + mod) % mod

		// Update DP (knapsack style)
  // Membuat slice untuk menyimpan hasil
		newDP := make([]int, 1<<p)
		copy(newDP, dp)
		for m := 0; m < (1 << p); m++ {
			if dp[m] == 0 {
				continue
			}
			newMask := m | mask
			newDP[newMask] = (newDP[newMask] + dp[m]*ways) % mod
		}
		dp = newDP
	}

	// Sum all non-empty subsets: any mask (0 included for subsets only containing 1)
	ans := (dp[0] - 1 + mod) % mod // non-empty subsets of number 1 only
	for m := 1; m < (1 << p); m++ {
		ans = (ans + dp[m]) % mod
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", squareFreeSubsets([]int{3, 4, 4, 5}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", squareFreeSubsets([]int{1}))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", squareFreeSubsets([]int{1, 2, 3, 4}))
	// Expected: 7
}
```
