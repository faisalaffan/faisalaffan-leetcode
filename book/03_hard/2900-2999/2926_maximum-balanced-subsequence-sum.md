# 2926 — Maximum Balanced Subsequence Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxBalancedSubsequenceSum(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2926: Maximum Balanced Subsequence Sum
// https://leetcode.com/problems/maximum-balanced-subsequence-sum/
// Difficulty: Hard
//
// A subsequence nums[i1], nums[i2], ..., nums[ik] is balanced if
// nums[i_{t+1}] - nums[i_t] >= i_{t+1} - i_t, which is equivalent to
// nums[i] - i being non-decreasing. Transform each element to key = nums[i] - i,
// then find the subsequence with non-decreasing keys maximizing sum of nums[i].
// Use BIT (Fenwick tree) with coordinate compression for DP: For each element,
// query max sum for keys <= current key, add nums[i], update BIT.
// O(N log N) time, O(N) space.

import (
	"fmt"
	"math"
	"sort"
)

func maxBalancedSubsequenceSum(nums []int) int64 {
	n := len(nums)
  // Alokasi slice integer
	keys := make([]int, n)
	for i, v := range nums {
		keys[i] = v - i
	}

	// Coordinate compression
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, keys)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)
	m := 1
	for i := 1; i < n; i++ {
		if sorted[i] != sorted[m-1] {
			sorted[m] = sorted[i]
			m++
		}
	}
	sorted = sorted[:m]

	// BIT for prefix maximum
  // Alokasi slice integer
	bit := make([]int64, m+2)
  // Range loop: iterasi dengan indeks + nilai
	for i := range bit {
		bit[i] = math.MinInt64
	}

	query := func(pos int) int64 {
		res := int64(math.MinInt64)
		for pos > 0 {
			if bit[pos] > res {
				res = bit[pos]
			}
			pos -= pos & -pos
		}
		return res
	}
	update := func(pos int, val int64) {
		for pos <= m {
			if val > bit[pos] {
				bit[pos] = val
			}
			pos += pos & -pos
		}
	}

	var ans int64 = int64(nums[0])
	for i, v := range nums {
		pos := sort.SearchInts(sorted, keys[i]) + 1 // 1-indexed BIT
		best := query(pos)
		if best == math.MinInt64 {
			best = 0
		}
		cur := best + int64(v)
		if cur > ans {
			ans = cur
		}
		update(pos, cur)
	}
	return ans
}

func main() {
	// Example: [3,3,5,6] => 14 (subsequence [3,5,6])
	// keys: [3,2,3,3]
	// 3 (pos 3): query(<=3)=minInt => cur=3
	// 3 (pos 2): query(<=2)=minInt => cur=3
	// 5 (pos 3): query(<=3)=max(3,3)=3 => cur=3+5=8
	// 6 (pos 3): query(<=3)=max(3,3,8)=8 => cur=8+6=14
	fmt.Println(maxBalancedSubsequenceSum([]int{3, 3, 5, 6}))

	// All equal: [5,5,5] => 15
	fmt.Println(maxBalancedSubsequenceSum([]int{5, 5, 5}))

	// All negative: [-1,-2,-3] => -1 (pick single max)
	fmt.Println(maxBalancedSubsequenceSum([]int{-1, -2, -3}))

	// Mixed
	fmt.Println(maxBalancedSubsequenceSum([]int{5, -10, 3}))

	// Increasing nums
	fmt.Println(maxBalancedSubsequenceSum([]int{1, 2, 3, 4, 5}))

	// Decreasing nums
	fmt.Println(maxBalancedSubsequenceSum([]int{5, 4, 3, 2, 1}))

	// Single element
	fmt.Println(maxBalancedSubsequenceSum([]int{-5}))

	// Complex case
	fmt.Println(maxBalancedSubsequenceSum([]int{10, 1, 2, 3, 4, 5}))
}
```
