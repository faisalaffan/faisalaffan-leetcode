# 1097 — Game Play Analysis V

## Deskripsi

**Soal:** [1097. Game Play Analysis V](https://leetcode.com/problems/game-play-analysis-v/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** O(N log N) for sorting, Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

**Fungsi Solusi:** `func roundTo2(f float64) float64`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1097: Game Play Analysis V
// https://leetcode.com/problems/game-play-analysis-v/
// Difficulty: Hard [Paid]
//
// For each install date (first login date of each player), compute:
// - Number of players who installed on that date
// - Number of those players who logged in again on the next day (day-1 retention)
// - Retention rate = day1_retention / installs, rounded to 2 decimal places

// Activity represents a row in the Activity table.
type Activity struct {
	PlayerID   int
	DeviceID   int
	EventDate  string // "YYYY-MM-DD"
	GamesPlayed int
}

// InstallRetention holds one result row.
type InstallRetention struct {
	InstallDate   string
	Installs      int
	Day1Retention int
	RetentionRate float64
}

// roundTo2 rounds to 2 decimal places.
func roundTo2(f float64) float64 {
	return float64(int(f*100+0.5)) / 100.0
}

// gamePlayAnalysisV computes install/retention stats per install date.
// Time: O(N log N) for sorting, Space: O(N)
func gamePlayAnalysisV(activities []Activity) []InstallRetention {
	if len(activities) == 0 {
		return nil
	}

	// Find first login date for each player.
  // Membuat map untuk pencarian O(1): key → value
	firstLogin := make(map[int]string) // player_id -> install date
	for _, a := range activities {
		if existing, ok := firstLogin[a.PlayerID]; !ok || a.EventDate < existing {
			firstLogin[a.PlayerID] = a.EventDate
		}
	}

	// Build player -> set of login dates for quick lookup.
  // Membuat map untuk pencarian O(1): key → value
	playerLogins := make(map[int]map[string]bool)
	for _, a := range activities {
		if playerLogins[a.PlayerID] == nil {
			playerLogins[a.PlayerID] = make(map[string]bool)
		}
		playerLogins[a.PlayerID][a.EventDate] = true
	}

	// nextDay returns the date after the given date (simple implementation).
	nextDay := func(date string) string {
		// Parse "YYYY-MM-DD".
		var y, m, d int
		fmt.Sscanf(date, "%d-%d-%d", &y, &m, &d)

		// Days per month (non-leap year for simplicity).
		daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		// Check leap year.
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			daysInMonth[2] = 29
		}

		d++
		if d > daysInMonth[m] {
			d = 1
			m++
			if m > 12 {
				m = 1
				y++
			}
		}
		return fmt.Sprintf("%d-%02d-%02d", y, m, d)
	}

	// Group by install date.
  // Membuat map untuk pencarian O(1): key → value
	installGroups := make(map[string][]int) // install date -> player IDs
	for playerID, installDate := range firstLogin {
		installGroups[installDate] = append(installGroups[installDate], playerID)
	}

	// Sort install dates.
	var dates []string
	for d := range installGroups {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	var results []InstallRetention

	for _, installDate := range dates {
		players := installGroups[installDate]
		installs := len(players)
		day1Retention := 0

		dayAfter := nextDay(installDate)

		for _, pid := range players {
			if playerLogins[pid][dayAfter] {
				day1Retention++
			}
		}

		rate := 0.0
		if installs > 0 {
			rate = roundTo2(float64(day1Retention) / float64(installs))
		}

		results = append(results, InstallRetention{
			InstallDate:   installDate,
			Installs:      installs,
			Day1Retention: day1Retention,
			RetentionRate: rate,
		})
	}

	return results
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 1097 Game Play Analysis V ===")

	activities := []Activity{
		{PlayerID: 1, EventDate: "2016-03-01", GamesPlayed: 5},
		{PlayerID: 1, EventDate: "2016-03-02", GamesPlayed: 6},
		{PlayerID: 2, EventDate: "2016-03-01", GamesPlayed: 2},
		{PlayerID: 2, EventDate: "2016-03-02", GamesPlayed: 3},
		{PlayerID: 3, EventDate: "2017-06-25", GamesPlayed: 1},
		{PlayerID: 3, EventDate: "2017-06-26", GamesPlayed: 1},
		{PlayerID: 4, EventDate: "2016-03-01", GamesPlayed: 7},
		// Player 5: installs 2017-06-25 but does NOT come back the next day.
		{PlayerID: 5, EventDate: "2017-06-25", GamesPlayed: 1},
		// Player 6: installs 2016-03-03, no next day login.
		{PlayerID: 6, EventDate: "2016-03-03", GamesPlayed: 4},
	}

	results := gamePlayAnalysisV(activities)
	fmt.Println("Installation Retention Analysis:")
	for _, r := range results {
		fmt.Printf("  Install Date: %s | Installs: %d | Day-1 Retention: %d | Rate: %.2f\n",
			r.InstallDate, r.Installs, r.Day1Retention, r.RetentionRate)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// Single player, no retention.
	acts2 := []Activity{
		{PlayerID: 1, EventDate: "2020-01-01", GamesPlayed: 1},
	}
	r2 := gamePlayAnalysisV(acts2)
	fmt.Println("Single player, no next-day login:")
	for _, r := range r2 {
		fmt.Printf("  %s | installs=%d ret=%d rate=%.2f\n", r.InstallDate, r.Installs, r.Day1Retention, r.RetentionRate)
	}

	// Single player, with retention.
	acts3 := []Activity{
		{PlayerID: 10, EventDate: "2020-01-01", GamesPlayed: 1},
		{PlayerID: 10, EventDate: "2020-01-02", GamesPlayed: 2},
		{PlayerID: 10, EventDate: "2020-01-03", GamesPlayed: 3},
	}
	r3 := gamePlayAnalysisV(acts3)
	fmt.Println("Single player, with next-day login:")
	for _, r := range r3 {
		fmt.Printf("  %s | installs=%d ret=%d rate=%.2f\n", r.InstallDate, r.Installs, r.Day1Retention, r.RetentionRate)
	}

	// Multiple install dates.
	acts4 := []Activity{
		{PlayerID: 1, EventDate: "2020-01-01", GamesPlayed: 1},
		{PlayerID: 1, EventDate: "2020-01-02", GamesPlayed: 1},
		{PlayerID: 2, EventDate: "2020-01-01", GamesPlayed: 1},
		{PlayerID: 3, EventDate: "2020-01-02", GamesPlayed: 1},
		{PlayerID: 3, EventDate: "2020-01-03", GamesPlayed: 1},
	}
	r4 := gamePlayAnalysisV(acts4)
	fmt.Println("Multiple install dates:")
	for _, r := range r4 {
		fmt.Printf("  %s | installs=%d ret=%d rate=%.2f\n", r.InstallDate, r.Installs, r.Day1Retention, r.RetentionRate)
	}

	// Empty.
	fmt.Println("Empty:", gamePlayAnalysisV(nil))
}
```
