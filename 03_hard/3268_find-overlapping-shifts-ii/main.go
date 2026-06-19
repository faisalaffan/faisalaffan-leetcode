package main

// LeetCode #3268: Find Overlapping Shifts II
// https://leetcode.com/problems/find-overlapping-shifts-ii/
// Difficulty: Hard [Paid]
//
// Given a list of shifts (intervals) and queries [l, r],
// for each query return the number of overlapping shift pairs
// entirely within the subarray shifts[l..r] (inclusive).
// Two shifts [a, b] and [c, d] overlap if they share any common time,
// i.e., a <= d && c <= b.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	shifts := [][]int{{1, 3}, {2, 5}, {6, 8}, {4, 7}}
	queries := [][]int{{0, 1}, {1, 3}, {0, 3}}
	fmt.Println(findOverlappingShiftsII(shifts, queries))
	// Example 2
	shifts2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	queries2 := [][]int{{0, 2}}
	fmt.Println(findOverlappingShiftsII(shifts2, queries2))
	// Example 3
	shifts3 := [][]int{{1, 10}, {2, 5}, {6, 9}, {3, 7}}
	queries3 := [][]int{{0, 1}, {0, 3}}
	fmt.Println(findOverlappingShiftsII(shifts3, queries3))
}

func findOverlappingShiftsII(shifts [][]int, queries [][]int) []int {
	n := len(shifts)

	// Precompute overlap for every pair of shifts.
	overlap := make([][]bool, n)
	for i := range overlap {
		overlap[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := shifts[i][0], shifts[i][1]
			c, d := shifts[j][0], shifts[j][1]
			if a <= d && c <= b {
				overlap[i][j] = true
				overlap[j][i] = true
			}
		}
	}

	// For each query [l, r], count overlapping pairs within [l, r].
	// Precompute prefix sums of overlap counts to answer queries in O(1).
	// pref[i][j] = number of overlapping pairs with first index < i and second index < j.
	pref := make([][]int, n+1)
	for i := range pref {
		pref[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := 0
			if overlap[i][j] {
				val = 1
			}
			pref[i+1][j+1] = pref[i][j+1] + pref[i+1][j] - pref[i][j] + val
		}
	}

	// Count pairs (i,j) with l <= i < j <= r.
	countInRange := func(l, r int) int {
		total := pref[r+1][r+1] - pref[l][r+1] - pref[r+1][l] + pref[l][l]
		// Remove the diagonal (i,i) which pref includes.
		diag := r - l + 1
		// Each pair (i,j) for i<j appears once in the pref sum.
		// The pref sum also includes (j,i) entries but those are counted separately.
		// Since our overlap matrix is symmetric, pref counts each pair twice.
		return total / 2
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		ans[qi] = countInRange(q[0], q[1])
	}
	return ans
}

// Alternative O(n log n) approach using Fenwick tree for
// online/streaming queries sorted by right endpoint.
// For simplicity we use the precomputation approach above
// which works well for n up to ~2000.

// Query-friendly approach (for larger n + many queries):
type event struct {
	l, r, idx int
}

func findOverlappingShiftsIIFenwick(shifts [][]int, queries [][]int) []int {
	n := len(shifts)
	m := len(queries)

	// Sort queries by right endpoint.
	qidx := make([]int, m)
	for i := range qidx {
		qidx[i] = i
	}
	sort.Slice(qidx, func(i, j int) bool {
		return queries[qidx[i]][1] < queries[qidx[j]][1]
	})

	// Group shifts by right endpoint.
	byRight := make([][]int, n)
	for i, s := range shifts {
		byRight[s[1]] = append(byRight[s[1]], s[0])
	}

	// Fenwick tree over left endpoints.
	tree := make([]int, n+2)
	add := func(pos, val int) {
		for pos <= n {
			tree[pos] += val
			pos += pos & -pos
		}
	}
	sum := func(pos int) int {
		s := 0
		for pos > 0 {
			s += tree[pos]
			pos -= pos & -pos
		}
		return s
	}
	rangeSum := func(l, r int) int {
		return sum(r) - sum(l-1)
	}

	ans := make([]int, m)
	shiftPtr := 0
	sortedShifts := make([]struct{ l, r int }, n)
	for i, s := range shifts {
		sortedShifts[i] = struct{ l, r int }{s[0], s[1]}
	}
	sort.Slice(sortedShifts, func(i, j int) bool {
		return sortedShifts[i].r < sortedShifts[j].r
	})

	for _, qi := range qidx {
		l, r := queries[qi][0], queries[qi][1]

		// Add shifts whose right endpoint <= r.
		for shiftPtr < n && sortedShifts[shiftPtr].r <= r {
			add(sortedShifts[shiftPtr].l, 1)
			shiftPtr++
		}

		// Count shifts with left endpoint >= l and right <= r.
		cnt := rangeSum(l, n)
		// Overlapping pairs = C(cnt, 2)
		ans[qi] = cnt * (cnt - 1) / 2
	}

	return ans
}
