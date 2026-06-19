package main

// LeetCode #2833: Furthest Point From Origin
// https://leetcode.com/problems/furthest-point-from-origin/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FurthestPointFromOrigin("L_RL__R"))
	fmt.Println(FurthestPointFromOrigin("_R__LL_"))
}

func FurthestPointFromOrigin(moves string) int {
	countL, countR, countUnderscore := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			countL++
		case 'R':
			countR++
		case '_':
			countUnderscore++
		}
	}
	diff := countL - countR
	if diff < 0 {
		diff = -diff
	}
	return diff + countUnderscore
}
