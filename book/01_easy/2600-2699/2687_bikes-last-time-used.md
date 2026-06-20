# 2687 — Bikes Last Time Used

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func BikesLastTimeUsed(rides []struct {
	BikeID int
	Time   int
}) map[int]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
