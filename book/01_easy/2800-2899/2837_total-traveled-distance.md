# 2837 — Total Traveled Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func TotalTraveledDistance(rides []struct { UserID int Distance int }) map[int]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2837: Total Traveled Distance
// https://leetcode.com/problems/total-traveled-distance/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL problem, adapted to Go. Computes total travel distance per user.

import "fmt"

func main() {
	rides := []struct {
		UserID   int
		Distance int
	}{
		{1, 10},
		{1, 15},
		{2, 20},
		{3, 0},
	}
	fmt.Println(TotalTraveledDistance(rides))
}

func TotalTraveledDistance(rides []struct {
	UserID   int
	Distance int
}) map[int]int {
	total := map[int]int{}
	for _, r := range rides {
		total[r.UserID] += r.Distance
	}
	return total
}
```
