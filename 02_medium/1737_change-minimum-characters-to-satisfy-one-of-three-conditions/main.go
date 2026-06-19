package main

// LeetCode #1737: Change Minimum Characters to Satisfy One of Three Conditions
// https://leetcode.com/problems/change-minimum-characters-to-satisfy-one-of-three-conditions/
// Difficulty: Medium
// Time: O(n + m), Space: O(26)

import "fmt"

func minCharacters(a string, b string) int {
	countA := make([]int, 26)
	countB := make([]int, 26)

	for _, ch := range a {
		countA[ch-'a']++
	}
	for _, ch := range b {
		countB[ch-'a']++
	}

	m, n := len(a), len(b)

	// Condition 3: make all characters in both strings the same
	ans := m + n
	for i := 0; i < 26; i++ {
		changes := (m - countA[i]) + (n - countB[i])
		if changes < ans {
			ans = changes
		}
	}

	// Condition 1: a < b lexicographically (every char in a < every char in b)
	// Condition 2: b < a lexicographically (every char in b < every char in a)
	for condition := 0; condition < 2; condition++ {
		prefixA := 0
		prefixB := 0
		for i := 0; i < 25; i++ {
			prefixA += countA[i]
			prefixB += countB[i]
			if condition == 0 {
				// a < b: change all a[i..25] to some smaller char, change all b[0..i] to some larger char
				changes := (m - prefixA) + prefixB
				if changes < ans {
					ans = changes
				}
			} else {
				// b < a: change all b[i..25] to some smaller char, change all a[0..i] to some larger char
				changes := (n - prefixB) + prefixA
				if changes < ans {
					ans = changes
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minCharacters("aba", "caa")) // Expected: 2
	fmt.Println(minCharacters("dabadd", "cda")) // Expected: 3
	fmt.Println(minCharacters("a", "a")) // Expected: 2
}
