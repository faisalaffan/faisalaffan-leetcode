# 3935 — Power Update After K Th Largest Insertion I

## Deskripsi

**Soal:** [3935. Power Update After K Th Largest Insertion I](https://leetcode.com/problems/power-update-after-k-th-largest-insertion-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log M + Q log M)  
**Kompleksitas Ruang:** O(M) where M = max value

**Algoritma:** Binary Search (pencarian biner), Prefix Sum (jumlah kumulatif), Fenwick Tree (Binary Indexed Tree)

**Fungsi Solusi:** `func NewBIT(size int) *BIT`

> **Ide Kunci:** Maintain sorted multiset via Fenwick tree. For each query,

## Solusi Go

```go
package main

// LeetCode #3935: Power Update After K-th Largest Insertion I
// https://leetcode.com/problems/power-update-after-k-th-largest-insertion-i/
// Difficulty: Medium [Paid]
// Time: O(N log M + Q log M) | Space: O(M) where M = max value
// Approach: Maintain sorted multiset via Fenwick tree. For each query,
// insert val, find k-th largest via binary search on BIT prefix sums,
// update p = p ^ kth % MOD.

import (
	"fmt"
	"sort"
)

const MOD = 1000000007

type BIT struct {
	tree []int
	size int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2), size: size}
}

func (b *BIT) Update(idx int, delta int) {
	idx++
	for idx <= b.size+1 {
		b.tree[idx] += delta
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func (b *BIT) KthLargest(k int) int {
	// Find smallest idx such that suffix sum (total - prefix) >= k
	// i.e., total - Query(idx-1) >= k
	// i.e., Query(idx-1) <= total - k
	total := b.Query(b.size - 1)
	target := total - k // number of elements strictly less than the kth largest

	// Binary search for first index with prefix sum > target
	lo, hi := 0, b.size-1
	for lo < hi {
		mid := (lo + hi) / 2
		if b.Query(mid) > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func PowerUpdateAfterKThLargestInsertionI(nums []int, p int, queries [][]int) []int {
	// Coordinate compress all values
  // Membuat slice untuk menyimpan hasil
	allVals := make([]int, 0, len(nums)+len(queries))
	allVals = append(allVals, nums...)
	for _, q := range queries {
		allVals = append(allVals, q[0])
	}
	sort.Ints(allVals)
	uniq := []int{allVals[0]}
	for i := 1; i < len(allVals); i++ {
		if allVals[i] != allVals[i-1] {
			uniq = append(uniq, allVals[i])
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	rank := make(map[int]int)
	for i, v := range uniq {
		rank[v] = i
	}
	m := len(uniq)

	bit := NewBIT(m)
	for _, v := range nums {
		bit.Update(rank[v], 1)
	}

	cur := int64(p)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))

	for idx, q := range queries {
		val, k := q[0], q[1]
		bit.Update(rank[val], 1)

		kthVal := uniq[bit.KthLargest(k)]

		cur = (cur ^ int64(kthVal)) % MOD
		ans[idx] = int(cur)
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{2}, 4, [][]int{{3, 1}, {1, 2}},
	)) // Expected: [64, 4096]

	// Example 2
	fmt.Println(PowerUpdateAfterKThLargestInsertionI(
		[]int{7, 5}, 6, [][]int{{4, 3}, {7, 2}},
	)) // Expected: [1296, 220296870]
}
```
