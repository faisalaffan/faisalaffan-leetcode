package main

// LeetCode #3200: Maximum Height of a Triangle
// https://leetcode.com/problems/maximum-height-of-a-triangle/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumHeightOfATriangle(2, 4))
	fmt.Println(MaximumHeightOfATriangle(2, 1))
	fmt.Println(MaximumHeightOfATriangle(10, 10))
}

// maxHeight tries building a triangle starting with the given first color.
func maxHeight(red, blue int, firstRed bool) int {
	h := 0
	need := 1
	for {
		if firstRed {
			if red < need {
				break
			}
			red -= need
		} else {
			if blue < need {
				break
			}
			blue -= need
		}
		h++
		need++
		firstRed = !firstRed
	}
	return h
}

// MaximumHeightOfATriangle returns the maximum height of a triangle using red and blue balls.
// Time: O(sqrt(n)). Space: O(1).
func MaximumHeightOfATriangle(red int, blue int) int {
	h1 := maxHeight(red, blue, true)
	h2 := maxHeight(red, blue, false)
	if h1 > h2 {
		return h1
	}
	return h2
}
