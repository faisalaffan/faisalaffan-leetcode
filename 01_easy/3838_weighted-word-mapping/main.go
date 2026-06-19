package main

// LeetCode #3838: Weighted Word Mapping
// https://leetcode.com/problems/weighted-word-mapping/
// Difficulty: Easy

import "fmt"

func main() {
	weights1 := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}
	fmt.Println(WeightedWordMapping([]string{"abcd", "def", "xyz"}, weights1))

	weights2 := make([]int, 26)
	for i := range weights2 {
		weights2[i] = 1
	}
	fmt.Println(WeightedWordMapping([]string{"a", "b", "c"}, weights2))
}

// Time: O(N * L)
// Space: O(1)
func WeightedWordMapping(words []string, weights []int) string {
	result := make([]byte, len(words))
	for i, word := range words {
		total := 0
		for _, ch := range word {
			total += weights[ch-'a']
		}
		result[i] = byte('z' - total%26)
	}
	return string(result)
}
