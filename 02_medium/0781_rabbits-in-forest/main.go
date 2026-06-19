package main

// LeetCode #781: Rabbits in Forest
// https://leetcode.com/problems/rabbits-in-forest/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

func numRabbits(answers []int) int {
	count := make(map[int]int)
	for _, a := range answers {
		count[a]++
	}

	result := 0
	for k, v := range count {
		groupSize := k + 1
		groups := (v + groupSize - 1) / groupSize
		result += groups * groupSize
	}

	return result
}
