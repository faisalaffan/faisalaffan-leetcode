package main

// LeetCode #3274: Check if Two Chessboard Squares Have the Same Color
// https://leetcode.com/problems/check-if-two-chessboard-squares-have-the-same-color/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "c3"))
	fmt.Println(CheckIfTwoChessboardSquaresHaveTheSameColor("a1", "h3"))
}

// CheckIfTwoChessboardSquaresHaveTheSameColor returns true if both squares are the same color on a chessboard.
// Time: O(1). Space: O(1).
func CheckIfTwoChessboardSquaresHaveTheSameColor(coordinate1 string, coordinate2 string) bool {
	c1 := (int(coordinate1[0]-'a') + int(coordinate1[1]-'1')) % 2
	c2 := (int(coordinate2[0]-'a') + int(coordinate2[1]-'1')) % 2
	return c1 == c2
}
