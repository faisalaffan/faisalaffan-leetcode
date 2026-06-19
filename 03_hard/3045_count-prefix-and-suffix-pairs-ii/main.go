package main

// LeetCode #3045: Count Prefix and Suffix Pairs II
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/
// Difficulty: Hard

import "fmt"

type trieNode3045 struct {
	son map[[2]byte]*trieNode3045
	cnt int
}

func countPrefixSuffixPairs(words []string) int64 {
	var ans int64
	root := &trieNode3045{son: make(map[[2]byte]*trieNode3045)}
	for _, s := range words {
		cur := root
		n := len(s)
		for i := 0; i < n; i++ {
			key := [2]byte{s[i], s[n-1-i]}
			if cur.son[key] == nil {
				cur.son[key] = &trieNode3045{son: make(map[[2]byte]*trieNode3045)}
			}
			cur = cur.son[key]
			ans += int64(cur.cnt)
		}
		cur.cnt++
	}
	return ans
}

func main() {
	fmt.Println(countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
}
