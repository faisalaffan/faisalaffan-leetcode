package main

// LeetCode #3076: Shortest Uncommon Substring in an Array
// https://leetcode.com/problems/shortest-uncommon-substring-in-an-array/
// Difficulty: Medium
// Time: O(n * L^2) | Space: O(n * L^2)

import "fmt"

func main() {
	fmt.Println(shortestUncommonSubstring([]string{"cab", "ad", "bad", "c"}))
	fmt.Println(shortestUncommonSubstring([]string{"abc", "bcd", "abcd"}))
}

func shortestUncommonSubstring(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)

	for i := 0; i < n; i++ {
		subs := map[string]bool{}
		for l := 0; l < len(arr[i]); l++ {
			for r := l + 1; r <= len(arr[i]); r++ {
				subs[arr[i][l:r]] = true
			}
		}
		best := ""
		for s := range subs {
			common := false
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				if contains(arr[j], s) {
					common = true
					break
				}
			}
			if !common {
				if best == "" || len(s) < len(best) || (len(s) == len(best) && s < best) {
					best = s
				}
			}
		}
		ans[i] = best
	}
	return ans
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
