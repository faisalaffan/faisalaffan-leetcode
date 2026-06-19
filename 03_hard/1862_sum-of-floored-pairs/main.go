package main

// LeetCode #1862: Sum of Floored Pairs
// https://leetcode.com/problems/sum-of-floored-pairs/
// Difficulty: Hard

import (
	"fmt"
)

func sumOfFlooredPairs(nums []int) int {
	const mod = 1_000_000_007
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	freq := make([]int, maxVal+1)
	for _, v := range nums {
		freq[v]++
	}
	prefix := make([]int, maxVal+1)
	for i := 1; i <= maxVal; i++ {
		prefix[i] = prefix[i-1] + freq[i]
	}

	ans := 0
	for x := 1; x <= maxVal; x++ {
		if freq[x] == 0 {
			continue
		}
		// For each multiple y = x * q, count how many numbers in [y, y+x-1]
		for y := x; y <= maxVal; y += x {
			q := y / x
			count := prefix[min(maxVal, y+x-1)] - prefix[y-1]
			ans = (ans + freq[x]*count%mod*q%mod) % mod
		}
	}
	return ans
}

func main() {
	// Example: [2,5,9] -> 10
	// floor(2/2)+floor(5/2)+floor(9/2)=1+2+4=7
	// floor(2/5)+floor(5/5)+floor(9/5)=0+1+1=2
	// floor(2/9)+floor(5/9)+floor(9/9)=0+0+1=1
	// Total: 7+2+1=10
	fmt.Println(sumOfFlooredPairs([]int{2, 5, 9}))

	// Additional test
	fmt.Println(sumOfFlooredPairs([]int{7, 7, 7, 7, 7, 7, 7}))
}
