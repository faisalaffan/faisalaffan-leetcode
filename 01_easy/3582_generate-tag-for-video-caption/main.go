package main

// LeetCode #3582: Generate Tag for Video Caption
// https://leetcode.com/problems/generate-tag-for-video-caption/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(GenerateTagForVideoCaption("Leetcode daily streak achieved"))
	fmt.Println(GenerateTagForVideoCaption("can I Go There"))
	fmt.Println(GenerateTagForVideoCaption("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
}

// Time: O(n)
// Space: O(n)
func GenerateTagForVideoCaption(caption string) string {
	words := strings.Fields(caption)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			runes := []rune(w)
			for j, r := range runes {
				if j == 0 {
					runes[j] = unicode.ToUpper(r)
				} else {
					runes[j] = unicode.ToLower(r)
				}
			}
			words[i] = string(runes)
		}
	}

	result := "#" + strings.Join(words, "")
	if len(result) > 100 {
		result = result[:100]
	}
	return result
}
