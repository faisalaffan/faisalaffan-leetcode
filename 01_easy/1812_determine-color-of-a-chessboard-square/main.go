package main

// LeetCode #1812: Determine Color of a Chessboard Square
// https://leetcode.com/problems/determine-color-of-a-chessboard-square/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func SquareIsWhite(coordinates string) bool {
	// (col + row) % 2 == 0 means dark, == 1 means light (white)
	col := int(coordinates[0] - 'a' + 1)
	row := int(coordinates[1] - '0')
	return (col+row)%2 == 1
}

func main() {
	fmt.Println(SquareIsWhite("a1"))
	fmt.Println(SquareIsWhite("h3"))
	fmt.Println(SquareIsWhite("c7"))
}
