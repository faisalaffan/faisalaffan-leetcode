package main

import "fmt"

// LeetCode #458: Poor Pigs
// https://leetcode.com/problems/poor-pigs/
// Difficulty: Hard
//
// Each pig has (minutesToTest/minutesToDie + 1) possible states:
// die at round 1, die at round 2, ..., survive all rounds.
// With P pigs, we can encode (states)^P bucket combinations.
// Find the smallest P such that states^P >= buckets.
// Uses iterative multiplication to avoid floating-point precision issues.

func main() {
	// Example 1: buckets=4, minToDie=15, minToTest=15 => states=2, 2^2 >= 4 => 2
	fmt.Println("Pigs:", poorPigs(4, 15, 15)) // 2
	// Example 2: buckets=1 => 0 pigs needed
	fmt.Println("Pigs:", poorPigs(1, 15, 15)) // 0
	// Example 3: buckets=1000, minToDie=15, minToTest=60 => states=5
	fmt.Println("Pigs:", poorPigs(1000, 15, 60)) // ceil(log5(1000)) = 5
	// Large exact power: states=5, 5^3=125 => 3
	fmt.Println("Pigs:", poorPigs(125, 1, 4)) // 3
	// Edge: 2 buckets, 1 round
	fmt.Println("Pigs:", poorPigs(2, 1, 1)) // 1
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets <= 1 {
		return 0
	}
	states := minutesToTest/minutesToDie + 1
	// Iterative approach: find smallest pigs such that states^pigs >= buckets
	pigs := 0
	covered := 1
	for covered < buckets {
		pigs++
		covered *= states
	}
	return pigs
}
