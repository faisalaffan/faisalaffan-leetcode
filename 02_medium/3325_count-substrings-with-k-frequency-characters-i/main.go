package main

// LeetCode #3325: Count Substrings With K-Frequency Characters I
// https://leetcode.com/problems/count-substrings-with-k-frequency-characters-i/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2)) // 4
	fmt.Println(numberOfSubstrings("abcde", 1)) // 15
	fmt.Println(numberOfSubstrings("aaaa", 2))  // 3
}

func numberOfSubstrings(s string, k int) int {
	n := len(s)
	cnt := [26]int{}
	ans, left := 0, 0

	for right := 0; right < n; right++ {
		idx := s[right] - 'a'
		cnt[idx]++

		for cnt[idx] >= k {
			ans += n - right
			cnt[s[left]-'a']--
			left++
		}
	}

	return ans
}
