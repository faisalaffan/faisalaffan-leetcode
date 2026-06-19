package main

// LeetCode #1849: Splitting a String Into Descending Consecutive Values
// https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SplitString("1234"))
	fmt.Println(SplitString("050043"))
	fmt.Println(SplitString("9080701"))
}

// Time: O(n^2), Space: O(n) for recursion
func SplitString(s string) bool {
	var dfs func(idx int, prev int64, count int) bool
	dfs = func(idx int, prev int64, count int) bool {
		if idx == len(s) {
			return count >= 2
		}
		num := int64(0)
		for i := idx; i < len(s); i++ {
			num = num*10 + int64(s[i]-'0')
			if num > 1<<62 {
				break
			}
			if count == 0 || prev-num == 1 {
				if dfs(i+1, num, count+1) {
					return true
				}
			}
			if num == 0 {
				break
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
