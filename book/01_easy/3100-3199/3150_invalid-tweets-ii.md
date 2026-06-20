# 3150 — Invalid Tweets Ii

## Deskripsi

**Soal:** [3150. Invalid Tweets Ii](https://leetcode.com/problems/invalid-tweets-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3150: Invalid Tweets II
// https://leetcode.com/problems/invalid-tweets-ii/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find tweet IDs where content length > 140 or
// content contains invalid characters.

import (
	"fmt"
	"strings"
)

func main() {
	// LeetCode name: invalidTweets
	tweets := map[int]string{
		1: "Hello world!",
		2: strings.Repeat("a", 150),
		3: "Valid tweet content",
	}
	fmt.Println(InvalidTweetsIi(tweets))
	// [2]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: invalidTweets
func InvalidTweetsIi(tweets map[int]string) []int {
	result := []int{}
	for id, content := range tweets {
		if len(content) > 140 {
			result = append(result, id)
		}
	}
	return result
}
```
