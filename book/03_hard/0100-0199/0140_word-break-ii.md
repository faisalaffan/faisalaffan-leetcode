# 0140 — Word Break Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func wordBreak(s string, wordDict []string) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

  // HashMap: O(1) lookup
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
			keys := make([]string, 0, len(expected))
			for k := range expected {
				keys = append(keys, k)
			}
			return keys
		}(), ", "))
	}
}
```
