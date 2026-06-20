# 1983 — Widest Pair Of Indices With Equal Range Sum

## Deskripsi

**Soal:** [1983. Widest Pair Of Indices With Equal Range Sum](https://leetcode.com/problems/widest-pair-of-indices-with-equal-range-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1983: Widest Pair of Indices With Equal Range Sum
// https://leetcode.com/problems/widest-pair-of-indices-with-equal-range-sum/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 2, 3}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{1, 1, 1}, []int{1, 1, 1}))
	fmt.Println(WidestPairOfIndicesWithEqualRangeSum([]int{0, 1}, []int{1, 0}))
}

// Time: O(n), Space: O(n)
func WidestPairOfIndicesWithEqualRangeSum(nums1 []int, nums2 []int) int {
  // Membuat map untuk pencarian O(1): key → value
	first := make(map[int]int)
	first[0] = -1
	maxWidth := 0
	prefix1, prefix2 := 0, 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums1); i++ {
		prefix1 += nums1[i]
		prefix2 += nums2[i]
		diff := prefix1 - prefix2

		if idx, ok := first[diff]; ok {
			if i-idx > maxWidth {
				maxWidth = i - idx
			}
		} else {
			first[diff] = i
		}
	}

	return maxWidth
}
```
