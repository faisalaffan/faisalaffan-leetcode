package main

// LeetCode #3216: Lexicographically Smallest String After a Swap
// https://leetcode.com/problems/lexicographically-smallest-string-after-a-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestStringAfterASwap("45320"))
	fmt.Println(LexicographicallySmallestStringAfterASwap("001"))
}

// LexicographicallySmallestStringAfterASwap makes the smallest string by swapping one pair of adjacent same-parity digits where left > right.
// Time: O(n). Space: O(n).
func LexicographicallySmallestStringAfterASwap(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		if b[i] > b[i+1] && (b[i]%2 == b[i+1]%2) {
			b[i], b[i+1] = b[i+1], b[i]
			break
		}
	}
	return string(b)
}
