# 1050 — Actors And Directors Who Cooperated At Least Three Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs []actorDirector) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1050: Actors and Directors Who Cooperated At Least Three Times
// https://leetcode.com/problems/actors-and-directors-who-cooperated-at-least-three-times/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)
// Note: This is a SQL problem. Go implementation simulates the query logic.

import "fmt"

type actorDirector struct {
	actorID   int
	directorID int
	timestamp int
}

func main() {
	pairs := []actorDirector{
		{1, 1, 0}, {1, 1, 1}, {1, 1, 2},
		{1, 2, 3}, {2, 1, 4}, {2, 1, 5},
	}
	fmt.Println(actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs)) // [[1 1]]

	pairs2 := []actorDirector{{1, 1, 0}, {1, 1, 1}}
	fmt.Println(actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs2)) // []
}

// LeetCode submission: actorsAndDirectorsWhoCooperatedAtLeastThreeTimes (SQL equivalent)
func actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs []actorDirector) [][]int {
  // HashMap: O(1) lookup
	count := make(map[[2]int]int)
	for _, p := range pairs {
		key := [2]int{p.actorID, p.directorID}
		count[key]++
	}
	var ans [][]int
	for k, v := range count {
		if v >= 3 {
			ans = append(ans, []int{k[0], k[1]})
		}
	}
	return ans
}
```
