package main

// LeetCode #477: Total Hamming Distance
// https://leetcode.com/problems/total-hamming-distance/
// Difficulty: Medium
// Time: O(n * 32) = O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalHammingDistance([]int{4, 14, 2}))
	fmt.Println(TotalHammingDistance([]int{4, 14, 4}))
}

func TotalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)

	for bit := 0; bit < 32; bit++ {
		countOnes := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				countOnes++
			}
		}
		countZeros := n - countOnes
		total += countOnes * countZeros
	}

	return total
}
