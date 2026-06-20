# 2995 — Viewers Turned Streamers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func viewersTurnedStreamers(sessions []Session) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2995: Viewers Turned Streamers (SQL simulation)
// https://leetcode.com/problems/viewers-turned-streamers/
// Difficulty: Hard [Paid]
//
// Find users whose first session was as a "viewer" and who later
// had a "streamer" session. Return each such user with their first
// streamer session date, sorted by user_id ascending.

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
	// Step 1: Find each user's first session
	type firstSeen struct {
		date     string
		sessType string
	}
  // Membuat map (HashMap) — pencarian O(1)
	userFirst := make(map[int]firstSeen)
	for _, s := range sessions {
		if existing, ok := userFirst[s.UserID]; !ok || s.SessionDate < existing.date {
			userFirst[s.UserID] = firstSeen{s.SessionDate, s.SessionType}
		}
	}

	// Step 2: For users whose first session was "viewer", find first "streamer" session
	type conversion struct {
		userID int
		date   string
	}
	var converted []conversion

	for uid, f := range userFirst {
		if f.sessType != "viewer" {
			continue
		}
		firstStreamerDate := ""
		for _, s := range sessions {
			if s.UserID == uid && s.SessionType == "streamer" {
				if firstStreamerDate == "" || s.SessionDate < firstStreamerDate {
					firstStreamerDate = s.SessionDate
				}
			}
		}
		if firstStreamerDate != "" {
			converted = append(converted, conversion{uid, firstStreamerDate})
		}
	}

  // Custom sort dengan comparator
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
	// Test 1: Basic conversion
	sessions := []Session{
		{1, "viewer", "2023-01-01"},
		{1, "streamer", "2023-02-01"},
		{2, "viewer", "2023-01-15"},
		{3, "streamer", "2023-01-10"},
	}
	fmt.Println("Test 1:")
	for _, r := range viewersTurnedStreamers(sessions) {
		fmt.Println(r)
	}

	// Test 2: First session is streamer (should not be included)
	fmt.Println("\nTest 2 (first session streamer):")
	sessions2 := []Session{
		{1, "streamer", "2023-01-01"},
		{1, "viewer", "2023-02-01"},
		{2, "viewer", "2023-01-15"},
		{2, "streamer", "2023-01-20"},
	}
	for _, r := range viewersTurnedStreamers(sessions2) {
		fmt.Println(r)
	}

	// Test 3: Multiple viewer sessions before first streamer
	fmt.Println("\nTest 3 (multiple viewers before streamer):")
	sessions3 := []Session{
		{1, "viewer", "2023-01-01"},
		{1, "viewer", "2023-01-10"},
		{1, "streamer", "2023-02-01"},
	}
	for _, r := range viewersTurnedStreamers(sessions3) {
		fmt.Println(r)
	}

	// Test 4: Viewer only, never converted
	fmt.Println("\nTest 4 (never converted):")
	sessions4 := []Session{
		{1, "viewer", "2023-01-01"},
	}
	fmt.Println(viewersTurnedStreamers(sessions4))
}
```
