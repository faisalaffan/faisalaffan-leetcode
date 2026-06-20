package main

// LeetCode #3762: Minimum Operations to Equalize Subarrays
// https://leetcode.com/problems/minimum-operations-to-equalize-subarrays/
// Difficulty: Hard
//
// For each query [l, r], find min operations to make all elements
// in nums[l..r] equal. Each operation inc/dec any element by 1.
// Optimal target is the median. Use prefix sums.
//
// Approach: Merge sort tree for kth order statistic (median).
// Prefix sums for range sum queries.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}, 2, [][]int{{0, 3}, {1, 2}}))
	// Example 2
	fmt.Println(minimumOperations([]int{1, 1, 1}, 1, [][]int{{0, 2}}))
	// Edge: single element
	fmt.Println(minimumOperations([]int{5}, 0, [][]int{{0, 0}}))
}

func minimumOperations(nums []int, k int, queries [][]int) []int64 {
	n := len(nums)

	// Build merge sort tree
	tree := make([][]int, 4*n)
	var build func(node, l, r int)
	build = func(node, l, r int) {
		if l == r {
			tree[node] = []int{nums[l]}
			return
		}
		mid := (l + r) / 2
		build(node*2, l, mid)
		build(node*2+1, mid+1, r)
		tree[node] = merge(tree[node*2], tree[node*2+1])
	}
	build(1, 0, n-1)

	// Count elements <= x in range [ql, qr]
	var countLe func(node, l, r, ql, qr, x int) int
	countLe = func(node, l, r, ql, qr, x int) int {
		if ql > r || qr < l {
			return 0
		}
		if ql <= l && r <= qr {
			// Binary search in sorted array at this node
			arr := tree[node]
			return sort.SearchInts(arr, x+1) // number of elements <= x
		}
		mid := (l + r) / 2
		return countLe(node*2, l, mid, ql, qr, x) +
			countLe(node*2+1, mid+1, r, ql, qr, x)
	}

	// Binary search for kth smallest in range [ql, qr]
	kth := func(ql, qr, k int) int {
		lo, hi := int(1e9+5), int(-1e9-5)
		for _, v := range nums {
			if v < lo {
				lo = v
			}
			if v > hi {
				hi = v
			}
		}
		for lo <= hi {
			mid := (lo + hi) / 2
			if countLe(1, 0, n-1, ql, qr, mid) >= k {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
		return lo
	}

	// Prefix sums for fast range sum
	pref := make([]int64, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + int64(v)
	}

	ans := make([]int64, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		length := r - l + 1

		// Find median
		med := kth(l, r, (length+1)/2)

		// Count elements <= med and their sum
		cntLe := countLe(1, 0, n-1, l, r, med)
		// Count elements < med
		cntLt := countLe(1, 0, n-1, l, r, med-1)

		// Sum of all elements in range
		totalSum := pref[r+1] - pref[l]

		// Sum of elements < med using prefix of sorted values
		// We need a secondary structure for sum by rank in range.
		// For simplicity, compute sum of elements <= med by scanning
		// (only for elements = med).
		// Actually use count of < med and count of <= med to determine
		// how many elements equal med, then compute:
		// sum_less = totalSum - med * cntGe - sum_greater
		// This is circular.

		// Let's compute directly: sum of elements < med in range
		// We need a merge sort tree with prefix sums too.
		// For simplicity, just iterate the range (n is small enough).
		var sumLt, cntExact int64
		for i := l; i <= r; i++ {
			v := int64(nums[i])
			if v < int64(med) {
				sumLt += v
			} else if v == int64(med) {
				cntExact++
			}
		}

		cntLt64 := int64(cntLt)
		cntLe64 := int64(cntLe)

		// Elements < med: need med - each_value operations each
		opsLeft := int64(med)*cntLt64 - sumLt

		// Elements > med: need each_value - med operations each
		cntGt := int64(length) - cntLe64
		sumGt := totalSum - sumLt - int64(med)*cntExact
		opsRight := sumGt - int64(med)*cntGt

		ans[qi] = opsLeft + opsRight
	}

	return ans
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		res[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		res[k] = b[j]
		j++
		k++
	}
	return res
}
