# 3832 — Find Users With Persistent Behavior Patterns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func findPersistentPatterns(actions []UserAction, minLength int) []UserPattern`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3832: Find Users with Persistent Behavior Patterns
// https://leetcode.com/problems/find-users-with-persistent-behavior-patterns/
// Difficulty: Hard (SQL / Database)
//
// Find users who exhibit the same behavior pattern for 3+ consecutive
// days. Port to Go with data structure simulation.

import (
	"fmt"
	"sort"
)

type UserAction struct {
	UserID    int
	Action    string
	ActionDay int
}

type UserPattern struct {
	UserID        int
	Pattern       string
	PatternLength int
}

func main() {
	// Example
	actions := []UserAction{
		{1, "login", 1},
		{1, "login", 2},
		{1, "login", 3},
		{2, "view", 1},
		{2, "view", 2},
	}
	fmt.Println(findPersistentPatterns(actions, 3))
	// Edge: no persistent users
	fmt.Println(findPersistentPatterns([]UserAction{{1, "a", 1}, {1, "b", 2}}, 3))
}

func findPersistentPatterns(actions []UserAction, minLength int) []UserPattern {
	// Group by user, sort by day
  // HashMap: O(1) lookup
	userActions := make(map[int][]struct {
		action string
		day    int
	})

	for _, a := range actions {
		userActions[a.UserID] = append(userActions[a.UserID], struct {
			action string
			day    int
		}{a.Action, a.ActionDay})
	}

	var result []UserPattern
	for userID, acts := range userActions {
  // Custom sort
		sort.Slice(acts, func(i, j int) bool {
			return acts[i].day < acts[j].day
		})

		// Find consecutive days with same action
		i := 0
		for i < len(acts) {
			j := i
			for j+1 < len(acts) &&
				acts[j+1].day == acts[j].day+1 &&
				acts[j+1].action == acts[j].action {
				j++
			}
			length := j - i + 1
			if length >= minLength {
				result = append(result, UserPattern{
					UserID:        userID,
					Pattern:       acts[i].action,
					PatternLength: length,
				})
			}
			i = j + 1
		}
	}

	// Sort by user, then pattern
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].UserID != result[j].UserID {
			return result[i].UserID < result[j].UserID
		}
		return result[i].Pattern < result[j].Pattern
	})

	return result
}
```
