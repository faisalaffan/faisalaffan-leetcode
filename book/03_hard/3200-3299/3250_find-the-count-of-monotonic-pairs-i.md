# 3250 — Find The Count Of Monotonic Pairs I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countOfPairs(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum, Monotonic Stack

**Waktu:** O(n * M) where M = max(nums) ≤ 50  |  **Ruang:** O(M)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3250: Find the Count of Monotonic Pairs I
// https://leetcode.com/problems/find-the-count-of-monotonic-pairs-i/
// Difficulty: Hard
//
// Given array nums of length n, count pairs (arr1, arr2) such that:
//   - arr1 is non-decreasing
//   - arr2 is non-increasing
//   - arr1[i] + arr2[i] == nums[i] for all i
//
// DP with prefix sums. Since arr2[i] = nums[i] - arr1[i], we track arr1.
// Condition: arr1[i-1] <= arr1[i] AND arr2[i-1] >= arr2[i]
// → nums[i-1] - arr1[i-1] >= nums[i] - arr1[i]
// → arr1[i] - arr1[i-1] >= nums[i] - nums[i-1]
// So: arr1[i] >= arr1[i-1] AND arr1[i] >= arr1[i-1] + nums[i] - nums[i-1]
//
// Time: O(n * M) where M = max(nums) ≤ 50
// Space: O(M)
//
// Constraints (for Part I): n ≤ 2000, nums[i] ≤ 50

import "fmt"

func main() {
	// Example 1: [2,3,2] => 4
	fmt.Println(countOfPairs([]int{2, 3, 2}))
	// Example 2: [5,5,5,5] => 126
	fmt.Println(countOfPairs([]int{5, 5, 5, 5}))
	// Example 3: [1,2,3,4] => 5
	fmt.Println(countOfPairs([]int{1, 2, 3, 4}))
	// Example 4: single element
	fmt.Println(countOfPairs([]int{5}))
	// Example 5: [1,1,1] => 10
	fmt.Println(countOfPairs([]int{1, 1, 1}))
}

const MOD = 1_000_000_007

func countOfPairs(nums []int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// dp[j] = number of ways where arr1[current] = j
  // Alokasi slice
	dp := make([]int64, maxVal+1)
	for j := 0; j <= nums[0]; j++ {
		dp[j] = 1
	}

	for i := 1; i < n; i++ {
		prev := nums[i-1]
		curr := nums[i]

		// prefix sums of dp
  // Alokasi slice
		prefix := make([]int64, maxVal+2)
		for j := 0; j <= maxVal; j++ {
			prefix[j+1] = (prefix[j] + dp[j]) % MOD
		}

  // Alokasi slice
		ndp := make([]int64, maxVal+1)
		for j := 0; j <= curr; j++ {
			// maxPrev = min(j, j + prev - curr)
			maxPrev := j
			if j+prev-curr < maxPrev {
				maxPrev = j + prev - curr
			}

			// We need dp[i-1][0..maxPrev] where dp[i-1][k] has arr1[i-1]=k
			// arr2[i-1] = prev - k >= arr2[i] = curr - j
			// → prev - k >= curr - j
			// → k <= prev - curr + j = j + prev - curr
			// Also k <= j for non-decreasing arr1
			// So maxPrev = min(j, j + prev - curr)
			if maxPrev >= 0 {
				// Actually we need to reconsider:
				// arr1[i-1] <= arr1[i] = j (non-decreasing)
				// arr2[i-1] >= arr2[i] → prev - arr1[i-1] >= curr - j
				// → arr1[i-1] <= prev - curr + j
				// So maxPrev = min(j, prev - curr + j) = j (since prev-curr+j <= j when prev <= curr, or > j when prev > curr)
				// Actually maxPrev = min(j, j + prev - curr) = j + min(0, prev - curr)

				// For dp[i-1][k] where k is arr1[i-1], we need:
				// k <= j AND k <= j + prev - curr
				// If prev - curr >= 0: both conditions give k <= j, so maxPrev = j
				// If prev - curr < 0: k <= j + (prev-curr), so maxPrev = j + prev - curr
				// But also arr1[i-1] can't exceed prev (since arr2[i-1] >= 0)
				if prev-curr >= 0 {
					maxPrev = j
				} else {
					maxPrev = j + prev - curr
				}
				if maxPrev > prev {
					maxPrev = prev
				}
				if maxPrev > j {
					maxPrev = j
				}
				if maxPrev >= 0 {
					ndp[j] = prefix[maxPrev+1]
				}
			}
		}

		dp = ndp
	}

	var ans int64
	for j := 0; j <= nums[n-1]; j++ {
		ans = (ans + dp[j]) % MOD
	}
	return int(ans)
}
```
