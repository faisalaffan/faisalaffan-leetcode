package main

// LeetCode #1629: Slowest Key
// https://leetcode.com/problems/slowest-key/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SlowestKey(releaseTimes []int, keysPressed string) byte {
	maxDuration := releaseTimes[0]
	result := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration || (duration == maxDuration && keysPressed[i] > result) {
			maxDuration = duration
			result = keysPressed[i]
		}
	}
	return result
}

func main() {
	fmt.Printf("%c\n", SlowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Printf("%c\n", SlowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}
