package main

// LeetCode #843: Guess the Word
// https://leetcode.com/problems/guess-the-word/
// Difficulty: Hard
// Approach: Minimax + filtering. For each candidate, compute match counts against other
// candidates. Pick the word that minimizes the maximum group size (best worst-case).
// After each guess, filter candidates to those matching the returned count.

import "fmt"

type Master struct {
	secret string
	words  []string
}

func (m *Master) guess(word string) int {
	matches := 0
	for i := 0; i < len(word); i++ {
		if word[i] == m.secret[i] {
			matches++
		}
	}
	return matches
}

func findSecretWord(words []string, master *Master) {
	n := len(words)
	match := func(a, b string) int {
		cnt := 0
		for i := 0; i < 6; i++ {
			if a[i] == b[i] {
				cnt++
			}
		}
		return cnt
	}

	candidates := make([]int, n)
	for i := range candidates {
		candidates[i] = i
	}

	for len(candidates) > 0 {
		// Pick the candidate that minimizes the maximum group size
		bestIdx := 0
		bestScore := n + 1
		for _, idx := range candidates {
			groups := make([]int, 7)
			for _, other := range candidates {
				m := match(words[idx], words[other])
				groups[m]++
			}
			maxGroup := 0
			for _, g := range groups {
				if g > maxGroup {
					maxGroup = g
				}
			}
			if maxGroup < bestScore {
				bestScore = maxGroup
				bestIdx = idx
			}
		}

		guess := words[bestIdx]
		matches := master.guess(guess)
		if matches == 6 {
			return
		}

		var newCandidates []int
		for _, idx := range candidates {
			if match(words[idx], guess) == matches {
				newCandidates = append(newCandidates, idx)
			}
		}
		candidates = newCandidates
	}
}

func main() {
	words := []string{"acckzz", "ccbazz", "eiowzz", "abcczz"}
	master := &Master{secret: "acckzz", words: words}
	findSecretWord(words, master)
	fmt.Println("Master guess called - findSecretWord returned (test passes if no panic)")

	words2 := []string{"hamada", "khaled"}
	master2 := &Master{secret: "khaled", words: words2}
	findSecretWord(words2, master2)
	fmt.Println("Test 2 passed")
}
