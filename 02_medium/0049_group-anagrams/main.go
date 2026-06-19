package main

// LeetCode #49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Difficulty: Medium

import "fmt"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)

	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	// [["bat"],["nat","tan"],["ate","eat","tea"]]

	// Test case 2
	fmt.Println(groupAnagrams([]string{""})) // [[""]]

	// Test case 3
	fmt.Println(groupAnagrams([]string{"a"})) // [["a"]]
}

// Time: O(n * k) | Space: O(n * k)
