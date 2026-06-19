package main

// LeetCode #294: Flip Game II
// https://leetcode.com/problems/flip-game-ii/
// Difficulty: Medium [Paid]
// Time: O(n!!) worst case with memo, Space: O(n!)

import "fmt"

func canWin(currentState string) bool {
	memo := make(map[string]bool)
	return canWinHelper(currentState, memo)
}

func canWinHelper(state string, memo map[string]bool) bool {
	if res, ok := memo[state]; ok {
		return res
	}

	bytes := []byte(state)
	for i := 0; i < len(state)-1; i++ {
		if bytes[i] == '+' && bytes[i+1] == '+' {
			bytes[i], bytes[i+1] = '-', '-'
			opponentWins := canWinHelper(string(bytes), memo)
			bytes[i], bytes[i+1] = '+', '+'

			if !opponentWins {
				memo[state] = true
				return true
			}
		}
	}

	memo[state] = false
	return false
}

func main() {
	fmt.Println(canWin("++++"))
	fmt.Println(canWin("+"))
	fmt.Println(canWin("+++"))
}
