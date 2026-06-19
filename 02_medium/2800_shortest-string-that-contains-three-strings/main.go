package main

// LeetCode #2800: Shortest String That Contains Three Strings
// https://leetcode.com/problems/shortest-string-that-contains-three-strings/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func merge2(x, y string) string {
	if len(x) < len(y) {
		return merge2(y, x)
	}
	// Check if x contains y
	for i := 0; i <= len(x)-len(y); i++ {
		if x[i:i+len(y)] == y {
			return x
		}
	}
	// Find longest overlap
	for i := 0; i < len(y); i++ {
		overlap := len(y) - i
		if overlap <= len(x) && x[len(x)-overlap:] == y[:overlap] {
			return x + y[i:]
		}
	}
	return x + y
}

func ShortestStringThatContainsThreeStrings(a string, b string, c string) string {
	strs := []string{a, b, c}
	perms := [][]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
		{1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}

	best := a + b + c
	for _, p := range perms {
		merged := merge2(merge2(strs[p[0]], strs[p[1]]), strs[p[2]])
		if len(merged) < len(best) || (len(merged) == len(best) && merged < best) {
			best = merged
		}
	}

	return best
}

func main() {
	fmt.Println(ShortestStringThatContainsThreeStrings("abc", "bcd", "cde"))
	fmt.Println(ShortestStringThatContainsThreeStrings("a", "ab", "abc"))
}
