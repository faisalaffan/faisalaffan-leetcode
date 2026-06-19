package main

// LeetCode #1576: Replace All ?'s to Avoid Consecutive Repeating Characters
// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
// Difficulty: Easy
//
// LeetCode submission: func modifyString(s string) string

import "fmt"

func main() {
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("?zs")) // "azs"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("ubv?w")) // "ubvaw"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("??yw?ipkj?")) // "abywcipkja"
}

// Time: O(n), Space: O(n)
func ReplaceAllSToAvoidConsecutiveRepeatingCharacters(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for c := byte('a'); c <= 'z'; c++ {
				if (i == 0 || res[i-1] != c) && (i == len(res)-1 || res[i+1] != c) {
					res[i] = c
					break
				}
			}
		}
	}
	return string(res)
}
