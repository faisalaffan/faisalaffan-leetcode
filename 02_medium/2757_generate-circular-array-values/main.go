package main

// LeetCode #2757: Generate Circular Array Values
// https://leetcode.com/problems/generate-circular-array-values/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func GenerateCircularArrayValues(arr []int, start int, count int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = arr[(start+i)%n]
	}
	return result
}

func main() {
	fmt.Println(GenerateCircularArrayValues([]int{1, 2, 3, 4}, 2, 6))
	fmt.Println(GenerateCircularArrayValues([]int{10, 20}, 1, 3))
}
