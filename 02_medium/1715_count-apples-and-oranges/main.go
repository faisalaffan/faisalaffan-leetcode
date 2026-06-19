package main

// LeetCode #1715: Count Apples and Oranges
// https://leetcode.com/problems/count-apples-and-oranges/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type Box struct {
	BoxID      int
	AppleCnt   int
	OrangeCnt  int
}

type Chest struct {
	ChestID    int
	AppleCnt   int
	OrangeCnt  int
}

func countFruits(boxes []Box, chestMap map[int]Chest) (int, int) {
	totalApples := 0
	totalOranges := 0
	for _, b := range boxes {
		totalApples += b.AppleCnt
		totalOranges += b.OrangeCnt
		if c, ok := chestMap[b.BoxID]; ok {
			totalApples += c.AppleCnt
			totalOranges += c.OrangeCnt
		}
	}
	return totalApples, totalOranges
}

func main() {
	boxes := []Box{
		{BoxID: 1, AppleCnt: 5, OrangeCnt: 3},
		{BoxID: 2, AppleCnt: 2, OrangeCnt: 8},
	}
	chests := map[int]Chest{
		1: {AppleCnt: 10, OrangeCnt: 4},
	}
	apples, oranges := countFruits(boxes, chests)
	fmt.Printf("Apples: %d, Oranges: %d\n", apples, oranges) // Expected: 17, 15
}
