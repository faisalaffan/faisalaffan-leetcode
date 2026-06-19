package main

// LeetCode #2682: Find the Losers of the Circular Game
// https://leetcode.com/problems/find-the-losers-of-the-circular-game/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheLosersOfTheCircularGame(5, 2))
	fmt.Println(FindTheLosersOfTheCircularGame(4, 4))
}

func FindTheLosersOfTheCircularGame(n int, k int) []int {
	visited := make([]bool, n)
	i := 0
	step := k
	for !visited[i] {
		visited[i] = true
		i = (i + step) % n
		step += k
	}

	result := []int{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			result = append(result, i+1) // 1-indexed
		}
	}
	return result
}
