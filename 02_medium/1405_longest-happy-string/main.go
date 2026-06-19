package main

// LeetCode #1405: Longest Happy String
// https://leetcode.com/problems/longest-happy-string/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(longestDiverseString(1, 1, 7)) // "ccaccbcc" or "ccbccacc"

	// Test case 2
	fmt.Println(longestDiverseString(7, 1, 0)) // "aabaa"

	// Test case 3
	fmt.Println(longestDiverseString(0, 8, 11)) // "ccbccbbccbbccbbccbc"
}

type charCount struct {
	count int
	char  byte
}

// Time: O(a+b+c) - building the result string
// Space: O(1) - constant extra space
func longestDiverseString(a int, b int, c int) string {
	pairs := []charCount{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	result := make([]byte, 0, a+b+c)

	for {
		// Sort by remaining count descending
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].count > pairs[j].count
		})

		placed := false
		for i := 0; i < 3; i++ {
			if pairs[i].count == 0 {
				break
			}
			n := len(result)
			// Check if we can place this character
			if n >= 2 && result[n-1] == pairs[i].char && result[n-2] == pairs[i].char {
				continue
			}
			result = append(result, pairs[i].char)
			pairs[i].count--
			placed = true
			break
		}

		if !placed {
			break
		}
	}

	return string(result)
}
