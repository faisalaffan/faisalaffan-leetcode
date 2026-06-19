package main

// LeetCode #3324: Find the Sequence of Strings Appeared on the Screen
// https://leetcode.com/problems/find-the-sequence-of-strings-appeared-on-the-screen/
// Difficulty: Medium
// Time: O(n * 26) Space: O(n * 26) for output

import "fmt"

func main() {
	fmt.Println(stringSequence("abc")) // [a aa ab aba abb abc]
	fmt.Println(stringSequence("ab"))  // [a aa ab]
	fmt.Println(stringSequence("z"))   // [a b c d e f g h i j k l m n o p q r s t u v w x y z]
}

func stringSequence(target string) []string {
	var result []string
	var cur []byte

	for _, ch := range target {
		cur = append(cur, 'a')
		result = append(result, string(cur))
		for cur[len(cur)-1] != byte(ch) {
			cur[len(cur)-1]++
			result = append(result, string(cur))
		}
	}

	return result
}
