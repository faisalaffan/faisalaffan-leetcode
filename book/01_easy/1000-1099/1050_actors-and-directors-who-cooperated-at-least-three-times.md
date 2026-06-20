# 1050 — Actors And Directors Who Cooperated At Least Three Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func actorsAndDirectorsWhoCooperatedAtLeastThreeTimes(pairs []actorDirector) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
