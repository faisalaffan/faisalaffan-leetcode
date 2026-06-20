# 3132 — Find The Integer Added To Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minimumAddedInteger(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3132: Find the Integer Added to Array II
// https://leetcode.com/problems/find-the-integer-added-to-array-ii/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minimumAddedInteger(nums1 []int, nums2 []int) int {
  // Sort O(n log n)
	sort.Ints(nums1)
  // Sort O(n log n)
	sort.Ints(nums2)

	// Try all pairs from nums1 as the two removed elements
  // Linear scan O(n)
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			// Check if nums2 can be matched after removing nums1[i] and nums1[j]
			diff := -1001
			idx := 0
			match := true
			for k := 0; k < len(nums1) && match; k++ {
				if k == i || k == j {
					continue
				}
				curDiff := nums2[idx] - nums1[k]
				if diff == -1001 {
					diff = curDiff
				} else if curDiff != diff {
					match = false
				}
				idx++
			}
			if match && diff >= 0 {
				return diff
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10})) // Expected: -2
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}))             // Expected: 2
}
```
