package main

// LeetCode #2103: Rings and Rods
// https://leetcode.com/problems/rings-and-rods/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RingsAndRods("B0B6G0R6R0R6G9"))    // 1
	fmt.Println(RingsAndRods("B0R0G0R9R0B0G0"))    // 1
	fmt.Println(RingsAndRods("G4"))                 // 0
}

// Time: O(n), Space: O(1)
func RingsAndRods(rings string) int {
	rods := make([]int, 10)
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		rod := rings[i+1] - '0'
		switch color {
		case 'R':
			rods[rod] |= 1
		case 'G':
			rods[rod] |= 2
		case 'B':
			rods[rod] |= 4
		}
	}

	count := 0
	for _, v := range rods {
		if v == 7 { // R|G|B = 1|2|4 = 7
			count++
		}
	}
	return count
}
