# 2688 — Find Active Users

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findActiveUsers(purchases []UserPurchase) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2688: Find Active Users
// https://leetcode.com/problems/find-active-users/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: Find users who made a second purchase within 7 days
// (inclusive) of another purchase. The 7-day window is inclusive of
// both start and end dates.

import (
	"fmt"
	"sort"
	"time"
)

// UserPurchase represents the Users database table.
type UserPurchase struct {
	UserID    int
	Item      string
	CreatedAt string // format: "YYYY-MM-DD"
	Amount    int
}

// findActiveUsers simulates the SQL query.
// Time: O(n log n) | Space: O(n)
// n = number of purchase records.
func findActiveUsers(purchases []UserPurchase) []int {
	// Group purchases by user_id.
  // Membuat map (HashMap) — pencarian O(1)
	userDates := make(map[int][]string)
	for _, p := range purchases {
		userDates[p.UserID] = append(userDates[p.UserID], p.CreatedAt)
	}

	// Count distinct dates per user.
  // Membuat map (HashMap) — pencarian O(1)
	userUniqueDates := make(map[int][]string)
	for uid, dates := range userDates {
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[string]bool)
		for _, d := range dates {
			if !seen[d] {
				seen[d] = true
				userUniqueDates[uid] = append(userUniqueDates[uid], d)
			}
		}
	}

	var active []int

	for uid, dates := range userUniqueDates {
		// Sort dates ascending.
		sort.Strings(dates)

		// Check consecutive purchases for <= 7 day gap.
		for i := 1; i < len(dates); i++ {
			prev, err1 := time.Parse("2006-01-02", dates[i-1])
			curr, err2 := time.Parse("2006-01-02", dates[i])
			if err1 != nil || err2 != nil {
				continue
			}
			diff := curr.Sub(prev).Hours() / 24
			if diff >= 0 && diff <= 7 {
				active = append(active, uid)
				break
			}
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(active)
	return active
}

func main() {
	// Test data from the problem description.
	purchases := []UserPurchase{
		// User 6 has purchases on Sep 10 and Sep 14 (4 days apart) -> active.
		{UserID: 6, Item: "item1", CreatedAt: "2023-09-10", Amount: 100},
		{UserID: 6, Item: "item2", CreatedAt: "2023-09-14", Amount: 200},
		// User 1 has purchases far apart (Sep 10 and Sep 20, 10 days) -> not active.
		{UserID: 1, Item: "item3", CreatedAt: "2023-09-10", Amount: 50},
		{UserID: 1, Item: "item4", CreatedAt: "2023-09-20", Amount: 75},
		// User 2 has only 1 purchase -> not active.
		{UserID: 2, Item: "item5", CreatedAt: "2023-09-10", Amount: 150},
		// User 3 has purchases exactly 7 days apart -> active.
		{UserID: 3, Item: "item6", CreatedAt: "2023-09-10", Amount: 60},
		{UserID: 3, Item: "item7", CreatedAt: "2023-09-17", Amount: 80},
	}

	results := findActiveUsers(purchases)

	fmt.Println("Active Users (user_id):")
	for _, uid := range results {
		fmt.Println(uid)
	}
	// Expected output:
	// 3
	// 6
}
```
