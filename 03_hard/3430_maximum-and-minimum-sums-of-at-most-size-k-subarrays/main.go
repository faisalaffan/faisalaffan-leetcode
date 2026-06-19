package main

// LeetCode #3430: Maximum and Minimum Sums of at Most Size K Subarrays
// https://leetcode.com/problems/maximum-and-minimum-sums-of-at-most-size-k-subarrays/
// Difficulty: Hard
//
// Monotonic stack for previous/next smaller/greater.
// Contribution counting with k constraint.

import "fmt"

func main() {
	fmt.Println(MaximumAndMinimumSumsOfAtMostSizeKSubarrays([]int{1, 3, 2}, 2))
}

func MaximumAndMinimumSumsOfAtMostSizeKSubarrays(nums []int, k int) int64 {
	n := len(nums)

	// Previous smaller element (strict)
	ps := make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ps[i] = stack[len(stack)-1]
		} else {
			ps[i] = -1
		}
		stack = append(stack, i)
	}

	// Next smaller element (strict)
	ns := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ns[i] = stack[len(stack)-1]
		} else {
			ns[i] = n
		}
		stack = append(stack, i)
	}

	// Previous greater element (strict)
	pg := make([]int, n)
	stack = stack[:0]
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			pg[i] = stack[len(stack)-1]
		} else {
			pg[i] = -1
		}
		stack = append(stack, i)
	}

	// Next greater element (strict)
	ng := make([]int, n)
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ng[i] = stack[len(stack)-1]
		} else {
			ng[i] = n
		}
		stack = append(stack, i)
	}

	contrib := func(left, right, idx int) int64 {
		l := int64(idx - left)
		r := int64(right - idx)
		total := int64(0)
		// Sum of subarray counts with length <= k that include idx
		// For each possible subarray length L: min(L, k) counts of being in a length-L subarray
		// Actually count # subarrays where idx is min/max and subarray length <= k
		// Left choices: 1..l, right choices: 1..r
		// For each leftLen a and rightLen b, subarray length = a+b-1
		// Count where a+b-1 <= k, i.e., a+b <= k+1

		// For each a in [1, l], b in [1, r], where a+b <= k+1
		// cnt = sum_{a=1}^{l} min(r, k+1-a)
		// Clamp a such that k+1-a >= 1 => a <= k
		maxA := l
		if maxA > int64(k) {
			maxA = int64(k)
		}
		for a := int64(1); a <= maxA; a++ {
			rem := int64(k+1) - a
			bLimit := r
			if rem < bLimit {
				bLimit = rem
			}
			if bLimit >= 1 {
				total += bLimit
			}
		}
		return total
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		minContrib := contrib(ps[i], ns[i], i)
		maxContrib := contrib(pg[i], ng[i], i)
		ans += int64(nums[i]) * (maxContrib - minContrib)
	}
	return ans
}
