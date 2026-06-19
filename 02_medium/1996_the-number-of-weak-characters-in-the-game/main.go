package main

// LeetCode #1996: The Number of Weak Characters in the Game
// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{5, 5}, {6, 3}, {3, 6}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{2, 2}, {3, 3}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{1, 5}, {10, 4}, {4, 3}}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func TheNumberOfWeakCharactersInTheGame(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	ans := 0
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return ans
}
