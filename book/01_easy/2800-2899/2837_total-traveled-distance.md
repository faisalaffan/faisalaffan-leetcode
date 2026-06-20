# 2837 — Total Traveled Distance

## Deskripsi

**Soal:** [2837. Total Traveled Distance](https://leetcode.com/problems/total-traveled-distance/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
