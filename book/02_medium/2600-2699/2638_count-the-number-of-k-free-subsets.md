# 2638 — Count The Number Of K Free Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countKFreeSubsets(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2638: Count the Number of K-Free Subsets
// https://leetcode.com/problems/count-the-number-of-k-free-subsets/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func countKFreeSubsets(nums []int, k int) int64 {
  // Sort O(n log n)
	sort.Ints(nums)
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Group by remainder modulo k
  // HashMap: O(1) lookup
	groups := make(map[int][]int)
	for v := range freq {
		groups[v%k] = append(groups[v%k], v)
	}

	ans := int64(1) // empty subset
	for _, vals := range groups {
  // Sort O(n log n)
		sort.Ints(vals)
		// DP within group: O(n) with constraint that consecutive vals with diff k can't both be taken
		dp0, dp1 := int64(1), int64(0)
		for i, v := range vals {
			waysSkip := dp0 + dp1
			var waysTake int64
			if i > 0 && v-vals[i-1] == k {
				waysTake = dp0 * ((int64(1) << uint(freq[v])) - 1)
			} else {
				waysTake = (dp0 + dp1) * ((int64(1) << uint(freq[v])) - 1)
			}
			dp0, dp1 = waysSkip, waysTake
		}
		ans *= (dp0 + dp1)
	}

	return ans - 1 // exclude empty subset
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countKFreeSubsets([]int{1, 2, 3}, 1))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", countKFreeSubsets([]int{1, 2, 3, 4}, 2))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", countKFreeSubsets([]int{1, 3, 5}, 2))
	// Expected: 4
}
```
