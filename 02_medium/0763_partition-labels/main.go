package main

// LeetCode #763: Partition Labels
// https://leetcode.com/problems/partition-labels/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
	last := make([]int, 26)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	result := make([]int, 0)
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
