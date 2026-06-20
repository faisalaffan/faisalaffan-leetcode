# 1709 — Biggest Window Between Visits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func parseDate(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Visit struct {
	UserId    int
	VisitDate string
}

type UserWindow struct {
	UserId       int
	BiggestWindow int
}

func main() {
	visits1 := []Visit{
		{1, "2020-01-01"},
		{1, "2020-01-10"},
		{1, "2020-01-28"},
		{2, "2020-01-05"},
		{2, "2020-01-20"},
	}
	result1 := biggestWindow(visits1)
	fmt.Println("Test 1:")
	for _, r := range result1 {
		fmt.Printf("  UserId: %d, BiggestWindow: %d\n", r.UserId, r.BiggestWindow)
	}

	visits2 := []Visit{
		{1, "2020-01-24"},
		{1, "2020-01-25"},
		{2, "2020-01-01"},
		{2, "2020-01-02"},
		{2, "2020-01-03"},
		{2, "2020-12-31"},
	}
	result2 := biggestWindow(visits2)
	fmt.Println("\nTest 2:")
	for _, r := range result2 {
		fmt.Printf("  UserId: %d, BiggestWindow: %d\n", r.UserId, r.BiggestWindow)
	}
}

func parseDate(s string) int {
	var year, month, day int
	fmt.Sscanf(s, "%d-%d-%d", &year, &month, &day)
	return year*365 + month*30 + day
}

func biggestWindow(visits []Visit) []UserWindow {
  // Membuat map (HashMap) — pencarian O(1)
	userVisits := make(map[int][]int)
  // Membuat map (HashMap) — pencarian O(1)
	userSet := make(map[int]bool)

	for _, v := range visits {
		userVisits[v.UserId] = append(userVisits[v.UserId], parseDate(v.VisitDate))
		userSet[v.UserId] = true
	}

	result := make([]UserWindow, 0)
	for uid := range userSet {
		dates := userVisits[uid]
  // Urutkan secara ascending — O(n log n)
		sort.Ints(dates)

		maxGap := 0
		for i := 1; i < len(dates); i++ {
			gap := dates[i] - dates[i-1]
			if gap > maxGap {
				maxGap = gap
			}
		}

		result = append(result, UserWindow{UserId: uid, BiggestWindow: maxGap})
	}

	return result
}
```
