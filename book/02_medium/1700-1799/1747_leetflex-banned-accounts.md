# 1747 — Leetflex Banned Accounts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func findBanned(logins []Login) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Login struct {
	AccountID int
	IPAddress string
	LoginTime int
}

func findBanned(logins []Login) []int {
	// Group by account
  // HashMap: O(1) lookup
	groups := make(map[int][]Login)
	for _, l := range logins {
		groups[l.AccountID] = append(groups[l.AccountID], l)
	}

  // HashMap: O(1) lookup
	banned := make(map[int]bool)
	for accID, records := range groups {
  // Custom sort
		sort.Slice(records, func(i, j int) bool {
			return records[i].LoginTime < records[j].LoginTime
		})
		// Track latest login time per IP for this account
  // HashMap: O(1) lookup
		lastTime := make(map[string]int)
		for _, r := range records {
			if prevTime, ok := lastTime[r.IPAddress]; ok {
				// Same IP, update last time
				_ = prevTime
			}
			lastTime[r.IPAddress] = r.LoginTime
		}

		// Check if any IP has concurrent sessions
  // HashMap: O(1) lookup
		active := make(map[string]int)
		for _, r := range records {
			if _, ok := active[r.IPAddress]; ok {
				// Check if there's a different IP with an active session
				for ip, t := range active {
					if ip != r.IPAddress && t < r.LoginTime {
						banned[accID] = true
					}
				}
				// End current session for this IP and start new one
				delete(active, r.IPAddress)
			}
			active[r.IPAddress] = r.LoginTime
		}
	}

  // Alokasi slice
	result := make([]int, 0, len(banned))
	for id := range banned {
		result = append(result, id)
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	logins := []Login{
		{1, "IP1", 1},
		{1, "IP2", 2},
		{1, "IP1", 3},
	}
	fmt.Println(findBanned(logins))
}
```
