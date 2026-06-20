# 2747 — Count Zero Request Servers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountZeroRequestServers(n int, logs [][]int, x int, queries []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Custom sort dengan comparator
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
  // Custom sort dengan comparator
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].time < qList[j].time
	})

  // Alokasi slice integer
	result := make([]int, len(queries))
  // Membuat map (HashMap) — pencarian O(1)
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
