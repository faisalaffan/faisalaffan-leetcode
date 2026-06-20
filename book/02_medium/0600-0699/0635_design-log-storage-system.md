# 0635 — Design Log Storage System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor() LogSystem`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Trie, Prefix Sum

**Waktu:** O(n) for put, O(n) for retrieve  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #635: Design Log Storage System
// https://leetcode.com/problems/design-log-storage-system/
// Difficulty: Medium [Paid]
// Time: O(n) for put, O(n) for retrieve
// Space: O(n)

import (
	"fmt"
	"strings"
)

type LogSystem struct {
	logs []LogEntry
}

type LogEntry struct {
	id   int
	time string
}

func Constructor() LogSystem {
	return LogSystem{logs: []LogEntry{}}
}

func (ls *LogSystem) Put(id int, timestamp string) {
	ls.logs = append(ls.logs, LogEntry{id: id, time: timestamp})
}

func (ls *LogSystem) Retrieve(start string, end string, granularity string) []int {
	startPrefix := truncate(start, granularity)
	endPrefix := truncate(end, granularity)

	result := []int{}
	for _, entry := range ls.logs {
		entryPrefix := truncate(entry.time, granularity)
		if entryPrefix >= startPrefix && entryPrefix <= endPrefix {
			result = append(result, entry.id)
		}
	}
	return result
}

func truncate(timestamp string, granularity string) string {
	parts := strings.Split(timestamp, ":")
	indices := map[string]int{
		"Year": 1, "Month": 2, "Day": 3,
		"Hour": 4, "Minute": 5, "Second": 6,
	}
	idx := indices[granularity]
	result := strings.Join(parts[:idx], ":")
	// Pad remaining with minimum values
	for i := idx; i < 6; i++ {
		result += ":00"
	}
	return result
}

func main() {
	ls := Constructor()
	ls.Put(1, "2017:01:01:23:59:59")
	ls.Put(2, "2017:01:01:22:59:59")
	ls.Put(3, "2016:01:01:00:00:00")
	fmt.Println(ls.Retrieve("2016:12:31:23:59:59", "2017:01:02:23:59:59", "Year"))
}
```
