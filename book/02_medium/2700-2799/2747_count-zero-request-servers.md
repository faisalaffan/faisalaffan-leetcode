# 2747 — Count Zero Request Servers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CountZeroRequestServers(n int, logs [][]int, x int, queries []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2747: Count Zero Request Servers
// https://leetcode.com/problems/count-zero-request-servers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Log struct {
	serverID int
	time     int
}

func CountZeroRequestServers(n int, logs [][]int, x int, queries []int) []int {
	// Process logs
	logList := make([]Log, len(logs))
	for i, l := range logs {
		logList[i] = Log{serverID: l[0], time: l[1]}
	}
  // Custom sort
	sort.Slice(logList, func(i, j int) bool {
		return logList[i].time < logList[j].time
	})

	// Map queries to their original indices
	type qItem struct {
		time int
		idx  int
	}
	qList := make([]qItem, len(queries))
	for i, q := range queries {
		qList[i] = qItem{time: q, idx: i}
	}
  // Custom sort
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].time < qList[j].time
	})

  // Alokasi slice
	result := make([]int, len(queries))
  // HashMap: O(1) lookup
	active := make(map[int]int)
	left := 0

	for _, q := range qList {
		end := q.time
		start := q.time - x

		// Add logs within [start, end]
		for left < len(logList) && logList[left].time <= end {
			active[logList[left].serverID]++
			left++
		}

		// Remove logs before start
		right := 0
		for right < len(logList) && logList[right].time < start {
			if _, ok := active[logList[right].serverID]; ok {
				active[logList[right].serverID]--
				if active[logList[right].serverID] == 0 {
					delete(active, logList[right].serverID)
				}
			}
			right++
		}

		result[q.idx] = n - len(active)
	}

	return result
}

func main() {
	fmt.Println(CountZeroRequestServers(3, [][]int{{1, 3}, {2, 6}, {1, 5}}, 5, []int{10, 11}))
	fmt.Println(CountZeroRequestServers(3, [][]int{{1, 1}, {2, 2}, {3, 3}}, 2, []int{1, 5}))
}
```
