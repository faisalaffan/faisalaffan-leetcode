# 1939 — Users That Actively Request Confirmation Messages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func UsersThatActivelyRequestConfirmationMessages(actions [][2]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1939: Users That Actively Request Confirmation Messages
// https://leetcode.com/problems/users-that-actively-request-confirmation-messages/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// user actions: each pair is (userID, action)
	actions := [][2]string{{"1", "confirmed"}, {"2", "timeout"}, {"1", "confirmed"}, {"3", "confirmed"}}
	fmt.Println(UsersThatActivelyRequestConfirmationMessages(actions)) // [1 3]
}

// Time: O(n log n), Space: O(n)
func UsersThatActivelyRequestConfirmationMessages(actions [][2]string) []int {
  // HashMap: O(1) lookup
	count := make(map[int]int)
	for _, a := range actions {
		userID := 0
		for _, c := range a[0] {
			userID = userID*10 + int(c-'0')
		}
		if a[1] == "confirmed" {
			count[userID]++
		}
	}

	var result []int
	for uid, c := range count {
		if c >= 2 {
			result = append(result, uid)
		}
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}
```
