# 3793 — Find Users With High Token Usage

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindUsersWithHighTokenUsage(prompts []Prompt) []UserTokenStat`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3793: Find Users with High Token Usage
// https://leetcode.com/problems/find-users-with-high-token-usage/
// Difficulty: Easy

import (
	"fmt"
	"math"
	"sort"
)

type Prompt struct {
	UserID int
	Prompt string
	Tokens int
}

type UserTokenStat struct {
	UserID      int
	PromptCount int
	AvgTokens   float64
}

func main() {
	prompts1 := []Prompt{
		{1, "Write a blog outline", 120},
		{1, "Generate SQL query", 80},
		{1, "Summarize an article", 200},
		{2, "Create resume bullet", 60},
		{2, "Improve LinkedIn bio", 70},
		{3, "Explain neural networks", 300},
		{3, "Generate interview Q&A", 250},
		{3, "Write cover letter", 180},
		{3, "Optimize Python code", 220},
	}
	fmt.Println(FindUsersWithHighTokenUsage(prompts1))

	prompts2 := []Prompt{
		{1, "Hello", 10},
		{1, "World", 20},
		{1, "Test", 30},
	}
	fmt.Println(FindUsersWithHighTokenUsage(prompts2))
}

// Time: O(n log n)
// Space: O(n)
func FindUsersWithHighTokenUsage(prompts []Prompt) []UserTokenStat {
	// Group tokens by user
  // HashMap: O(1) lookup
	userMap := make(map[int][]int)
	for _, p := range prompts {
		userMap[p.UserID] = append(userMap[p.UserID], p.Tokens)
	}

	var result []UserTokenStat
	for uid, tokens := range userMap {
		if len(tokens) < 3 {
			continue
		}

		sum := 0
		for _, t := range tokens {
			sum += t
		}
		avg := float64(sum) / float64(len(tokens))

		hasAboveAvg := false
		for _, t := range tokens {
			if float64(t) > avg {
				hasAboveAvg = true
				break
			}
		}
		if !hasAboveAvg {
			continue
		}

		avg = math.Round(avg*100) / 100

		result = append(result, UserTokenStat{
			UserID:      uid,
			PromptCount: len(tokens),
			AvgTokens:   avg,
		})
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].AvgTokens != result[j].AvgTokens {
			return result[i].AvgTokens > result[j].AvgTokens
		}
		return result[i].UserID < result[j].UserID
	})

	return result
}
```
