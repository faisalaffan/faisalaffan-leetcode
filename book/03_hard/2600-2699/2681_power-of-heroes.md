# 2681 — Power Of Heroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sumOfPower(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2681: Power of Heroes
// https://leetcode.com/problems/power-of-heroes/
// Difficulty: Hard
//
// The power of any non-empty subset of heroes is max(nums)^2 * min(nums).
// Sum over all non-empty subsets. Return modulo 1_000_000_007.
// Sort ascending, then for each element as max, compute contribution
// using a running sum of min * 2^(distance) for earlier elements.
// Time: O(n log n), Space: O(1)

import (
	"fmt"
	"sort"
)

const mod = 1_000_000_007

func main() {
	// Example 1: [2,1,4] -> 141
	fmt.Println(sumOfPower([]int{2, 1, 4}))
	// Example 2: [1,1,1] -> 7
	fmt.Println(sumOfPower([]int{1, 1, 1}))
	// Example 3: [5] -> 125
	fmt.Println(sumOfPower([]int{5}))
}

func sumOfPower(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)

	ans := int64(0)
	// sumMin tracks sum of (min * 2^(distance to current)) for previous elements
	sumMin := int64(0)

	for _, v := range nums {
		vv := int64(v)

		// Contribution with v as max:
		// 1. Single-element subset: v^2 * v = v^3
		// 2. Multi-element: v^2 * sumMin (which has 2^dist factor built in)
		ans = (ans + (vv*vv%mod)*vv%mod) % mod                 // single element
		ans = (ans + (vv*vv%mod)*sumMin%mod) % mod              // multi-element

		// Update sumMin for next iteration:
		// sumMin_new = v + 2 * sumMin_old
		// Because all previous elements' weights double (one more position between min and max)
		// and v itself becomes a candidate min for future maxes
		sumMin = (sumMin*2 + vv) % mod
	}

	return int(ans)
}
```
