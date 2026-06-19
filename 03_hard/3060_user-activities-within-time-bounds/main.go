package main

// LeetCode #3060: User Activities Within Time Bounds (SQL simulation)
// https://leetcode.com/problems/user-activities-within-time-bounds/
// Difficulty: Hard [Paid]

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
	byUser := make(map[int][]UserSession)
	for _, s := range sessions {
		byUser[s.UserID] = append(byUser[s.UserID], s)
	}
	var result []int
	for uid, sList := range byUser {
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
			if found { break }
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
	parse := func(s string) time.Time { t, _ := time.Parse(layout, s); return t }
	sessions := []UserSession{
		{101, parse("2023-01-01 08:00:00"), parse("2023-01-01 10:00:00"), 1, "Viewer"},
		{102, parse("2023-01-01 09:00:00"), parse("2023-01-01 11:00:00"), 3, "Viewer"},
		{102, parse("2023-01-01 13:00:00"), parse("2023-01-01 14:00:00"), 4, "Viewer"},
	}
	fmt.Println(userActivitiesWithinTimeBounds(sessions))
}
