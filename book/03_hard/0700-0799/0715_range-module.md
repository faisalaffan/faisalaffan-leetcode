# 0715 — Range Module

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() RangeModule
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #715: Range Module
// https://leetcode.com/problems/range-module/
// Difficulty: Hard
//
// Maintain sorted disjoint intervals with addRange, queryRange, removeRange.
// All operations O(n) (n = number of intervals).

type RangeModule struct {
	intervals [][2]int
}

func Constructor() RangeModule {
	return RangeModule{intervals: [][2]int{}}
}

func (rm *RangeModule) AddRange(left int, right int) {
	var merged [][2]int
	i := 0
	n := len(rm.intervals)

	for i < n && rm.intervals[i][1] < left {
		merged = append(merged, rm.intervals[i])
		i++
	}
	for i < n && rm.intervals[i][0] <= right {
		if rm.intervals[i][0] < left {
			left = rm.intervals[i][0]
		}
		if rm.intervals[i][1] > right {
			right = rm.intervals[i][1]
		}
		i++
	}
	merged = append(merged, [2]int{left, right})
	for i < n {
		merged = append(merged, rm.intervals[i])
		i++
	}
	rm.intervals = merged
}

func (rm *RangeModule) QueryRange(left int, right int) bool {
	idx := sort.Search(len(rm.intervals), func(i int) bool {
		return rm.intervals[i][1] > left
	})
	return idx < len(rm.intervals) && rm.intervals[idx][0] <= left && rm.intervals[idx][1] >= right
}

func (rm *RangeModule) RemoveRange(left int, right int) {
	var result [][2]int
	for _, interval := range rm.intervals {
		if interval[1] <= left || interval[0] >= right {
			result = append(result, interval)
			continue
		}
		if interval[0] < left {
			result = append(result, [2]int{interval[0], left})
		}
		if interval[1] > right {
			result = append(result, [2]int{right, interval[1]})
		}
	}
	rm.intervals = result
}

func main() {
	rm := Constructor()
	rm.AddRange(10, 20)
	rm.AddRange(25, 30)
	fmt.Println(rm.QueryRange(10, 14)) // true
	fmt.Println(rm.QueryRange(10, 20)) // true
	fmt.Println(rm.QueryRange(14, 21)) // true (10-20 covers 14-20, but 20-21 uncovered) -> false
	fmt.Println(rm.QueryRange(20, 22)) // false
	fmt.Println(rm.QueryRange(15, 20)) // true
	rm.RemoveRange(14, 16)
	fmt.Println(rm.QueryRange(10, 14)) // true
	fmt.Println(rm.QueryRange(13, 15)) // false
	fmt.Println(rm.QueryRange(16, 17)) // true
	rm.AddRange(5, 8)
	fmt.Println(rm.QueryRange(5, 8))  // true
	fmt.Println(rm.QueryRange(0, 5))  // false
	fmt.Println(rm.QueryRange(8, 9))  // false

	fmt.Println("---")

	// Edge: full overlap removal
	rm2 := Constructor()
	rm2.AddRange(1, 10)
	rm2.RemoveRange(3, 7)
	fmt.Println(rm2.QueryRange(1, 3)) // true
	fmt.Println(rm2.QueryRange(3, 7)) // false
	fmt.Println(rm2.QueryRange(7, 10)) // true
	rm2.AddRange(3, 7)
	fmt.Println(rm2.QueryRange(3, 7)) // true
}
```
