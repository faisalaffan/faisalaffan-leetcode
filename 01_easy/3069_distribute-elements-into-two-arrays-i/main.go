package main

// LeetCode #3069: Distribute Elements Into Two Arrays I
// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: resultArray
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{2, 1, 3, 4})) // [2, 3, 4, 1]
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{5, 4, 3, 8})) // [5, 3, 4, 8]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: resultArray
func DistributeElementsIntoTwoArraysI(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}
