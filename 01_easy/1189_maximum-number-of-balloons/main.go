package main

// LeetCode #1189: Maximum Number of Balloons
// https://leetcode.com/problems/maximum-number-of-balloons/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))           // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))    // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))            // 0
}

// LeetCode submission: maxNumberOfBalloons
func maxNumberOfBalloons(text string) int {
	count := [26]int{}
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	ans := count[1]           // b
	ans = min(ans, count[0])  // a
	ans = min(ans, count[11]/2) // l (needs 2)
	ans = min(ans, count[14]/2) // o (needs 2)
	ans = min(ans, count[13]) // n
	return ans
}
