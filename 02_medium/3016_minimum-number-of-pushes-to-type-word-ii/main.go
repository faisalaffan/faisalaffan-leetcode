package main

// LeetCode #3016: Minimum Number of Pushes to Type Word II
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumPushes("abcde"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

func minimumPushes(word string) int {
	cnt := make([]int, 26)
	for _, ch := range word {
		cnt[ch-'a']++
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	ans := 0
	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return ans
}
