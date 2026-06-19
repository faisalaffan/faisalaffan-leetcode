package main

// LeetCode #293: Flip Game
// https://leetcode.com/problems/flip-game/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(n) for output
func GeneratePossibleNextMoves(currentState string) []string {
	var res []string
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			flipped := currentState[:i] + "--" + currentState[i+2:]
			res = append(res, flipped)
		}
	}
	return res
}

func main() {
	fmt.Println(GeneratePossibleNextMoves("++++"))
	fmt.Println(GeneratePossibleNextMoves("+"))
}
