package main

// LeetCode #3187: Peaks in Array
// https://leetcode.com/problems/peaks-in-array/
// Difficulty: Hard
//
// Given an array nums and queries of two types:
//   1 [l, r] -> count peaks in subarray nums[l..r] (endpoints excluded)
//   2 [idx, val] -> set nums[idx] = val
// A peak is nums[i] > nums[i-1] && nums[i] > nums[i+1].
// Use a Binary Indexed Tree to track peak status per index.
// Updates only affect positions idx-1, idx, idx+1.

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	queries := [][]int{
		{2, 0, 5},
		{2, 1, 3},
		{2, 2, 6},
	}
	fmt.Println(countOfPeaks(nums, queries))
}

type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{tree: make([]int, n+2), n: n}
}

func (b *BIT) Add(i, delta int) {
	if i < 0 || i >= b.n {
		return
	}
	for idx := i + 1; idx <= b.n+1; idx += idx & -idx {
		b.tree[idx] += delta
	}
}

func (b *BIT) Sum(i int) int {
	if i < 0 {
		return 0
	}
	if i >= b.n {
		i = b.n - 1
	}
	res := 0
	for idx := i + 1; idx > 0; idx -= idx & -idx {
		res += b.tree[idx]
	}
	return res
}

func (b *BIT) RangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return b.Sum(r) - b.Sum(l-1)
}

func isPeak(nums []int, i int) bool {
	if i <= 0 || i >= len(nums)-1 {
		return false
	}
	return nums[i] > nums[i-1] && nums[i] > nums[i+1]
}

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	bit := NewBIT(n)

	// Initialize BIT with current peaks
	for i := 1; i < n-1; i++ {
		if isPeak(nums, i) {
			bit.Add(i, 1)
		}
	}

	ans := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 {
			// Count peaks in [l, r]
			l, r := q[1], q[2]
			// Endpoints cannot be peaks, so check [l+1, r-1]
			if r-l < 2 {
				ans = append(ans, 0)
			} else {
				ans = append(ans, bit.RangeSum(l+1, r-1))
			}
		} else {
			// Update: set nums[idx] = val
			idx, val := q[1], q[2]
			// Check old status for affected positions
			affected := []int{idx - 1, idx, idx + 1}
			oldStatus := make([]bool, 3)
			for p, pos := range affected {
				oldStatus[p] = isPeak(nums, pos)
			}

			nums[idx] = val

			// Check new status
			for p, pos := range affected {
				newStatus := isPeak(nums, pos)
				if oldStatus[p] != newStatus {
					if newStatus {
						bit.Add(pos, 1)
					} else {
						bit.Add(pos, -1)
					}
				}
			}
		}
	}
	return ans
}
