# 0262 — Trips And Users

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func roundTo2(f float64) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(T + U + D log D) where T=#trips, U=#users, D=#distinct dates  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #262: Trips and Users
// https://leetcode.com/problems/trips-and-users/
// Difficulty: Hard
//
// Find the cancellation rate of unbanned users (both client and driver must not be banned)
// for each day between "2013-10-01" and "2013-10-03".
// Cancellation rate = number of cancelled trips / total trips (by unbanned users).

// User represents a row in the Users table.
type User struct {
	UsersID int
	Banned  string // "Yes" or "No"
	Role    string // "client", "driver", "partner"
}

// Trip represents a row in the Trips table.
type Trip struct {
	ID        int
	ClientID  int
	DriverID  int
	CityID    int
	Status    string // "completed", "cancelled_by_driver", "cancelled_by_client"
	RequestAt string // date "YYYY-MM-DD"
}

// DailyRate holds one output row.
type DailyRate struct {
	Day    string
	Rate   float64 // cancellation rate, rounded to 2 decimal places
}

// RoundTo2 rounds f to 2 decimal places.
func roundTo2(f float64) float64 {
	return float64(int(f*100+0.5)) / 100.0
}

// tripsAndUsers computes the daily cancellation rate for unbanned users.
// Time: O(T + U + D log D) where T=#trips, U=#users, D=#distinct dates
func tripsAndUsers(trips []Trip, users []User) []DailyRate {
	// Build banned user set.
  // HashMap: O(1) lookup
	banned := make(map[int]bool)
	for _, u := range users {
		if u.Banned == "Yes" {
			banned[u.UsersID] = true
		}
	}

	// Group trips by date, only including trips where both client and driver are unbanned.
	type dateStats struct {
		total       int
		cancelled   int
	}
  // HashMap: O(1) lookup
	byDay := make(map[string]*dateStats)

	for _, t := range trips {
		if banned[t.ClientID] || banned[t.DriverID] {
			continue
		}
		if _, ok := byDay[t.RequestAt]; !ok {
			byDay[t.RequestAt] = &dateStats{}
		}
		byDay[t.RequestAt].total++
		if t.Status != "completed" {
			byDay[t.RequestAt].cancelled++
		}
	}

	// Collect dates in sorted order.
	var dates []string
	for d := range byDay {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	var result []DailyRate
	for _, d := range dates {
		s := byDay[d]
		rate := 0.0
		if s.total > 0 {
			rate = roundTo2(float64(s.cancelled) / float64(s.total))
		}
		result = append(result, DailyRate{Day: d, Rate: rate})
	}
	return result
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0262 Trips and Users ===")

	users := []User{
		{UsersID: 1, Banned: "No", Role: "client"},
		{UsersID: 2, Banned: "Yes", Role: "client"},
		{UsersID: 3, Banned: "No", Role: "client"},
		{UsersID: 4, Banned: "No", Role: "client"},
		{UsersID: 10, Banned: "No", Role: "driver"},
		{UsersID: 11, Banned: "No", Role: "driver"},
		{UsersID: 12, Banned: "No", Role: "driver"},
		{UsersID: 13, Banned: "No", Role: "driver"},
	}

	trips := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-01"},
		{ID: 2, ClientID: 2, DriverID: 11, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-01"},
		{ID: 3, ClientID: 3, DriverID: 12, CityID: 6, Status: "completed", RequestAt: "2013-10-01"},
		{ID: 4, ClientID: 4, DriverID: 13, CityID: 6, Status: "cancelled_by_client", RequestAt: "2013-10-01"},
		{ID: 5, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 6, ClientID: 2, DriverID: 11, CityID: 6, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 7, ClientID: 3, DriverID: 12, CityID: 6, Status: "completed", RequestAt: "2013-10-02"},
		{ID: 8, ClientID: 2, DriverID: 12, CityID: 12, Status: "completed", RequestAt: "2013-10-03"},
		{ID: 9, ClientID: 3, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-03"},
		{ID: 10, ClientID: 4, DriverID: 13, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-03"},
	}

	fmt.Println("Trips:", len(trips), "Users:", len(users))
	results := tripsAndUsers(trips, users)
	fmt.Println("\nDaily Cancellation Rate (unbanned users only):")
	for _, r := range results {
		fmt.Printf("  %s  rate=%.2f\n", r.Day, r.Rate)
	}

	// --- Edge cases ---
	fmt.Println("\n--- Edge Cases ---")

	// No trips.
	empty := tripsAndUsers(nil, users)
	fmt.Println("Empty trips:", len(empty))

	// All trips cancelled.
	trips2 := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "cancelled_by_client", RequestAt: "2013-10-01"},
		{ID: 2, ClientID: 3, DriverID: 11, CityID: 1, Status: "cancelled_by_driver", RequestAt: "2013-10-01"},
	}
	r2 := tripsAndUsers(trips2, users)
	for _, r := range r2 {
		fmt.Printf("  All cancelled: %s rate=%.2f (expected 1.00)\n", r.Day, r.Rate)
	}

	// All banned users (trips filtered out).
	usersBanned := []User{
		{UsersID: 1, Banned: "Yes", Role: "client"},
		{UsersID: 10, Banned: "Yes", Role: "driver"},
	}
	trips3 := []Trip{
		{ID: 1, ClientID: 1, DriverID: 10, CityID: 1, Status: "completed", RequestAt: "2013-10-01"},
	}
	r3 := tripsAndUsers(trips3, usersBanned)
	fmt.Println("All banned users:", len(r3), "(expected 0)")
}
```
