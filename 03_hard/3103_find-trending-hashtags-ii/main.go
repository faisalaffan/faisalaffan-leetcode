package main

// LeetCode #3103: Find Trending Hashtags II
// https://leetcode.com/problems/find-trending-hashtags-ii/
// Difficulty: Hard [Paid]
//
// Given a list of tweets (each with tweet_id, user_id, tweet_date, tweet_text),
// find the top 3 trending hashtags. A hashtag is a word starting with '#'.
// Sort results by count descending, then hashtag alphabetically.
// "II" variant: considers tag frequency across all tweets (not per-tweet dedup).

import (
	"fmt"
	"sort"
	"strings"
)

type HashtagCount struct {
	hashtag string
	count   int
}

func findTrendingHashtags(tweets [][]string) []string {
	counts := make(map[string]int)

	for _, tweet := range tweets {
		if len(tweet) < 4 {
			continue
		}
		text := tweet[3]

		// Extract hashtags: words starting with #
		words := strings.Fields(text)
		for _, word := range words {
			if len(word) > 1 && word[0] == '#' {
				hashtag := strings.ToLower(word)
				counts[hashtag]++
			}
		}
	}

	// Convert to slice and sort by count desc, then hashtag asc
	var hcs []HashtagCount
	for h, c := range counts {
		hcs = append(hcs, HashtagCount{h, c})
	}
	sort.Slice(hcs, func(i, j int) bool {
		if hcs[i].count != hcs[j].count {
			return hcs[i].count > hcs[j].count
		}
		return hcs[i].hashtag < hcs[j].hashtag
	})

	// Return top 3
	topN := 3
	if len(hcs) < topN {
		topN = len(hcs)
	}
	result := make([]string, topN)
	for i := 0; i < topN; i++ {
		result[i] = hcs[i].hashtag
	}
	return result
}

func main() {
	// Test case 1
	tweets1 := [][]string{
		{"1", "u1", "2024-01-01", "#tech is great #coding"},
		{"2", "u2", "2024-01-01", "#python #coding"},
		{"3", "u3", "2024-01-02", "#tech is the future"},
		{"4", "u1", "2024-01-02", "#coding makes me happy"},
		{"5", "u2", "2024-01-03", "#python #python"},
	}
	fmt.Println("Test 1:", findTrendingHashtags(tweets1))
	// Expected: [#coding #python #tech] or [#coding #tech #python] depending on counts

	// Test case 2: no tweets
	tweets2 := [][]string{}
	fmt.Println("Test 2:", findTrendingHashtags(tweets2))
	// Expected: []

	// Test case 3: no hashtags
	tweets3 := [][]string{
		{"1", "u1", "2024-01-01", "hello world"},
	}
	fmt.Println("Test 3:", findTrendingHashtags(tweets3))
	// Expected: []

	// Test case 4: tie-breaking
	tweets4 := [][]string{
		{"1", "u1", "2024-01-01", "#abc nothing"},
		{"2", "u2", "2024-01-01", "#xyz nothing"},
	}
	fmt.Println("Test 4:", findTrendingHashtags(tweets4))
	// Expected: [#abc #xyz] (alphabetical tie-break)
}
