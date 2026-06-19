package main

// LeetCode #2819: Minimum Relative Loss After Buying Chocolates
// https://leetcode.com/problems/minimum-relative-loss-after-buying-chocolates/
// Difficulty: Hard [Paid]
//
// For each query (k, m), Bob picks exactly m chocolates. For price <= k, Bob pays
// full price (contribution = p). For price > k, Bob pays k and Alice pays rest
// (contribution = 2k - p). Minimize sum of contributions.
// Sort prices, precompute prefix sums. For each query, find optimal split point:
// x cheapest from left (price <= k) and m-x most expensive from right (price > k).
// Loss function is convex -- use ternary search on x.
// O(N log N + Q log N) time, O(N) space.

import (
	"fmt"
	"sort"
)

func minRelativeLoss(prices []int, queries [][]int) []int64 {
	sort.Ints(prices)
	n := len(prices)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(prices[i])
	}

	result := make([]int64, len(queries))
	for qi, q := range queries {
		k, m := q[0], q[1]

		// split = first index where price > k
		split := sort.Search(n, func(i int) bool { return prices[i] > k })

		// x = number of chocolates picked from left (prices <= k)
		// Must pick exactly m chocolates total.
		// x ranges from max(0, m - (n-split)) to min(split, m)
		minX := m - (n - split)
		if minX < 0 {
			minX = 0
		}
		maxX := split
		if m < maxX {
			maxX = m
		}

		// loss(x) = prefix[x] + 2k*(m-x) - (prefix[n] - prefix[n-m+x])
		// The derivative loss(x+1)-loss(x) = prices[x] + prices[n-m+x] - 2k
		// is monotonic (convex). Binary search for minimum.
		lo, hi := minX, maxX
		for lo < hi {
			mid := (lo + hi) / 2
			// At mid, what is the slope?
			diff := prices[mid] + prices[n-m+mid] - 2*k
			if diff >= 0 {
				hi = mid
			} else {
				lo = mid + 1
			}
		}

		// Check lo and lo-1
		best := loss(prefix, prices, n, m, k, lo)
		if lo-1 >= minX {
			cand := loss(prefix, prices, n, m, k, lo-1)
			if cand < best {
				best = cand
			}
		}

		result[qi] = best
	}
	return result
}

func loss(prefix []int64, prices []int, n, m, k, x int) int64 {
	leftSum := prefix[x]
	rightNeed := m - x
	if rightNeed < 0 {
		return 1 << 62
	}
	rightSum := prefix[n] - prefix[n-rightNeed]
	return leftSum + int64(rightNeed)*int64(2*k) - rightSum
}

func main() {
	// Example 1
	res1 := minRelativeLoss([]int{1, 9, 22, 10, 19}, [][]int{{18, 4}, {5, 2}})
	for _, v := range res1 {
		fmt.Println(v)
	}

	// Example 2
	res2 := minRelativeLoss([]int{1, 5, 4, 3, 7, 11, 9}, [][]int{{5, 4}, {5, 7}, {7, 3}, {4, 5}})
	for _, v := range res2 {
		fmt.Println(v)
	}

	// Example 3
	res3 := minRelativeLoss([]int{5, 6, 7}, [][]int{{10, 1}, {5, 3}, {3, 3}})
	for _, v := range res3 {
		fmt.Println(v)
	}

	// Single chocolate, various queries
	fmt.Println(minRelativeLoss([]int{10}, [][]int{{5, 1}, {20, 1}}))

	// All prices <= k (should pick any m chocolates)
	fmt.Println(minRelativeLoss([]int{1, 2, 3, 4, 5}, [][]int{{10, 3}}))

	// All prices > k
	fmt.Println(minRelativeLoss([]int{10, 20, 30}, [][]int{{5, 2}}))

	// m = n (must pick all)
	fmt.Println(minRelativeLoss([]int{2, 3, 5}, [][]int{{4, 3}}))
}
