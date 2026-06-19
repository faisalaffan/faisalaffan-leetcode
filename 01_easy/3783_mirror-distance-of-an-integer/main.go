package main

// LeetCode #3783: Mirror Distance of an Integer
// https://leetcode.com/problems/mirror-distance-of-an-integer/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MirrorDistanceOfAnInteger(25))
	fmt.Println(MirrorDistanceOfAnInteger(10))
	fmt.Println(MirrorDistanceOfAnInteger(7))
}

// Time: O(log n)
// Space: O(1)
func MirrorDistanceOfAnInteger(n int) int {
	rev := 0
	for x := n; x > 0; x /= 10 {
		rev = rev*10 + x%10
	}
	if n > rev {
		return n - rev
	}
	return rev - n
}
