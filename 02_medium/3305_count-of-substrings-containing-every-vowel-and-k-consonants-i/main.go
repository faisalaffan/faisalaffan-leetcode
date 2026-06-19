package main

// LeetCode #3305: Count of Substrings Containing Every Vowel and K Consonants I
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))         // 0
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstrings("aeiou", 0))          // 1
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var ans int64
	for i := 0; i < n; i++ {
		vowelCnt := make(map[byte]int)
		cons := 0
		for j := i; j < n; j++ {
			if isVowel(word[j]) {
				vowelCnt[word[j]]++
			} else {
				cons++
				if cons > k {
					break
				}
			}
			if cons == k && len(vowelCnt) == 5 {
				ans++
			}
		}
	}
	return ans
}
