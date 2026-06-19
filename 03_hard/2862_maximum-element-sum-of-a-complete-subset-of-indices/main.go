package main

// LeetCode #2862: Maximum Element-Sum of a Complete Subset of Indices
// https://leetcode.com/problems/maximum-element-sum-of-a-complete-subset-of-indices/
// Difficulty: Hard
//
// A subset S of indices {1..n} is "complete" if for any i, j in S, i*j <= n implies i*j in S.
// This is equivalent to: all indices with the same squarefree kernel form a maximal complete
// subset. The squarefree kernel of an index is its value with all square factors removed.
// Group indices by squarefree kernel, sum their nums values, return the max sum.

import "fmt"

func maximumElementSumOfCompleteSubsetOfIndices(nums []int) int64 {
	n := len(nums)

	// Compute smallest prime factor (spf) for numbers up to n
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if spf[i] == 0 {
			spf[i] = i
			if i*i <= n {
				for j := i * i; j <= n; j += i {
					if spf[j] == 0 {
						spf[j] = i
					}
				}
			}
		}
	}

	// Compute squarefree kernel for each index
	sf := make([]int, n+1)
	sf[1] = 1
	for i := 2; i <= n; i++ {
		p := spf[i]
		cnt := 0
		x := i / p
		for x%p == 0 {
			x /= p
			cnt++
		}
		if cnt%2 == 1 {
			// p appears odd number of times => include p in squarefree kernel
			sf[i] = p * sf[x]
		} else {
			// p appears even number of times => squarefree kernel same as x
			p2 := p * p
			if x%p2 == 0 {
				// p^2 still divides (after removing pairs), further reduce
				// Actually we already divided out all p's, so just use x
			}
			sf[i] = sf[x]
		}
	}

	// Group sums by squarefree kernel
	sumByCore := make(map[int]int64)
	var ans int64 = 0
	for i := 1; i <= n; i++ {
		core := sf[i]
		sumByCore[core] += int64(nums[i-1])
		if sumByCore[core] > ans {
			ans = sumByCore[core]
		}
	}

	return ans
}

func main() {
	// Example 1: n=4, nums=[1,2,3,4]
	// Squarefree kernels: 1->1, 2->2, 3->3, 4->1
	// Groups: sf(1)={1,4}: sum=5, sf(2)={2}: sum=2, sf(3)={3}: sum=3 => max=5
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{1, 2, 3, 4}))

	// Example 2: n=3, nums=[10,20,30]
	// Kernels: 1->1, 2->2, 3->3
	// Groups: {1}:10, {2}:20, {3}:30 => max=30
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{10, 20, 30}))

	// Example 3: n=5, nums=[5,10,15,20,25]
	// Kernels: 1->1, 2->2, 3->3, 4->1, 5->5
	// sf(1)={1,4}: 5+20=25, sf(2)={2}:10, sf(3)={3}:15, sf(5)={5}:25
	// max=25
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{5, 10, 15, 20, 25}))

	// n=1
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{100}))

	// n=6, nums=[-1,-2,-3,-4,-5,-6]
	// sf(1)={1,4}: -5, sf(2)={2,8}... n=6 so sf(2)={2}: -2
	// sf(3)={3}: -3, sf(5)={5}: -5, sf(6)={2,3,6}: center each at index 6
	// 6 = 2*3, both with odd exponent 1 => sf(6)=6. So {6}: -6
	// max = max(-2, -3, -5, -5, -6, -5 from sf(1)) = -2
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{-1, -2, -3, -4, -5, -6}))

	// All same squarefree kernel
	fmt.Println(maximumElementSumOfCompleteSubsetOfIndices([]int{2, 3, 5}))
}
