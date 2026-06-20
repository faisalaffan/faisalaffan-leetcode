# 3670 — Maximum Product Of Two Integers With No Common Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func maximumProductOfTwoIntegersWithNoCommonBits(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n + B*2^B)  |  **Ruang:** O(2^B)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3670: Maximum Product of Two Integers With No Common Bits
// https://leetcode.com/problems/maximum-product-of-two-integers-with-no-common-bits/
// Difficulty: Medium
// Time: O(n + B*2^B) | Space: O(2^B)

import (
	"fmt"
	"math/bits"
)

func maximumProductOfTwoIntegersWithNoCommonBits(nums []int) int64 {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	maxBits := 32 - bits.LeadingZeros32(uint32(maxVal))
	if maxBits == 0 {
		maxBits = 1
	}
	size := 1 << maxBits

  // Alokasi slice
	dp := make([]int, size)
	for _, x := range nums {
		if dp[x] < x {
			dp[x] = x
		}
	}

	// SOS DP
	for b := 0; b < maxBits; b++ {
		half := 1 << b
		step := half << 1
		for base := 0; base < size; base += step {
			upper := base + half
			for m := 0; m < half; m++ {
				u := upper + m
				l := base + m
				if dp[u] < dp[l] {
					dp[u] = dp[l]
				}
			}
		}
	}

	var ans int64 = 0
	full := size - 1
	for _, x := range nums {
		complement := (^x) & full
		y := dp[complement]
		if y > 0 {
			prod := int64(x) * int64(y)
			if prod > ans {
				ans = prod
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{3, 5, 2}))
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{1, 2, 4, 8}))
	fmt.Println(maximumProductOfTwoIntegersWithNoCommonBits([]int{5, 10, 3, 6}))
}
```
