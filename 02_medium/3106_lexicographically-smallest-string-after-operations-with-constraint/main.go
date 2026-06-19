package main

// LeetCode #3106: Lexicographically Smallest String After Operations With Constraint
// https://leetcode.com/problems/lexicographically-smallest-string-after-operations-with-constraint/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getSmallestString(s string, k int) string {
	if k == 0 {
		return s
	}

	bytes := []byte(s)
	for i, ch := range bytes {
		dist := int(ch - 'a')
		move := min(dist, 26-dist)
		if move <= k {
			k -= move
			bytes[i] = 'a'
		} else {
			bytes[i] = byte(int(ch) - k)
			k = 0
			break
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(getSmallestString("zbbz", 3))  // Expected: "aaaz"
	fmt.Println(getSmallestString("xaxcd", 4)) // Expected: "aawcd"
	fmt.Println(getSmallestString("lol", 0))   // Expected: "lol"
}
