# 1939 — Users That Actively Request Confirmation Messages

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func UsersThatActivelyRequestConfirmationMessages(actions [][2]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}
```
