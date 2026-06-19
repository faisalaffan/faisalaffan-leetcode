package main

// LeetCode #3360: Stone Removal Game
// https://leetcode.com/problems/stone-removal-game/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(StoneRemovalGame(10))
	fmt.Println(StoneRemovalGame(7))
}

// StoneRemovalGame returns true if Alice wins the stone removal game.
// Alice goes first; they remove stones starting from 1 and increase by 1 each turn.
// Alice wins if she can make the last move.
// Time: O(sqrt(n)). Space: O(1).
func StoneRemovalGame(n int) bool {
	turn := 0 // 0 for Alice, 1 for Bob
	remove := 1
	for n >= remove {
		n -= remove
		remove++
		turn = 1 - turn
	}
	// If Alice made the last move, Alice wins, else Bob wins
	return turn != 0
}
