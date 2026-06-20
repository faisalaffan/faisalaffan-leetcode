# 2597 — The Number Of Beautiful Subsets

## Deskripsi

**Soal:** [2597. The Number Of Beautiful Subsets](https://leetcode.com/problems/the-number-of-beautiful-subsets/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func beautifulSubsets(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2597: The Number of Beautiful Subsets
// https://leetcode.com/problems/the-number-of-beautiful-subsets/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Group by residue modulo k
  // Membuat map untuk pencarian O(1): key → value
	grouped := make(map[int][]int)
	for v := range freq {
		grouped[v%k] = append(grouped[v%k], v)
	}

	var ans int = 1 // empty subset
	for _, vals := range grouped {
		sort.Ints(vals)
		// DP within each group: dp0 (ways not taking current), dp1 (ways taking current)
		dp0, dp1 := 1, 0 // dp0 = ways for "not taking previous value", dp1 = ways for "taking previous value"
		for i, v := range vals {
			ways0 := dp0 + dp1 // skip current value
			ways1 := dp0 * (1 << uint(freq[v]-1))
			if i > 0 && v-vals[i-1] == k {
				ways1 = dp0 * ((1 << uint(freq[v])) - 1)
			} else {
				ways1 = (dp0 + dp1) * ((1 << uint(freq[v])) - 1)
			}
			dp0, dp1 = ways0, ways1
		}
		ans *= (dp0 + dp1)
	}

	return ans - 1 // exclude empty subset
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubsets([]int{2, 4, 6}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", beautifulSubsets([]int{1}, 1))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", beautifulSubsets([]int{1, 2, 3, 4}, 1))
	// Expected: 7
}
```
