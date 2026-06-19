package main

// LeetCode #3696: Maximum Distance Between Unequal Words in Array I
// https://leetcode.com/problems/maximum-distance-between-unequal-words-in-array-i/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"leetcode", "leetcode", "codeforces"}))
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"a", "b", "c", "a", "a"}))
	fmt.Println(MaximumDistanceBetweenUnequalWordsInArrayI([]string{"z", "z", "z"}))
}

// Time: O(n)
// Space: O(1)
func MaximumDistanceBetweenUnequalWordsInArrayI(words []string) int {
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
