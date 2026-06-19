package main

// LeetCode #2569: Handling Sum Queries After Update
// https://leetcode.com/problems/handling-sum-queries-after-update/
// Difficulty: Hard

import "fmt"

// handleQuery processes 3 types of queries on nums1 and nums2:
//   1 l r: flip bits in nums1[l..r] (0->1, 1->0)
//   2 x 0: nums2[j] += nums1[j] * x for all j
//   3 0 0: append sum(nums2) to result
//
// Approach: Segment tree on nums1 tracking count of 1s per segment with lazy flip.
// Type 2: sum2 += x * countOnesInNums1 (since nums1[j] is 0 or 1).
// Type 3: append current sum2.
//
// Complexity: O((n+q)*log n) time, O(n) space

// segTree tracks count of 1s in nums1 segments with lazy range flips
type segTree struct {
	n    int
	ones []int  // count of 1s in each segment
	lazy []bool // pending flip flag
}

func newSegTree(nums []int) *segTree {
	n := len(nums)
	st := &segTree{
		n:    n,
		ones: make([]int, 4*n),
		lazy: make([]bool, 4*n),
	}
	st.build(nums, 0, 0, n-1)
	return st
}

func (st *segTree) build(nums []int, idx, l, r int) {
	if l == r {
		st.ones[idx] = nums[l]
		return
	}
	mid := (l + r) / 2
	st.build(nums, 2*idx+1, l, mid)
	st.build(nums, 2*idx+2, mid+1, r)
	st.ones[idx] = st.ones[2*idx+1] + st.ones[2*idx+2]
}

func (st *segTree) push(idx, l, r int) {
	if st.lazy[idx] {
		// Flip: new count of 1s = segment length - old count
		st.ones[idx] = (r - l + 1) - st.ones[idx]
		if l != r {
			st.lazy[2*idx+1] = !st.lazy[2*idx+1]
			st.lazy[2*idx+2] = !st.lazy[2*idx+2]
		}
		st.lazy[idx] = false
	}
}

func (st *segTree) update(idx, l, r, ql, qr int) {
	st.push(idx, l, r)
	if ql > r || qr < l {
		return
	}
	if ql <= l && r <= qr {
		st.lazy[idx] = !st.lazy[idx]
		st.push(idx, l, r)
		return
	}
	mid := (l + r) / 2
	st.update(2*idx+1, l, mid, ql, qr)
	st.update(2*idx+2, mid+1, r, ql, qr)
	st.ones[idx] = st.ones[2*idx+1] + st.ones[2*idx+2]
}

func (st *segTree) flipRange(l, r int) {
	st.update(0, 0, st.n-1, l, r)
}

func (st *segTree) totalOnes() int {
	st.push(0, 0, st.n-1)
	return st.ones[0]
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	st := newSegTree(nums1)
	sum2 := int64(0)
	for _, v := range nums2 {
		sum2 += int64(v)
	}

	var result []int64
	for _, q := range queries {
		switch q[0] {
		case 1:
			st.flipRange(q[1], q[2])
		case 2:
			x := int64(q[1])
			ones := int64(st.totalOnes())
			sum2 += ones * x
		case 3:
			result = append(result, sum2)
		}
	}
	return result
}

func main() {
	// Test cases
	nums1 := []int{1, 0, 1}
	nums2 := []int{0, 0, 0}
	queries := [][]int{{1, 0, 1}, {2, 1, 0}, {3, 0, 0}, {2, 2, 0}, {3, 0, 0}}
	fmt.Println("Test 1: ->", handleQuery(nums1, nums2, queries))

	nums1 = []int{1, 1, 1}
	nums2 = []int{0, 0, 0}
	queries = [][]int{{2, 2, 0}, {3, 0, 0}, {1, 0, 2}, {3, 0, 0}}
	fmt.Println("Test 2: ->", handleQuery(nums1, nums2, queries))

	// Edge cases
	fmt.Println("Test 3: single element ->", handleQuery([]int{0}, []int{5}, [][]int{{3, 0, 0}}))
	fmt.Println("Test 4: flip and add ->", handleQuery([]int{1}, []int{5}, [][]int{{2, 3, 0}, {3, 0, 0}}))
	fmt.Println("Test 5: multiple ops ->", handleQuery(
		[]int{1, 0, 1, 0, 1},
		[]int{1, 2, 3, 4, 5},
		[][]int{{1, 0, 4}, {2, 10, 0}, {3, 0, 0}},
	))
}
