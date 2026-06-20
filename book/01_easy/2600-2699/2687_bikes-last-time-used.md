# 2687 — Bikes Last Time Used

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func BikesLastTimeUsed(rides []struct { BikeID int Time int }) map[int]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2687: Bikes Last Time Used
// https://leetcode.com/problems/bikes-last-time-used/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL/JS problem adapted to Go. Find last time each bike was used.

import "fmt"

func main() {
	rides := []struct {
		BikeID int
		Time   int
	}{
		{1, 100},
		{2, 150},
		{1, 200},
		{3, 50},
	}
	fmt.Println(BikesLastTimeUsed(rides))
}

func BikesLastTimeUsed(rides []struct {
	BikeID int
	Time   int
}) map[int]int {
	lastUsed := map[int]int{}
	for _, r := range rides {
		if r.Time > lastUsed[r.BikeID] {
			lastUsed[r.BikeID] = r.Time
		}
	}
	return lastUsed
}
```
