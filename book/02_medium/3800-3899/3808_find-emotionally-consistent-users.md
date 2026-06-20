# 3808 — Find Emotionally Consistent Users

## Deskripsi

**Soal:** [3808. Find Emotionally Consistent Users](https://leetcode.com/problems/find-emotionally-consistent-users/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N + U log U) where N = reactions, U = unique users  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3808: Find Emotionally Consistent Users
// https://leetcode.com/problems/find-emotionally-consistent-users/
// Difficulty: Medium (SQL problem — implemented in Go)

import (
	"fmt"
	"sort"
)

type Reaction struct {
	UserID    int
	ContentID int
	Reaction  string
}

type ConsistentUser struct {
	UserID           int
	DominantReaction string
	ReactionRatio    float64
}

func main() {
	reactions := []Reaction{
		{1, 1, "like"}, {1, 2, "like"}, {1, 3, "like"}, {1, 4, "like"}, {1, 5, "love"},
		{2, 1, "like"}, {2, 2, "love"}, {2, 3, "wow"}, {2, 4, "sad"}, {2, 5, "angry"},
		{3, 1, "love"}, {3, 2, "love"}, {3, 3, "love"}, {3, 4, "love"}, {3, 5, "love"},
	}

	result := FindEmotionallyConsistentUsers(reactions)
	for _, r := range result {
		fmt.Printf("%d %s %.2f\n", r.UserID, r.DominantReaction, r.ReactionRatio)
	}
}

// Time: O(N + U log U) where N = reactions, U = unique users
// Space: O(N)
func FindEmotionallyConsistentUsers(reactions []Reaction) []ConsistentUser {
	// userID -> reaction -> count
  // Membuat map untuk pencarian O(1): key → value
	userReactionCounts := make(map[int]map[string]int)
  // Membuat map untuk pencarian O(1): key → value
	userTotal := make(map[int]int)

	for _, r := range reactions {
		if userReactionCounts[r.UserID] == nil {
			userReactionCounts[r.UserID] = make(map[string]int)
		}
		userReactionCounts[r.UserID][r.Reaction]++
		userTotal[r.UserID]++
	}

	var result []ConsistentUser
	for userID, reactionCounts := range userReactionCounts {
		total := userTotal[userID]
		if total < 5 {
			continue
		}

		maxCnt := 0
		dominant := ""
		for reaction, cnt := range reactionCounts {
			if cnt > maxCnt || (cnt == maxCnt && reaction < dominant) {
				maxCnt = cnt
				dominant = reaction
			}
		}

		ratio := float64(maxCnt) / float64(total)
		if ratio >= 0.60 {
			result = append(result, ConsistentUser{
				UserID:           userID,
				DominantReaction: dominant,
				ReactionRatio:    ratio,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].ReactionRatio != result[j].ReactionRatio {
			return result[i].ReactionRatio > result[j].ReactionRatio
		}
		return result[i].UserID < result[j].UserID
	})

	return result
}
```
