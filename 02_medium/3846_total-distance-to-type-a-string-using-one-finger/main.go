package main

// LeetCode #3846: Total Distance to Type a String Using One Finger
// https://leetcode.com/problems/total-distance-to-type-a-string-using-one-finger/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Precompute keyboard positions, simulate typing from 'a'.

import "fmt"

func TotalDistanceToTypeAStringUsingOneFinger(s string) int {
	// Keyboard layout (row, col)
	keyboard := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

	pos := make(map[byte][2]int)
	for r, row := range keyboard {
		for c, ch := range row {
			pos[byte(ch)] = [2]int{r, c}
		}
	}

	total := 0
	cur := pos['a']
	for i := 0; i < len(s); i++ {
		next := pos[s[i]]
		dist := abs(cur[0]-next[0]) + abs(cur[1]-next[1])
		total += dist
		cur = next
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("hello")) // Expected: 17

	// Example 2
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("a")) // Expected: 0

	// Example 3
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("qaz")) // q: (0,0), a: (1,0), z: (2,0) = |0-1|+|0-0| + |1-2|+|0-0| = 1+1 = 2
}
