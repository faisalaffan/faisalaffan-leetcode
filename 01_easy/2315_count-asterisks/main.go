package main

// LeetCode #2315: Count Asterisks
// https://leetcode.com/problems/count-asterisks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountAsterisks("l|*e*et|c**o|*de|")) // 2
	fmt.Println(CountAsterisks("iamprogrammer"))      // 0
}

func CountAsterisks(s string) int {
	count := 0
	bar := false
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			bar = !bar
		} else if s[i] == '*' && !bar {
			count++
		}
	}
	return count
}
