# 1127 — User Purchase Platform

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func getUserPurchasePlatform(spending []UserSpend) []PlatformStats
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1127: User Purchase Platform
// https://leetcode.com/problems/user-purchase-platform/
// Difficulty: Hard [Paid]
//
// Given a table of user spending with platform (mobile, desktop) and date,
// find for each date the total amount and number of users broken down by
// platform category: 'mobile' (mobile only), 'desktop' (desktop only),
// 'both' (both platforms).

import (
	"fmt"
	"sort"
)

// UserSpend represents a single purchase record.
type UserSpend struct {
	UserID    int
	Platform  string // "mobile" or "desktop"
	Amount    float64
	Date      string // "YYYY-MM-DD"
}

// PlatformStats holds the output per date per platform category.
type PlatformStats struct {
	Date         string
	Platform     string // "mobile", "desktop", or "both"
	TotalAmount  float64
	TotalUsers   int
}

func main() {
	// Test case
	spending := []UserSpend{
		{1, "mobile", 100, "2019-07-01"},
		{1, "desktop", 100, "2019-07-01"},
		{2, "mobile", 100, "2019-07-01"},
		{2, "desktop", 100, "2019-07-01"},
		{3, "mobile", 100, "2019-07-01"},
	}

	results := getUserPurchasePlatform(spending)
	for _, r := range results {
		fmt.Printf("%s | %s | total_amount=%.0f | total_users=%d\n", r.Date, r.Platform, r.TotalAmount, r.TotalUsers)
	}
	fmt.Println("---")

	// Single platform only
	spending = []UserSpend{
		{1, "mobile", 50, "2019-08-01"},
		{2, "mobile", 75, "2019-08-01"},
	}
	results = getUserPurchasePlatform(spending)
	for _, r := range results {
		fmt.Printf("%s | %s | total_amount=%.0f | total_users=%d\n", r.Date, r.Platform, r.TotalAmount, r.TotalUsers)
	}
}

// getUserPurchasePlatform computes, per date, the breakdown by platform category.
func getUserPurchasePlatform(spending []UserSpend) []PlatformStats {
	// Group user-platform by date
	type spentInfo struct {
		total float64
		count int
	}

	// userPlatformSet[date][userID][platform] = amount
  // Membuat map (HashMap) — pencarian O(1)
	dateUserPlatform := make(map[string]map[int]map[string]float64)

	for _, s := range spending {
		if dateUserPlatform[s.Date] == nil {
			dateUserPlatform[s.Date] = make(map[int]map[string]float64)
		}
		if dateUserPlatform[s.Date][s.UserID] == nil {
			dateUserPlatform[s.Date][s.UserID] = make(map[string]float64)
		}
		dateUserPlatform[s.Date][s.UserID][s.Platform] += s.Amount
	}

	dates := make([]string, 0, len(dateUserPlatform))
	for d := range dateUserPlatform {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	var results []PlatformStats

	for _, date := range dates {
		users := dateUserPlatform[date]

		mobileTot := 0.0
		mobileUsers := 0
		desktopTot := 0.0
		desktopUsers := 0
		bothTot := 0.0
		bothUsers := 0

		for _, platforms := range users {
			_, hasMobile := platforms["mobile"]
			_, hasDesktop := platforms["desktop"]

			if hasMobile && hasDesktop {
				bothTot += platforms["mobile"] + platforms["desktop"]
				bothUsers++
			} else if hasMobile {
				mobileTot += platforms["mobile"]
				mobileUsers++
			} else if hasDesktop {
				desktopTot += platforms["desktop"]
				desktopUsers++
			}
		}

		if mobileUsers > 0 {
			results = append(results, PlatformStats{date, "mobile", mobileTot, mobileUsers})
		}
		if desktopUsers > 0 {
			results = append(results, PlatformStats{date, "desktop", desktopTot, desktopUsers})
		}
		if bothUsers > 0 {
			results = append(results, PlatformStats{date, "both", bothTot, bothUsers})
		}
	}

	return results
}
```
