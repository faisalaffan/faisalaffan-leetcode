package main

// LeetCode #1255: Maximum Score Words Formed by Letters
// https://leetcode.com/problems/maximum-score-words-formed-by-letters/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1255. Maximum Score Words Formed by Letters")
	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(words, letters, score, ":", maxScoreWords(words, letters, score), "(expected 23)")

	words2 := []string{"xxxz", "ax", "bx", "cx"}
	letters2 := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
	score2 := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
	fmt.Println(words2, letters2, score2, ":", maxScoreWords(words2, letters2, score2), "(expected 27)")
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterCount := [26]int{}
	for _, ch := range letters {
		letterCount[ch-'a']++
	}

	n := len(words)
	wordCount := make([][26]int, n)
	wordScore := make([]int, n)
	for i, word := range words {
		for _, ch := range word {
			wordCount[i][ch-'a']++
			wordScore[i] += score[ch-'a']
		}
	}

	maxScore := 0
	var dfs func(idx int, curScore int, avail [26]int)
	dfs = func(idx int, curScore int, avail [26]int) {
		if idx == n {
			if curScore > maxScore {
				maxScore = curScore
			}
			return
		}
		// Skip this word.
		dfs(idx+1, curScore, avail)
		// Try to include this word.
		canUse := true
		for i := 0; i < 26; i++ {
			if wordCount[idx][i] > avail[i] {
				canUse = false
				break
			}
		}
		if canUse {
			for i := 0; i < 26; i++ {
				avail[i] -= wordCount[idx][i]
			}
			dfs(idx+1, curScore+wordScore[idx], avail)
		}
	}

	dfs(0, 0, letterCount)
	return maxScore
}
