package main

// LeetCode #488: Zuma Game
// https://leetcode.com/problems/zuma-game/
// Difficulty: Hard
// Approach: DFS + memoization. Try each hand ball at each position in the board.
// After placing, remove consecutive 3+ balls and continue recursively.

import (
	"fmt"
)

func main() {
	fmt.Println("488 - Zuma Game")

	// Test cases
	fmt.Printf("findMinStep(\"WRRBBW\", \"RB\") = %d (expected: -1)\n",
		findMinStep("WRRBBW", "RB"))
	fmt.Printf("findMinStep(\"WWRRBBWW\", \"WRBRW\") = %d (expected: 2)\n",
		findMinStep("WWRRBBWW", "WRBRW"))
	fmt.Printf("findMinStep(\"G\", \"GGGGG\") = %d (expected: 2)\n",
		findMinStep("G", "GGGGG"))
	fmt.Printf("findMinStep(\"RBYYBBRRB\", \"YRBGB\") = %d (expected: 3)\n",
		findMinStep("RBYYBBRRB", "YRBGB"))
	fmt.Printf("findMinStep(\"\", \"\") = %d (expected: 0)\n",
		findMinStep("", ""))
	fmt.Printf("findMinStep(\"RRWWRRBBRR\", \"WB\") = %d\n",
		findMinStep("RRWWRRBBRR", "WB"))
}

func findMinStep(board string, hand string) int {
	// Count hand balls
	handCount := make([]int, 26)
	for _, ch := range hand {
		handCount[ch-'A']++
	}

	result := dfs(board, handCount)
	if result > len(hand) {
		return -1
	}
	return result
}

func dfs(board string, handCount []int) int {
	if board == "" {
		return 0
	}

	// Prune: if there are balls in hand that don't exist on board, skip
	// and also if count is insufficient

	// Try every possible placement
	minUsed := len(handCount)*5 + 1 // larger than any possible answer

	for i := 0; i < len(board); i++ {
		// Try to insert a ball from hand
		for color := 0; color < 26; color++ {
			if handCount[color] == 0 {
				continue
			}
			// Check if this placement makes sense: the ball color should match
			// a neighbor, or there should be at least 2 of the same color
			ball := byte('A' + color)
			if board[i] == ball {
				// Insert ball at position i
				handCount[color]--
				newBoard := removeConsecutive(board[:i] + string(ball) + board[i:])
				used := 1 + dfs(newBoard, handCount)
				if used < minUsed {
					minUsed = used
				}
				handCount[color]++
			} else if i > 0 && i < len(board)-1 && board[i-1] == board[i+1] && board[i-1] == ball {
				// Special case: inserting between two same balls to form 3
				handCount[color]--
				newBoard := removeConsecutive(board[:i] + string(ball) + board[i:])
				used := 1 + dfs(newBoard, handCount)
				if used < minUsed {
					minUsed = used
				}
				handCount[color]++
			}
		}

		// Skip consecutive same characters to avoid duplicate placements
		for i+1 < len(board) && board[i] == board[i+1] {
			i++
		}
	}

	return minUsed
}

// removeConsecutive removes all groups of 3+ consecutive same characters
func removeConsecutive(s string) string {
	// Keep removing until no more groups of 3+
	for {
		newS := removeOnce(s)
		if newS == s {
			break
		}
		s = newS
	}
	return s
}

func removeOnce(s string) string {
	if len(s) < 3 {
		return s
	}

	// Find groups of 3+ consecutive same characters
	result := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if j-i < 3 {
			// Keep this group
			result = append(result, s[i:j]...)
		}
		// If >=3, skip (remove them)
		i = j
	}
	return string(result)
}
