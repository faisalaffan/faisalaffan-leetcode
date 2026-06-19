package main

// LeetCode #1872: Stone Game VIII
// https://leetcode.com/problems/stone-game-viii/
// Difficulty: Hard

import "fmt"

func stoneGameViii(stones []int) int {
	n := len(stones)
	prefix := make([]int, n)
	prefix[0] = stones[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	// dp[i] = maximum score difference (current player - other) when
	// considering stones from index i onward (i is the first stone).
	// The player picks some index j >= i, takes prefix[j], then the rest
	// becomes a game for the other player.
	// dp[i] = max(prefix[j] - dp[j+1]) for j in [i, n-2]
	// dp[n-1] = 0 (only one stone, can't make a move with i < n-1 requirement)

	// dp[i] = max(prefix[i] - dp[i+1], dp[i+1])
	dpI := 0
	for i := n - 2; i >= 0; i-- {
		dpI = max(prefix[i]-dpI, dpI)
	}
	return dpI
}

func main() {
	// Example: [-1,2,-3,4,-5] -> 5
	fmt.Println(stoneGameViii([]int{-1, 2, -3, 4, -5}))

	// Additional test
	fmt.Println(stoneGameViii([]int{1, 2, 3, 4, 5}))
}
