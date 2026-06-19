package main

// LeetCode #916: Word Subsets
// https://leetcode.com/problems/word-subsets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
}

// Time: O((n + m) * L) | Space: O(1)
func WordSubsets(words1 []string, words2 []string) []string {
	maxCnt := [26]int{}
	for _, word := range words2 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if cnt[i] > maxCnt[i] {
				maxCnt[i] = cnt[i]
			}
		}
	}

	var ans []string
	for _, word := range words1 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if cnt[i] < maxCnt[i] {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, word)
		}
	}

	return ans
}
