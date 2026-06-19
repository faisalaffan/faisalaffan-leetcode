package main

// LeetCode #2818: Apply Operations to Maximize Score
// https://leetcode.com/problems/apply-operations-to-maximize-score/
// Difficulty: Hard
//
// Monotonic stack to find subarray dominance + prime score + greedy. For each
// element compute its prime score (distinct prime factors), then find how many
// subarrays it dominates. Sort by value descending and apply k operations.
// O(N * sqrt(maxVal) + N log N) time, O(N) space.

import (
	"fmt"
	"sort"
)

const mod2818 = 1000000007

func primeScore(n int) int {
	count := 0
	remaining := n
	for p := 2; p*p <= remaining; p++ {
		if remaining%p == 0 {
			count++
			for remaining%p == 0 {
				remaining /= p
			}
		}
	}
	if remaining > 1 {
		count++
	}
	return count
}

func powMod(a, e int64) int64 {
	res := int64(1)
	a %= mod2818
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % mod2818
		}
		a = (a * a) % mod2818
		e >>= 1
	}
	return res
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	scores := make([]int, n)
	for i, v := range nums {
		scores[i] = primeScore(v)
	}

	// Previous greater (or equal) element index
	prev := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && scores[stack[len(stack)-1]] < scores[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			prev[i] = -1
		} else {
			prev[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next greater (strictly greater) element index
	next := make([]int, n)
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && scores[stack[len(stack)-1]] <= scores[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			next[i] = n
		} else {
			next[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Sort indices by value descending (if tie, by index ascending)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		if nums[indices[i]] != nums[indices[j]] {
			return nums[indices[i]] > nums[indices[j]]
		}
		return indices[i] < indices[j]
	})

	result := int64(1)
	for _, idx := range indices {
		leftCount := idx - prev[idx]
		rightCount := next[idx] - idx
		applications := leftCount * rightCount
		use := applications
		if use > k {
			use = k
		}
		if use > 0 {
			result = (result * powMod(int64(nums[idx]), int64(use))) % mod2818
			k -= use
		}
		if k == 0 {
			break
		}
	}

	return int(result)
}

func main() {
	// Example: nums=[8,3,9,3,8], k=2 => 81
	fmt.Println(maximumScore([]int{8, 3, 9, 3, 8}, 2))
	// Single element
	fmt.Println(maximumScore([]int{5}, 1))
	// K larger than total subarrays
	fmt.Println(maximumScore([]int{2, 3}, 3))
	// All same values
	fmt.Println(maximumScore([]int{4, 4, 4}, 2))
	// Prime-heavy
	fmt.Println(maximumScore([]int{19, 12, 14, 6, 10}, 3))
}
