# 2426 — Number Of Pairs Satisfying Inequality

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func newBIT(size int) *BIT`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting, Fenwick Tree

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2426: Number of Pairs Satisfying Inequality
// https://leetcode.com/problems/number-of-pairs-satisfying-inequality/
// Difficulty: Hard
//
// Given nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff for i < j.
// Rearranged: (nums1[i] - nums2[i]) <= (nums1[j] - nums2[j]) + diff.
// Let arr[k] = nums1[k] - nums2[k]. Then for i < j: arr[i] <= arr[j] + diff.
// Processing left to right, at position j count previous i where
// arr[i] <= arr[j] + diff. Use BIT (Fenwick Tree) on compressed values.

import (
	"fmt"
	"sort"
)

type BIT struct {
	tree []int
}

func newBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) add(idx int) {
	for i := idx; i < len(b.tree); i += i & -i {
		b.tree[i]++
	}
}

func (b *BIT) sum(idx int) int {
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

func numberOfPairs(input [][]int, diff int) int64 {
	nums1, nums2 := input[0], input[1]
	n := len(nums1)

  // Alokasi slice
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nums1[i] - nums2[i]
	}

	// Coordinate compression
  // Alokasi slice
	allVals := make([]int, 0, n*2)
	for _, v := range arr {
		allVals = append(allVals, v, v+diff)
	}
  // Sort O(n log n)
	sort.Ints(allVals)
	uniq := 1
	for i := 1; i < len(allVals); i++ {
		if allVals[i] != allVals[uniq-1] {
			allVals[uniq] = allVals[i]
			uniq++
		}
	}
	allVals = allVals[:uniq]

	compress := func(x int) int {
		return sort.SearchInts(allVals, x) + 1
	}

	bt := newBIT(len(allVals) + 2)
	var ans int64

	for j := 0; j < n; j++ {
		// Count previous arr[i] where arr[i] <= arr[j] + diff
		target := arr[j] + diff
		pos := compress(target)
		ans += int64(bt.sum(pos))
		bt.add(compress(arr[j]))
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(numberOfPairs([][]int{{3, 2, 5}, {2, 2, 1}}, 1))
	// Example 2
	fmt.Println(numberOfPairs([][]int{{3, -1}, {-2, 2}}, -1))
	// Single element
	fmt.Println(numberOfPairs([][]int{{1}, {1}}, 0))
	// All equal
	fmt.Println(numberOfPairs([][]int{{1, 1, 1}, {1, 1, 1}}, 0))
	// Larger range
	fmt.Println(numberOfPairs([][]int{{1, 3, 5, 7}, {2, 4, 6, 8}}, 2))
}
```
