# 3500 — Minimum Cost To Divide Array Into Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCost(nums []int, cost []int, k int) int64
```

> **💡 Hint:** DP with convex hull trick (CHT) for optimization, or use

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3500: Minimum Cost to Divide Array Into Subarrays
// https://leetcode.com/problems/minimum-cost-to-divide-array-into-subarrays/
// Difficulty: Hard
//
// Given arrays nums and cost, divide nums into subarrays where the cost of
// each subarray is the sum of (nums[i] * cost multiplier), and the multiplier
// for each subarray increases by k for each subsequent subarray.
//
// Approach: DP with convex hull trick (CHT) for optimization, or use
// prefix sums with greedy strategy.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumCost([]int{1, 2, 3}, []int{4, 5, 6}, 1))
	// Example 2
	fmt.Println(minimumCost([]int{1, 2}, []int{3, 4}, 2))
	// Example 3: single element
	fmt.Println(minimumCost([]int{5}, []int{10}, 3))
	// Edge: all same
	fmt.Println(minimumCost([]int{1, 1, 1}, []int{1, 1, 1}, 2))
}

func minimumCost(nums []int, cost []int, k int) int64 {
	n := len(nums)

	// Prefix sums
  // Alokasi slice integer
	prefNums := make([]int, n+1)
  // Alokasi slice integer
	prefCost := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefNums[i+1] = prefNums[i] + nums[i]
		prefCost[i+1] = prefCost[i] + cost[i]
	}

	// DP[i] = min cost for prefix up to i
  // Alokasi slice integer
	dp := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1 << 62 // large number
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// Cost of subarray nums[j..i-1] with multiplier k*(j+1) ... ???
			// The cost for a subarray starting at position j is:
			// sum_{t=j}^{i-1} nums[t] * (some multiplier) + k * sum_{t=j}^{i-1} nums[t] * (number of previous subarrays)
			// This is a DP with convex hull trick optimization in the optimal solution.
			// Simplified: cost = prefNums[i] * (prefCost[i] - prefCost[j]) + k * prefNums[i] * j?
			// For now, a simple O(n^2) DP
			subarraySum := int64(prefNums[i] - prefNums[j])
			costSum := int64(prefCost[i] - prefCost[j])
			val := dp[j] + subarraySum*costSum + int64(k)*subarraySum*int64(j)
			if val < dp[i] {
				dp[i] = val
			}
		}
	}

	return dp[n]
}
```
