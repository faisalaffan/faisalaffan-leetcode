package main

// LeetCode #2399: Check Distances Between Same Letters
// https://leetcode.com/problems/check-distances-between-same-letters/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CheckDistancesBetweenSameLetters("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})) // true
	fmt.Println(CheckDistancesBetweenSameLetters("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))    // false
}

func CheckDistancesBetweenSameLetters(s string, distance []int) bool {
	first := [26]int{}
	for i := 0; i < 26; i++ {
		first[i] = -1
	}
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if first[idx] == -1 {
			first[idx] = i
		} else if i-first[idx]-1 != distance[idx] {
			return false
		}
	}
	return true
}
