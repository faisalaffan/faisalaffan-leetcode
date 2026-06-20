# 1847 — Closest Room

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func closestRoom(rooms [][]int, queries [][]int) []int
```

> **💡 Hint:** sort rooms by size descending, sort queries by minSize descending.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1847: Closest Room
// https://leetcode.com/problems/closest-room/
// Difficulty: Hard
//
// There are rooms with (roomId, size). For each query (preferred, minSize),
// find the room with size >= minSize and roomId closest to preferred.
// If tie, choose the smaller roomId.
//
// Approach: sort rooms by size descending, sort queries by minSize descending.
// Process queries in order, adding eligible rooms to a sorted list.
// For each query, use binary search to find the closest roomId.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1:
	// rooms = [[2,2],[1,2],[3,2]]
	// queries = [[3,1],[3,3],[5,2]]
	// Expected: [3,-1,3]
	fmt.Println(closestRoom([][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}}))

	// Example 2:
	// rooms = [[1,4],[2,3],[3,5],[4,1],[5,2]]
	// queries = [[2,3],[2,4],[2,5]]
	// Expected: [2,1,-1]
	fmt.Println(closestRoom([][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}}))

	// Single room, single query
	fmt.Println(closestRoom([][]int{{1, 10}}, [][]int{{5, 5}}))

	// No room meets min size
	fmt.Println(closestRoom([][]int{{1, 5}}, [][]int{{3, 10}}))

	// Multiple rooms, exact match
	fmt.Println(closestRoom([][]int{{10, 20}, {20, 30}, {30, 40}}, [][]int{{25, 25}}))

	// Tie-breaking: smaller roomId wins
	fmt.Println(closestRoom([][]int{{5, 10}, {7, 10}}, [][]int{{6, 5}}))
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	// Sort rooms by size descending
  // Custom sort dengan comparator
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] > rooms[j][1]
	})

	// Sort queries by minSize descending, keeping original index
	type query struct {
		preferred int
		minSize   int
		idx       int
	}
	sortedQueries := make([]query, len(queries))
	for i, q := range queries {
		sortedQueries[i] = query{preferred: q[0], minSize: q[1], idx: i}
	}
  // Custom sort dengan comparator
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i].minSize > sortedQueries[j].minSize
	})

  // Alokasi slice integer
	ans := make([]int, len(queries))
	avail := []int{} // sorted room IDs
	roomIdx := 0

	for _, q := range sortedQueries {
		// Add all rooms with size >= query.minSize
		for roomIdx < len(rooms) && rooms[roomIdx][1] >= q.minSize {
			// Insert room ID in sorted order
			id := rooms[roomIdx][0]
			pos := sort.SearchInts(avail, id)
			avail = append(avail, 0)
			copy(avail[pos+1:], avail[pos:])
			avail[pos] = id
			roomIdx++
		}

		if len(avail) == 0 {
			ans[q.idx] = -1
			continue
		}

		// Binary search for closest room ID
		pos := sort.SearchInts(avail, q.preferred)
		if pos == 0 {
			ans[q.idx] = avail[0]
		} else if pos == len(avail) {
			ans[q.idx] = avail[len(avail)-1]
		} else {
			leftDiff := q.preferred - avail[pos-1]
			rightDiff := avail[pos] - q.preferred
			if leftDiff <= rightDiff {
				ans[q.idx] = avail[pos-1]
			} else {
				ans[q.idx] = avail[pos]
			}
		}
	}
	return ans
}
```
