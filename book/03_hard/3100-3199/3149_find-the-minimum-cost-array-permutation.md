# 3149 — Find The Minimum Cost Array Permutation

## Deskripsi

**Soal:** [3149. Find The Minimum Cost Array Permutation](https://leetcode.com/problems/find-the-minimum-cost-array-permutation/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func findMinCostArrayPermutation(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #3149: Find the Minimum Cost Array Permutation
// https://leetcode.com/problems/find-the-minimum-cost-array-permutation/
// Difficulty: Hard
//
// Given an integer array nums of length n, find a permutation arr of [0, 1, ..., n-1]
// that minimizes: arr[0] + sum_{i=1}^{n-1} |arr[i] - nums[arr[i-1]]|
// Use DP with bitmask (TSP-like). Among optimal permutations, return lexicographically smallest.

import (
	"fmt"
	"math"
)

func findMinCostArrayPermutation(nums []int) []int {
	n := len(nums)
	totalMasks := 1 << n

	// dp[mask][last] = min cost to form subset `mask` ending with `last`
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, totalMasks)
  // Membuat slice 2D untuk DP/tabel
	parent := make([][]int, totalMasks) // to reconstruct path

	for mask := 0; mask < totalMasks; mask++ {
		dp[mask] = make([]int, n)
		parent[mask] = make([]int, n)
		for i := 0; i < n; i++ {
			dp[mask][i] = math.MaxInt32
			parent[mask][i] = -1
		}
	}

	// Base: start with any single element
	for i := 0; i < n; i++ {
		dp[1<<i][i] = i // cost = arr[0] = i (first element value)
		parent[1<<i][i] = -1
	}

	// DP over masks
	for mask := 0; mask < totalMasks; mask++ {
		for last := 0; last < n; last++ {
			if dp[mask][last] == math.MaxInt32 {
				continue
			}
			for nxt := 0; nxt < n; nxt++ {
				if mask&(1<<nxt) != 0 {
					continue
				}
				newMask := mask | (1 << nxt)
				cost := dp[mask][last] + abs(nxt-nums[last])
				if cost < dp[newMask][nxt] {
					dp[newMask][nxt] = cost
					parent[newMask][nxt] = last
				}
			}
		}
	}

	fullMask := totalMasks - 1

	// Find min cost and last element (preferring lexicographically smaller full permutation)
	minCost := math.MaxInt32
  // Membuat slice untuk menyimpan hasil
	bestCandidates := make([]int, 0)

	for last := 0; last < n; last++ {
		if dp[fullMask][last] < minCost {
			minCost = dp[fullMask][last]
			bestCandidates = []int{last}
		} else if dp[fullMask][last] == minCost {
			bestCandidates = append(bestCandidates, last)
		}
	}

	// Reconstruct and pick lexicographically smallest
  // Membuat slice untuk menyimpan hasil
	bestPerm := make([]int, n)
	first := true

	for _, last := range bestCandidates {
  // Membuat slice untuk menyimpan hasil
		perm := make([]int, n)
		pos := n - 1
		mask := fullMask
		cur := last

		for cur != -1 {
			perm[pos] = cur
			prev := parent[mask][cur]
			mask ^= (1 << cur)
			cur = prev
			pos--
		}

		if first || lexSmaller(perm, bestPerm) {
			copy(bestPerm, perm)
			first = false
		}
	}

	return bestPerm
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lexSmaller(a, b []int) bool {
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return false
}

func main() {
	// Test case 1: minimal
	nums := []int{1, 2}
	fmt.Println("Test 1:", findMinCostArrayPermutation(nums))

	// Test case 2: 3 elements
	nums2 := []int{0, 1, 2}
	fmt.Println("Test 2:", findMinCostArrayPermutation(nums2))

	// Test case 3: 4 elements
	nums3 := []int{0, 2, 1, 3}
	fmt.Println("Test 3:", findMinCostArrayPermutation(nums3))

	// Test case 4: single element
	nums4 := []int{0}
	fmt.Println("Test 4:", findMinCostArrayPermutation(nums4))
	// Expected: [0]

	// Test case 5
	nums5 := []int{1, 0, 2}
	fmt.Println("Test 5:", findMinCostArrayPermutation(nums5))
}
```
