package main

// LeetCode #91: Decode Ways
// https://leetcode.com/problems/decode-ways/
// Difficulty: Medium

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 {
			dp[i] += dp[i-1]
		}

		twoDigits := int(s[i-2]-'0')*10 + oneDigit
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	fmt.Println(numDecodings("12")) // 2

	// Test case 2
	fmt.Println(numDecodings("226")) // 3

	// Test case 3
	fmt.Println(numDecodings("06")) // 0
}

// Time: O(n) | Space: O(n)
