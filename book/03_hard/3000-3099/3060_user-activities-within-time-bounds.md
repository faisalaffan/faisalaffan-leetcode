# 3060 — User Activities Within Time Bounds

## Deskripsi

**Soal:** [3060. User Activities Within Time Bounds](https://leetcode.com/problems/user-activities-within-time-bounds/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func userActivitiesWithinTimeBounds(sessions []UserSession) []int`

> **Ide Kunci:** Find users who have two consecutive sessions of the same type

## Solusi Go

```go
package main

// LeetCode #3060: User Activities Within Time Bounds (SQL simulation)
// https://leetcode.com/problems/user-activities-within-time-bounds/
// Difficulty: Hard [Paid]
//
// Approach: Find users who have two consecutive sessions of the same type
// within 12 hours of each other (end of first to start of second).

import (
	"fmt"
	"sort"
	"time"
)

type UserSession struct {
	UserID       int
	SessionStart time.Time
	SessionEnd   time.Time
	SessionID    int
	SessionType  string
}

func userActivitiesWithinTimeBounds(sessions []UserSession) []int {
  // Membuat map untuk pencarian O(1): key → value
	byUser := make(map[int][]UserSession)
	for _, s := range sessions {
		byUser[s.UserID] = append(byUser[s.UserID], s)
	}
	var result []int
	for uid, sList := range byUser {
  // Membuat map untuk pencarian O(1): key → value
		byType := make(map[string][]UserSession)
		for _, s := range sList {
			byType[s.SessionType] = append(byType[s.SessionType], s)
		}
		found := false
		for _, typedSessions := range byType {
			sort.Slice(typedSessions, func(i, j int) bool {
				return typedSessions[i].SessionStart.Before(typedSessions[j].SessionStart)
			})
			for i := 1; i < len(typedSessions); i++ {
				gap := typedSessions[i].SessionStart.Sub(typedSessions[i-1].SessionEnd)
				if gap <= 12*time.Hour && gap >= 0 {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			result = append(result, uid)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	layout := "2006-01-02 15:04:05"
	parse := func(s string) time.Time {
		t, _ := time.Parse(layout, s)
		return t
	}

	// Example 1: User 102 has two Viewer sessions 2h apart (10:00 to 13:00)
	sessions1 := []UserSession{
		{101, parse("2023-01-01 08:00:00"), parse("2023-01-01 10:00:00"), 1, "Viewer"},
		{102, parse("2023-01-01 09:00:00"), parse("2023-01-01 11:00:00"), 3, "Viewer"},
		{102, parse("2023-01-01 13:00:00"), parse("2023-01-01 14:00:00"), 4, "Viewer"},
	}
	fmt.Println("Example 1:", userActivitiesWithinTimeBounds(sessions1))
	// Expected: [102] (gap from 11:00 to 13:00 = 2h <= 12h)

	// Example 2: gap exactly 12 hours
	sessions2 := []UserSession{
		{201, parse("2023-01-01 08:00:00"), parse("2023-01-01 09:00:00"), 1, "TypeA"},
		{201, parse("2023-01-01 21:00:00"), parse("2023-01-01 22:00:00"), 2, "TypeA"},
	}
	fmt.Println("Example 2 (exactly 12h):", userActivitiesWithinTimeBounds(sessions2))
	// Expected: [201] (gap = 12h exactly)

	// Example 3: gap > 12 hours
	sessions3 := []UserSession{
		{301, parse("2023-01-01 08:00:00"), parse("2023-01-01 09:00:00"), 1, "TypeA"},
		{301, parse("2023-01-01 22:00:00"), parse("2023-01-01 23:00:00"), 2, "TypeA"},
	}
	fmt.Println("Example 3 (gap > 12h):", userActivitiesWithinTimeBounds(sessions3))
	// Expected: [] (gap = 13h > 12h)

	// Example 4: different type sessions, gap within bound
	sessions4 := []UserSession{
		{401, parse("2023-01-01 08:00:00"), parse("2023-01-01 09:00:00"), 1, "TypeA"},
		{401, parse("2023-01-01 10:00:00"), parse("2023-01-01 11:00:00"), 2, "TypeB"},
	}
	fmt.Println("Example 4 (diff types):", userActivitiesWithinTimeBounds(sessions4))
	// Expected: [] (different types, so grouped separately, each group has only 1 session)

	// Example 5: multiple users
	sessions5 := []UserSession{
		{1, parse("2023-01-01 08:00:00"), parse("2023-01-01 09:00:00"), 1, "A"},
		{2, parse("2023-01-01 10:00:00"), parse("2023-01-01 11:00:00"), 2, "A"},
		{1, parse("2023-01-01 12:00:00"), parse("2023-01-01 13:00:00"), 3, "A"},
		{2, parse("2023-01-01 14:00:00"), parse("2023-01-01 15:00:00"), 4, "B"},
	}
	fmt.Println("Example 5 (multiple):", userActivitiesWithinTimeBounds(sessions5))
}
```
