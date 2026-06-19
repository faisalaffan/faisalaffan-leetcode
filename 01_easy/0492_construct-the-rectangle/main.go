package main

// LeetCode #492: Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func ConstructTheRectangle(area int) []int {
	w := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			w = i
		}
	}
	return []int{area / w, w}
}

func main() {
	fmt.Println(ConstructTheRectangle(4))
	fmt.Println(ConstructTheRectangle(37))
	fmt.Println(ConstructTheRectangle(122122))
}
