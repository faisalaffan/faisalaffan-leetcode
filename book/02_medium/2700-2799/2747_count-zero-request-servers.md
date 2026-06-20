# 2747 — Count Zero Request Servers

## Deskripsi

**Soal:** [2747. Count Zero Request Servers](https://leetcode.com/problems/count-zero-request-servers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func CountZeroRequestServers(n int, logs [][]int, x int, queries []int) []int`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	logList := make([]Log, len(logs))
	for i, l := range logs {
		logList[i] = Log{serverID: l[0], time: l[1]}
	}
	sort.Slice(logList, func(i, j int) bool {
		return logList[i].time < logList[j].time
	})

	// Map queries to their original indices
	type qItem struct {
		time int
		idx  int
	}
  // Membuat slice untuk menyimpan hasil
	qList := make([]qItem, len(queries))
	for i, q := range queries {
		qList[i] = qItem{time: q, idx: i}
	}
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].time < qList[j].time
	})

  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(queries))
  // Membuat map untuk pencarian O(1): key → value
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
