package main

// LeetCode #3673: Find Zombie Sessions
// https://leetcode.com/problems/find-zombie-sessions/
// Difficulty: Hard (SQL / Database)
//
// Find sessions that are still active (zombie) — sessions with no
// logout event recorded after the login. Port to Go with data
// structures.

import "fmt"

type Session struct {
	SessionID int
	UserID    int
	LoginTime int
	Logout    bool // true if logout recorded
}

func main() {
	// Example 1
	sessions := []Session{
		{1, 1, 100, false},
		{2, 1, 200, true},
		{3, 2, 150, false},
	}
	fmt.Println(findZombieSessions(sessions))

	// Example 2: all logged out
	sessions2 := []Session{
		{1, 1, 100, true},
		{2, 2, 200, true},
	}
	fmt.Println(findZombieSessions(sessions2))
}

func findZombieSessions(sessions []Session) []Session {
	var zombie []Session
	for _, s := range sessions {
		if !s.Logout {
			zombie = append(zombie, s)
		}
	}

	// Sort by SessionID for determinism
	for i := 0; i < len(zombie); i++ {
		for j := i + 1; j < len(zombie); j++ {
			if zombie[i].SessionID > zombie[j].SessionID {
				zombie[i], zombie[j] = zombie[j], zombie[i]
			}
		}
	}
	return zombie
}
