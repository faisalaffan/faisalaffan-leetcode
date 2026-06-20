# 1454 — Active Users

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func activeUsers(accounts []struct { id int name string }, logins []struct { userID int loginDate string }) []activeResult`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n) for sorting logins  |  **Ruang:** O(n) for maps

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1454: Active Users
// https://leetcode.com/problems/active-users/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	accounts := []struct {
		id   int
		name string
	}{
		{1, "Daniel"},
		{2, "Bob"},
		{3, "John"},
		{4, "Alice"},
	}
	logins := []struct {
		userID    int
		loginDate string
	}{
		{1, "2020-01-01"}, {1, "2020-01-02"}, {1, "2020-01-03"},
		{1, "2020-01-06"}, {1, "2020-01-07"},
		{2, "2020-01-01"}, {2, "2020-01-02"}, {2, "2020-01-03"},
		{2, "2020-01-05"}, {2, "2020-01-06"}, {2, "2020-01-07"},
		{3, "2020-01-05"}, {3, "2020-01-06"}, {3, "2020-01-07"},
		{3, "2020-01-08"},
		{4, "2020-01-01"},
	}

	result := activeUsers(accounts, logins)
	for _, r := range result {
		fmt.Printf("%d %s\n", r.id, r.name)
	}
}

type activeResult struct {
	id   int
	name string
}

// Time: O(n log n) for sorting logins
// Space: O(n) for maps
func activeUsers(accounts []struct {
	id   int
	name string
}, logins []struct {
	userID    int
	loginDate string
}) []activeResult {
	// Group logins by user
  // HashMap: O(1) lookup
	userLogins := make(map[int][]string)
	for _, l := range logins {
		userLogins[l.userID] = append(userLogins[l.userID], l.loginDate)
	}

	// Sort logins for each user and check for 5+ consecutive days
  // HashMap: O(1) lookup
	activeIDs := make(map[int]bool)
	for userID, dates := range userLogins {
		sort.Strings(dates)

		// Remove duplicates
		unique := make([]string, 0, len(dates))
  // HashMap: O(1) lookup
		seen := make(map[string]bool)
		for _, d := range dates {
			if !seen[d] {
				seen[d] = true
				unique = append(unique, d)
			}
		}

		// Check for 5 consecutive days
		consecutive := 1
		for i := 1; i < len(unique); i++ {
			if isNextDay(unique[i-1], unique[i]) {
				consecutive++
				if consecutive >= 5 {
					activeIDs[userID] = true
					break
				}
			} else {
				consecutive = 1
			}
		}
	}

	var result []activeResult
	for _, a := range accounts {
		if activeIDs[a.id] {
			result = append(result, activeResult{a.id, a.name})
		}
	}

	return result
}

func isNextDay(day1, day2 string) bool {
	// Simple date comparison (YYYY-MM-DD format)
	y1, m1, d1 := parseDate(day1)
	y2, m2, d2 := parseDate(day2)

	// Check if day2 is the next day after day1
	if y1 == y2 && m1 == m2 && d2-d1 == 1 {
		return true
	}
	// Handle month/year boundaries (simplified)
	daysInMonth := map[int]int{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30,
		7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}
	if y1 == y2 && m2 == m1+1 && d1 == daysInMonth[m1] && d2 == 1 {
		return true
	}
	if y2 == y1+1 && m1 == 12 && m2 == 1 && d1 == 31 && d2 == 1 {
		return true
	}
	return false
}

func parseDate(date string) (int, int, int) {
	var y, m, d int
	fmt.Sscanf(date, "%d-%d-%d", &y, &m, &d)
	return y, m, d
}
```
