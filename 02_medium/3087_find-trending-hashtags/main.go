package main

// LeetCode #3087: Find Trending Hashtags
// https://leetcode.com/problems/find-trending-hashtags/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func findTrendingHashtags(tweets []string) []string {
	freq := make(map[string]int)
	for _, tweet := range tweets {
		words := strings.Fields(tweet)
		for _, w := range words {
			if strings.HasPrefix(w, "#") {
				hashtag := strings.ToLower(w)
				freq[hashtag]++
			}
		}
	}

	type ht struct {
		tag string
		cnt int
	}
	var list []ht
	for tag, cnt := range freq {
		list = append(list, ht{tag, cnt})
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].cnt != list[j].cnt {
			return list[i].cnt > list[j].cnt
		}
		return list[i].tag < list[j].tag
	})

	ans := make([]string, 0, min(3, len(list)))
	for i := 0; i < min(3, len(list)); i++ {
		ans = append(ans, list[i].tag)
	}
	return ans
}

func main() {
	fmt.Println(findTrendingHashtags([]string{
		"Good morning #tech #coding",
		"Loving #coding today #go",
		"#Tech is amazing #golang",
		"#coding #coding #coding",
	}))
	fmt.Println(findTrendingHashtags([]string{
		"#a #b #c",
		"#a #b",
		"#a",
	}))
}
