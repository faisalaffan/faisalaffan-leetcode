package main

// LeetCode #1073: Adding Two Negabinary Numbers
// https://leetcode.com/problems/adding-two-negabinary-numbers/
// Difficulty: Medium
//
// Approach: Sum digits from right to left with carry in base -2
// Time: O(max(m, n))
// Space: O(max(m, n))

import "fmt"

func main() {
	fmt.Println(addNegabinary([]int{1, 1, 1, 1, 1}, []int{1, 0, 1})) // [1,0,0,0,0]
	fmt.Println(addNegabinary([]int{0}, []int{0}))                   // [0]
}

func addNegabinary(arr1 []int, arr2 []int) []int {
	i, j := len(arr1)-1, len(arr2)-1
	carry := 0
	result := make([]int, 0)

	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >= 0 {
			carry += arr2[j]
			j--
		}

		result = append(result, carry&1)
		carry = -(carry >> 1)
	}

	// Reverse and remove leading zeros
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	// Remove leading zeros
	start := 0
	for start < len(result)-1 && result[start] == 0 {
		start++
	}

	return result[start:]
}
