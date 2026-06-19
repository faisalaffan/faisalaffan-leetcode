package main

// LeetCode #717: 1-bit and 2-bit Characters
// https://leetcode.com/problems/1-bit-and-2-bit-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))          // true
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))      // false
	fmt.Println(isOneBitCharacter([]int{0}))                // true
}

// isOneBitCharacter checks if the last character must be a one-bit character.
// Time: O(n). Space: O(1).
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == len(bits)-1
}
