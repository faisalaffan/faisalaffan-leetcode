# 2143 — Choose Numbers From Two Arrays In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countSubarrays(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2143: Choose Numbers From Two Arrays in Range
// https://leetcode.com/problems/choose-numbers-from-two-arrays-in-range/
// Difficulty: Hard [Paid]
//
// Count subarrays [l, r] where sum(nums1[l..r]) == sum(nums2[l..r]) and the
// combined multiset has at most one odd value. Use prefix difference tracking:
// diff = sum1[i] - sum2[i], oddCount = number of odd values up to i.
// For each r, find matching l with same diff and oddCount diff <= 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}))

	// Example 2
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}, []int{3, 4, 2, 1, 5}))

	// Example 3
	fmt.Println(countSubarrays([]int{0, 0, 0}, []int{0, 0, 0}))

	// Simple case
	fmt.Println(countSubarrays([]int{1, 1}, []int{1, 1}))
}

func countSubarrays(nums1 []int, nums2 []int) int {
	n := len(nums1)

	// prefixDiff[i] = sum(nums1[0..i-1]) - sum(nums2[0..i-1])
  // Alokasi slice
	prefixDiff := make([]int, n+1)
	// prefixOdd[i] = count of odd values in nums1[0..i-1] + nums2[0..i-1]
  // Alokasi slice
	prefixOdd := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixDiff[i+1] = prefixDiff[i] + nums1[i] - nums2[i]
		oddCount := 0
		if nums1[i]%2 == 1 {
			oddCount++
		}
		if nums2[i]%2 == 1 {
			oddCount++
		}
		prefixOdd[i+1] = prefixOdd[i] + oddCount
	}

	// Map: prefixDiff -> list of prefixOdd values at that diff
	// We'll use a map from diff to a map of oddCount->frequency
  // HashMap: O(1) lookup
	diffMap := make(map[int]map[int]int)
	ans := 0

	// l starts at 0 (empty prefix), r runs from 1 to n
	// We iterate r, and before processing, we add prefix l=r to the map
	// For r, we want l where prefixDiff[r] == prefixDiff[l] and prefixOdd[r] - prefixOdd[l] <= 1

	// Initialize with l=0 (empty prefix)
	diffMap[0] = map[int]int{0: 1}

	for r := 1; r <= n; r++ {
		diff := prefixDiff[r]
		odd := prefixOdd[r]

		// Query: find l where prefixDiff[l] == diff and odd - prefixOdd[l] <= 1
		if m, ok := diffMap[diff]; ok {
			// Add all l where oddDiff <= 1
			// oddDiff = odd - prefixOdd[l]
			// prefixOdd[l] >= odd - 1 and prefixOdd[l] <= odd
			for oddL, freq := range m {
				if odd-oddL <= 1 {
					ans += freq
				}
			}
		}

		// Add current prefix to map
		if _, ok := diffMap[diff]; !ok {
			diffMap[diff] = make(map[int]int)
		}
		diffMap[diff][odd]++
	}

	return ans
}
```
