package main

// LeetCode #3712: Sum of Elements With Frequency Divisible by K
// https://leetcode.com/problems/sum-of-elements-with-frequency-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 2, 3, 3, 3, 3, 4}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{4, 4, 4, 1, 2, 3}, 3))
}

// Time: O(n)
// Space: O(1)
func SumOfElementsWithFrequencyDivisibleByK(nums []int, k int) int {
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}

	sum := 0
	for v := 1; v <= 100; v++ {
		if cnt[v]%k == 0 {
			sum += v * cnt[v]
		}
	}
	return sum
}
