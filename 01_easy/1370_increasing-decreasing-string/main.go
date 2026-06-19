package main

// LeetCode #1370: Increasing Decreasing String
// https://leetcode.com/problems/increasing-decreasing-string/
// Difficulty: Easy
//
// LeetCode submission: func sortString(s string) string

import "fmt"

func main() {
	fmt.Println(IncreasingDecreasingString("aaaabbbbcccc")) // "abccbaabccba"
	fmt.Println(IncreasingDecreasingString("rat"))          // "art"
	fmt.Println(IncreasingDecreasingString("leetcode"))     // "cdelotee"
}

// Time: O(n), Space: O(n)
func IncreasingDecreasingString(s string) string {
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	res := make([]byte, 0, len(s))
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if freq[i] > 0 {
				res = append(res, byte('a'+i))
				freq[i]--
			}
		}
	}
	return string(res)
}
