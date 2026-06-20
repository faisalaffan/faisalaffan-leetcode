# 3832 — Find Users With Persistent Behavior Patterns

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findPersistentPatterns(actions []UserAction, minLength int) []UserPattern
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
  // Custom sort dengan comparator
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
  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		if result[i].UserID != result[j].UserID {
			return result[i].UserID < result[j].UserID
		}
		return result[i].Pattern < result[j].Pattern
	})

	return result
}
```
