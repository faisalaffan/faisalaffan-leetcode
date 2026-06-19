package main

// LeetCode #948: Bag of Tokens
// https://leetcode.com/problems/bag-of-tokens/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score, maxScore := 0, 0

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}
	return maxScore
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
