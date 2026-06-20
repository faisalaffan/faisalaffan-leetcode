# 1972 — First And Last Call On The Same Day

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func firstAndLastCallOnTheSameDay(calls [][]interface{}) [][]interface`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1972: First and Last Call On the Same Day
// https://leetcode.com/problems/first-and-last-call-on-the-same-day/
// Difficulty: Hard [Paid]
//
// Given a table calls (caller_id, recipient_id, call_time) with call_time
// as a datetime, find all users whose first and last call on the same day
// was with the same person. For each such user, return user_id, the other
// participant's id, and the date.
//
// This is a SQL-style problem implemented in Go.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	calls := [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 2, "2021-01-01 17:00:00"},
	}
	// User 1: first call 9am with 2, last call 5pm with 2 -> same person (2)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 2: different person for first vs last
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 3, "2021-01-01 17:00:00"},
	}
	// User 1: first with 2, last with 3 -> different, no result
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 3: multiple users
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 2, "2021-01-01 17:00:00"},
		{2, 1, "2021-01-01 08:00:00"},
		{2, 1, "2021-01-01 18:00:00"},
	}
	// User 1: first with 2, last with 2 -> (1, 2, 2021-01-01)
	// User 2: first with 1, last with 1 -> (2, 1, 2021-01-01)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 4: multiple days
	calls = [][]interface{}{
		{1, 2, "2021-01-01 09:00:00"},
		{1, 3, "2021-01-01 17:00:00"},
		{1, 2, "2021-01-02 09:00:00"},
		{1, 2, "2021-01-02 17:00:00"},
	}
	// User 1 on 2021-01-02: first with 2, last with 2 -> (1, 2, 2021-01-02)
	fmt.Println(firstAndLastCallOnTheSameDay(calls))

	// Test case 5: empty
	fmt.Println(firstAndLastCallOnTheSameDay([][]interface{}{}))
}

func firstAndLastCallOnTheSameDay(calls [][]interface{}) [][]interface{} {
	// Group calls by (user, date)
	// Each call involves two users: caller and recipient
	type callInfo struct {
		other int
		time  string
	}
  // HashMap: O(1) lookup
	userDayCalls := make(map[[2]string][]callInfo) // (user_id, date) -> calls

	for _, c := range calls {
		caller := c[0].(int)
		recipient := c[1].(int)
		timeStr := c[2].(string)
		date := timeStr[:10] // YYYY-MM-DD

		// Add for caller
		key1 := [2]string{fmt.Sprintf("%d", caller), date}
		userDayCalls[key1] = append(userDayCalls[key1], callInfo{recipient, timeStr})

		// Add for recipient (they also participated in the call)
		key2 := [2]string{fmt.Sprintf("%d", recipient), date}
		userDayCalls[key2] = append(userDayCalls[key2], callInfo{caller, timeStr})
	}

	var result [][]interface{}

	for key, calls := range userDayCalls {
		// Sort calls by time
  // Custom sort
		sort.Slice(calls, func(i, j int) bool {
			return calls[i].time < calls[j].time
		})

		firstOther := calls[0].other
		lastOther := calls[len(calls)-1].other

		if firstOther == lastOther {
			uid := 0
			fmt.Sscanf(key[0], "%d", &uid)
			result = append(result, []interface{}{uid, firstOther, key[1]})
		}
	}

	// Sort result by user_id, then date
  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		ui := result[i][0].(int)
		uj := result[j][0].(int)
		if ui != uj {
			return ui < uj
		}
		return result[i][2].(string) < result[j][2].(string)
	})

	if result == nil {
		return [][]interface{}{}
	}
	return result
}
```
