# 1709 — Biggest Window Between Visits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func biggestWindow(visits []Visit) []UserWindow`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
	userVisits := make(map[int][]int)
  // HashMap: O(1) lookup
	userSet := make(map[int]bool)

	for _, v := range visits {
		userVisits[v.UserId] = append(userVisits[v.UserId], parseDate(v.VisitDate))
		userSet[v.UserId] = true
	}

	result := make([]UserWindow, 0)
	for uid := range userSet {
		dates := userVisits[uid]
  // Sort O(n log n)
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
