package main

// LeetCode #2964: Number of Divisible Triplet Sums
// https://leetcode.com/problems/number-of-divisible-triplet-sums/
// Difficulty: Medium

import "fmt"

func numberOfDivisibleTripletSums(nums []int, d int) int {
	n := len(nums)
	count := 0

	// pre[rem] = count of elements before current j with remainder rem
	pre := make(map[int]int)

	for j := 0; j < n; j++ {
		// For each k > j, find nums[i] (i < j) such that
		// (nums[i] + nums[j] + nums[k]) % d == 0
		for k := j + 1; k < n; k++ {
			needed := (d - (nums[j]+nums[k])%d) % d
			count += pre[needed]
		}
		// Add current element's remainder to pre for future j
		pre[nums[j]%d]++
	}

	return count
}

func main() {
	// Test case 1: nums = [3,3,4,7,8], d = 5 -> 3
	fmt.Println(numberOfDivisibleTripletSums([]int{3, 3, 4, 7, 8}, 5)) // 3

	// Test case 2: nums = [3,3,3,3], d = 3 -> 4
	fmt.Println(numberOfDivisibleTripletSums([]int{3, 3, 3, 3}, 3)) // 4

	// Test case 3: nums = [1,2,3], d = 3 -> 1 (1+2+3=6, 6%3==0)
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 2, 3}, 3)) // 1

	// Test case 4: nums = [1,1,1], d = 1 -> 1 (all triplets sum to 3, 3%1==0)
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 1, 1}, 1)) // 1

	// Test case 5: nums = [1,2,3,4], d = 2
	// Triplets: (0,1,2)=6%2=0, (0,1,3)=7%2=1, (0,2,3)=8%2=0, (1,2,3)=9%2=1 -> 2
	fmt.Println(numberOfDivisibleTripletSums([]int{1, 2, 3, 4}, 2)) // 2
}

// Time: O(n^2) | Space: O(d)
