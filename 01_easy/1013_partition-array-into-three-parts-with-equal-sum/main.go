package main

// LeetCode #1013: Partition Array Into Three Parts With Equal Sum
// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1})) // true
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1})) // false
	fmt.Println(partitionArrayIntoThreePartsWithEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))    // true
}

// LeetCode submission: canThreePartsEqualSum
func canThreePartsEqualSum(arr []int) bool {
	total := 0
	for _, v := range arr {
		total += v
	}
	if total%3 != 0 {
		return false
	}
	target := total / 3
	sum, count := 0, 0
	for _, v := range arr {
		sum += v
		if sum == target {
			count++
			sum = 0
		}
	}
	return count >= 3
}

func partitionArrayIntoThreePartsWithEqualSum(arr []int) bool {
	return canThreePartsEqualSum(arr)
}
