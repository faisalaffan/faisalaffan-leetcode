package main

// LeetCode #3024: Type of Triangle
// https://leetcode.com/problems/type-of-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	fmt.Println(TypeOfTriangle([]int{3, 3, 3})) // equilateral
	fmt.Println(TypeOfTriangle([]int{3, 4, 5})) // scalene
	fmt.Println(TypeOfTriangle([]int{3, 3, 5})) // isosceles
	fmt.Println(TypeOfTriangle([]int{1, 2, 3})) // none
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: triangleType
func TypeOfTriangle(nums []int) string {
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]

	// Check if valid triangle
	if a+b <= c {
		return "none"
	}

	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || a == c {
		return "isosceles"
	}
	return "scalene"
}
