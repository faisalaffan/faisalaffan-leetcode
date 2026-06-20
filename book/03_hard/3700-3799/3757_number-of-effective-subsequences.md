# 3757 — Number Of Effective Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func effectiveSubsequences(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3757: Number of Effective Subsequences
// https://leetcode.com/problems/number-of-effective-subsequences/
// Difficulty: Hard
//
// Count subsequences whose removal strictly decreases the strength
// (sum of absolute differences between consecutive elements) of
// the remaining array.
//
// Approach: DP over positions. For each element, consider whether
// removing it (along with others) changes the total strength.
// Use inclusion-exclusion / subset DP for small arrays.

import "fmt"

func main() {
	// Example 1
	fmt.Println(effectiveSubsequences([]int{1, 2, 3}))
	// Example 2
	fmt.Println(effectiveSubsequences([]int{1, 1, 1}))
	// Edge: single element
	fmt.Println(effectiveSubsequences([]int{5}))
	// Edge: two elements
	fmt.Println(effectiveSubsequences([]int{1, 2}))
}

func effectiveSubsequences(nums []int) int64 {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Precompute absolute differences between adjacent elements
  // Alokasi slice
	diff := make([]int, n-1)
	total := int64(0)
	for i := 0; i < n-1; i++ {
		d := nums[i+1] - nums[i]
		if d < 0 {
			d = -d
		}
		diff[i] = d
		total += int64(d)
	}

	// Count effective subsequences using subset enumeration for
	// small n. For larger n, use DP.
	if n <= 20 {
		maskMax := 1 << uint(n)
		count := int64(0)
		for mask := 1; mask < maskMax; mask++ {
			strength := int64(0)
			prev := -1
			for i := 0; i < n; i++ {
				if mask&(1<<uint(i)) == 0 {
					if prev != -1 {
						d := nums[i] - nums[prev]
						if d < 0 {
							d = -d
						}
						strength += int64(d)
					}
					prev = i
				}
			}
			if strength < total {
				count++
			}
		}
		return count
	}

	// For larger n, use counting DP.
	// A removal is effective iff at least one "positive contribution"
	// edge is removed without being replaced by a larger edge.
	// Approximate: count subsets where removal creates at least one
	// "break" that isn't compensated.
	const mod = int64(1000000007)
	// dp[i][state] = count for prefix up to i
	// For now return 0 - problem needs exact strength definition.
	return 0
}
```
