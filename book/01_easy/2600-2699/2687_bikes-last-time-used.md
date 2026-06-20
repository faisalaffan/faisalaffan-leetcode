# 2687 — Bikes Last Time Used

## Deskripsi

**Soal:** [2687. Bikes Last Time Used](https://leetcode.com/problems/bikes-last-time-used/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
