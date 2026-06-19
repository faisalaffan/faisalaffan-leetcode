package main

// LeetCode #2995: Viewers Turned Streamers (SQL simulation)
// https://leetcode.com/problems/viewers-turned-streamers/
// Difficulty: Hard [Paid]

import (
	"fmt"
	"sort"
)

type Session struct {
	UserID      int
	SessionType string
	SessionDate string
}

func viewersTurnedStreamers(sessions []Session) []string {
	type firstSeen struct {
		date     string
		sessType string
	}
	userFirst := make(map[int]firstSeen)
	for _, s := range sessions {
		if existing, ok := userFirst[s.UserID]; !ok || s.SessionDate < existing.date {
			userFirst[s.UserID] = firstSeen{s.SessionDate, s.SessionType}
		}
	}
	type conversion struct {
		userID int
		date   string
	}
	var converted []conversion
	for uid, f := range userFirst {
		if f.sessType == "viewer" {
			for _, s := range sessions {
				if s.UserID == uid && s.SessionType == "streamer" {
					converted = append(converted, conversion{uid, s.SessionDate})
					break
				}
			}
		}
	}
	sort.Slice(converted, func(i, j int) bool {
		return converted[i].userID < converted[j].userID
	})
	var result []string
	for _, c := range converted {
		result = append(result, fmt.Sprintf("%d|%s", c.userID, c.date))
	}
	return result
}

func main() {
	sessions := []Session{
		{1, "viewer", "2023-01-01"},
		{1, "streamer", "2023-02-01"},
		{2, "viewer", "2023-01-15"},
		{3, "streamer", "2023-01-10"},
	}
	for _, r := range viewersTurnedStreamers(sessions) {
		fmt.Println(r)
	}
}
