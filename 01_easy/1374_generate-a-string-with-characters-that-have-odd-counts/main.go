package main

// LeetCode #1374: Generate a String With Characters That Have Odd Counts
// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
// Difficulty: Easy
//
// LeetCode submission: func generateTheString(n int) string

import "fmt"

func main() {
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(4)) // "aaab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(2)) // "ab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(7)) // "aaaaaaa"
}

// Time: O(n), Space: O(n)
func GenerateAStringWithCharactersThatHaveOddCounts(n int) string {
	if n%2 == 1 {
		return string(makeN('a', n))
	}
	return string(makeN('a', n-1)) + "b"
}

func makeN(ch byte, n int) []byte {
	res := make([]byte, n)
	for i := range res {
		res[i] = ch
	}
	return res
}
