# 1149 — Article Views Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func articleViewsII(views [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1149: Article Views II
// https://leetcode.com/problems/article-views-ii/
// Difficulty: Medium [Paid]

// Given a table with viewer_id, article_id, and view_date, find all viewers
// who viewed at least two distinct articles on the same day.
// Return result sorted by viewer_id.

// Time: O(n log n)
// Space: O(n)

func articleViewsII(views [][]int) []int {
	// views[i] = [viewer_id, article_id, view_date]
	type view struct {
		viewerID int
		date     int
		article  int
	}

	type key struct {
		viewerID int
		date     int
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[key]map[int]bool)

	for _, v := range views {
		k := key{v[0], v[2]}
		if seen[k] == nil {
			seen[k] = make(map[int]bool)
		}
		seen[k][v[1]] = true
	}

  // Membuat map (HashMap) — pencarian O(1)
	resultSet := make(map[int]bool)
	for k, articles := range seen {
		if len(articles) >= 2 {
			resultSet[k.viewerID] = true
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, len(resultSet))
	for id := range resultSet {
		result = append(result, id)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	views := [][]int{
		{1, 1, 20200101},
		{1, 2, 20200101},
		{2, 3, 20200101},
		{1, 3, 20200102},
	}
	fmt.Printf("%v (expected: [1])\n", articleViewsII(views))

	views2 := [][]int{
		{1, 1, 20200101},
		{1, 2, 20200101},
		{2, 1, 20200101},
		{2, 3, 20200101},
	}
	fmt.Printf("%v (expected: [1 2])\n", articleViewsII(views2))

	views3 := [][]int{
		{1, 1, 20200101},
		{1, 1, 20200101},
	}
	fmt.Printf("%v (expected: [])\n", articleViewsII(views3))
}
```
