package main

// LeetCode #3706: Maximum Distance Between Unequal Words in Array II
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-ii/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumDistanceBetweenUnequalWordsInArrayIi(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		if words[i] != words[0] {
			if i+1 > ans {
				ans = i + 1
			}
		}
		if words[i] != words[n-1] {
			if n-i > ans {
				ans = n - i
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(maximumDistanceBetweenUnequalWordsInArrayIi([]string{"z", "z", "z"}))
}
