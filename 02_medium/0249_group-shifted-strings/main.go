package main

// LeetCode #249: Group Shifted Strings
// https://leetcode.com/problems/group-shifted-strings/
// Difficulty: Medium [Paid]
// Time: O(n * m), Space: O(n * m)

import (
	"fmt"
	"strings"
)

func groupStrings(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func getKey(s string) string {
	if len(s) == 0 {
		return ""
	}

	shift := s[0] - 'a'
	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		diff := (int(s[i]-'a') - int(shift) + 26) % 26
		sb.WriteByte(byte(diff + 'a'))
	}

	return sb.String()
}

func main() {
	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
	fmt.Println(groupStrings([]string{"a"}))
	fmt.Println(groupStrings([]string{"ab", "bc", "cd"}))
}
