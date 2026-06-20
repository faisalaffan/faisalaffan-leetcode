# 3139 — Minimum Cost To Equalize Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minCostToEqualizeArray(nums []int, cost1, cost2 int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3139: Minimum Cost to Equalize Array
// https://leetcode.com/problems/minimum-cost-to-equalize-array/
// Difficulty: Hard
//
// You can increment one element by 1 (cost1) or two different elements by 1 (cost2).
// Find min cost to make all elements equal, modulo 1e9+7.
// Strategy: try all possible target values from max(nums) up to a bound.
// Use pair operations (cost2) as much as possible since they're cheaper per increment.

import (
	"fmt"
	"math"
)

const MOD = 1000000007

func minCostToEqualizeArray(nums []int, cost1, cost2 int) int {
	n := len(nums)
	minVal, maxVal := nums[0], nums[0]
	var sum int64 = 0
	for _, v := range nums {
		sum += int64(v)
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}

	if n == 1 {
		return 0
	}

	// If cost1*2 <= cost2, just use cost1 for all increments
	if cost1*2 <= cost2 {
		totalCost := int64(0)
		for _, v := range nums {
			totalCost += int64(maxVal-v) * int64(cost1)
		}
		return int(totalCost % MOD)
	}

	ans := int64(math.MaxInt64)
	limit := maxVal + n*2 + 5

	for target := maxVal; target <= limit; target++ {
		totalIncs := int64(target)*int64(n) - sum
		maxDeficit := int64(target - minVal)

		// Max pairs = min(totalIncs/2, totalIncs - maxDeficit)
		pairs := totalIncs / 2
		if pairs > totalIncs-maxDeficit {
			pairs = totalIncs - maxDeficit
		}
		cost := pairs*int64(cost2) + (totalIncs-2*pairs)*int64(cost1)
		if cost < ans {
			ans = cost
		}
	}

	return int(ans % MOD)
}

func main() {
	// Test case 1
	nums := []int{4, 1}
	cost1 := 5
	cost2 := 2
	fmt.Println("Test 1:", minCostToEqualizeArray(nums, cost1, cost2))
	// Expected: 15

	// Test case 2: all same
	nums2 := []int{3, 3, 3}
	fmt.Println("Test 2:", minCostToEqualizeArray(nums2, 2, 3))
	// Expected: 0

	// Test case 3: single element
	nums3 := []int{5}
	fmt.Println("Test 3:", minCostToEqualizeArray(nums3, 1, 2))
	// Expected: 0

	// Test case 4: cost1 cheaper
	nums4 := []int{1, 5}
	fmt.Println("Test 4:", minCostToEqualizeArray(nums4, 1, 10))
	// Expected: 4 (use cost1: (5-1)*1 = 4)

	// Test case 5: pair cheaper
	nums5 := []int{1, 5}
	fmt.Println("Test 5:", minCostToEqualizeArray(nums5, 5, 1))
	// Expected: 6? Let's compute: target=5, incs=[4,0], pairs=0, cost=4*5=20.
	// target=6, incs=[5,1], pairs=1, cost=1*1+(6-2)*5=1+20=21
	// target=5: incs=[4,0], totalIncs=4, maxDeficit=4, pairs=0, cost=4*5=20
	// Hmm, answer depends on better target
}
```
