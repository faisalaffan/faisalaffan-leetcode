# 2321 — Maximum Score Of Spliced Array

## Deskripsi

**Soal:** [2321. Maximum Score Of Spliced Array](https://leetcode.com/problems/maximum-score-of-spliced-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func maximumsSplicedArray(nums1, nums2 []int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// 2321. Maximum Score Of Spliced Array
// ----------------------------------------------------------------
// We may swap a single contiguous subarray between nums1 and nums2.
// After the swap, return the maximum possible sum of EITHER array.
//
// Gain from swapping subarray [l,r]:
//   nums2 → nums1: gain = sum(nums2[l..r] - nums1[l..r])
//   nums1 → nums2: gain = sum(nums1[l..r] - nums2[l..r])
//
// Answer = max(sum(nums1) + bestKadane(diff),
//              sum(nums2) + bestKadane(-diff))
// where diff[i] = nums2[i] - nums1[i].

func maximumsSplicedArray(nums1, nums2 []int) int {
	n := len(nums1)
	sum1, sum2 := 0, 0
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		diff[i] = nums2[i] - nums1[i]
	}

	g1 := maxSubarray(diff)   // best gain from swapping nums2 INTO nums1
	g2 := maxSubarrayNeg(diff) // best gain from swapping nums1 INTO nums2

	ans := sum1 + g1
	if sum2+g2 > ans {
		ans = sum2 + g2
	}
	return ans
}

// maxSubarray returns the maximum subarray sum (Kadane); empty subarray → 0.
func maxSubarray(arr []int) int {
	best, cur := 0, 0
	for _, x := range arr {
		cur += x
		if cur < 0 {
			cur = 0
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// maxSubarrayNeg returns maxSubarray of -arr.
func maxSubarrayNeg(arr []int) int {
	best, cur := 0, 0
	for _, x := range arr {
		cur -= x // equivalent to cur += (-x)
		if cur < 0 {
			cur = 0
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

// ---------------------------------------------------------------------------
//  Wrapper

func MaximumScoreOfSplicedArray() interface{} {
	return maximumsSplicedArray([]int{60, 60, 60}, []int{10, 90, 10})
}

func main() {
	fmt.Println(MaximumScoreOfSplicedArray())

	tests := []struct {
		a, b []int
		want int
	}{
		{[]int{60, 60, 60}, []int{10, 90, 10}, 210},
		{[]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}, 220},
		{[]int{7, 11, 13}, []int{1, 1, 1}, 31},
		{[]int{1, 2, 3}, []int{3, 2, 1}, 8},
	}
	for _, tc := range tests {
		got := maximumsSplicedArray(tc.a, tc.b)
		if got != tc.want {
			fmt.Printf("FAIL nums1=%v nums2=%v: got %d, want %d\n",
				tc.a, tc.b, got, tc.want)
		}
	}
	fmt.Println("Done testing 2321.")
}
```
