# 3611 — Find Overbooked Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindOverbookedEmployees(shifts [][]int, limit int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3611: Find Overbooked Employees
// https://leetcode.com/problems/find-overbooked-employees/
// Difficulty: Medium
// Complexity: O(n log n) time, O(n) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	shifts := [][]int{{1, 3}, {2, 5}, {1, 4}}
	limit := 2
	fmt.Println("Test 1:", FindOverbookedEmployees(shifts, limit))
	// Test case 2
	shifts2 := [][]int{{1, 2}, {3, 4}}
	limit2 := 1
	fmt.Println("Test 2:", FindOverbookedEmployees(shifts2, limit2))
	// Test case 3
	shifts3 := [][]int{{1, 5}, {2, 3}, {4, 6}}
	limit3 := 2
	fmt.Println("Test 3:", FindOverbookedEmployees(shifts3, limit3))
}

func FindOverbookedEmployees(shifts [][]int, limit int) int {
	type event struct {
		time int
		typ  int // 1=start, -1=end
	}
	var events []event
	for _, s := range shifts {
		events = append(events, event{s[0], 1}, event{s[1], -1})
	}
  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		if events[i].time != events[j].time {
			return events[i].time < events[j].time
		}
		return events[i].typ < events[j].typ
	})
	overbooked := 0
	count := 0
	for _, e := range events {
		count += e.typ
		if count > limit {
			overbooked++
		}
	}
	if overbooked > 0 {
		return overbooked
	}
	return 0
}
```
