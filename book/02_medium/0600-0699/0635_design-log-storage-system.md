# 0635 — Design Log Storage System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() LogSystem
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie, Prefix Sum

**Kompleksitas Waktu:** O(n) for put, O(n) for retrieve  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
