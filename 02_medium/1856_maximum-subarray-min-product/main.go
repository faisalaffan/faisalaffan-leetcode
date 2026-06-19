package main

// LeetCode #1856: Maximum Subarray Min-Product
// https://leetcode.com/problems/maximum-subarray-min-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxSumMinProduct([]int{1, 2, 3, 2}))
	fmt.Println(MaxSumMinProduct([]int{2, 3, 3, 1, 2}))
	fmt.Println(MaxSumMinProduct([]int{3, 1, 5, 6, 4, 2}))
}

const mod = 1000000007

// Time: O(n), Space: O(n)
func MaxSumMinProduct(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Monotonic stack to find previous smaller and next smaller elements
	left := make([]int, n)  // left[i] = index of previous smaller element
	right := make([]int, n) // right[i] = index of next smaller element

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxProd := int64(0)
	for i := 0; i < n; i++ {
		sum := int64(prefix[right[i]] - prefix[left[i]+1])
		prod := sum * int64(nums[i])
		if prod > maxProd {
			maxProd = prod
		}
	}
	return int(maxProd % mod)
}
