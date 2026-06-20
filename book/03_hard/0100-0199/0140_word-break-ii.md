# 0140 — Word Break Ii

## Deskripsi

**Soal:** [0140. Word Break Ii](https://leetcode.com/problems/word-break-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func wordBreak(s string, wordDict []string) []string`

## Solusi Go

```go
package main

// LeetCode #140: Word Break II
// https://leetcode.com/problems/word-break-ii/
// Difficulty: Hard

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
  // Membuat map untuk pencarian O(1): key → value
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

  // Membuat map untuk pencarian O(1): key → value
	memo := make(map[string][]string)

	var dfs func(string) []string
	dfs = func(remaining string) []string {
		if results, ok := memo[remaining]; ok {
			return results
		}

		results := []string{}
		if wordSet[remaining] {
			results = append(results, remaining)
		}

		for i := 1; i < len(remaining); i++ {
			prefix := remaining[:i]
			if wordSet[prefix] {
				suffixSentences := dfs(remaining[i:])
				for _, sentence := range suffixSentences {
					results = append(results, prefix+" "+sentence)
				}
			}
		}

		memo[remaining] = results
		return results
	}

	result := dfs(s)

	// The problem expects sentences that don't include the last word as a standalone sentence
	// when the full string is a word itself. We need to filter if the full string matches
	// but we already handle that. But the example shows:
	// "catsanddog" -> ["cats and dog","cat sand dog"]
	// The full string "catsanddog" is NOT in the dict, so no issue in this case.

	return result
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	result := wordBreak(s, wordDict)
	fmt.Printf("wordBreak(%q, %v) = %v\n", s, wordDict, result)

	// Check that result contains expected sentences (order-independent)
	expected := map[string]bool{
		"cats and dog": true,
		"cat sand dog": true,
	}

	pass := true
	if len(result) != len(expected) {
		pass = false
	} else {
		for _, r := range result {
			if !expected[r] {
				pass = false
				break
			}
		}
	}

	if pass {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %v\n", strings.Join(func() []string {
  // Membuat slice untuk menyimpan hasil
			keys := make([]string, 0, len(expected))
			for k := range expected {
				keys = append(keys, k)
			}
			return keys
		}(), ", "))
	}
}
```
