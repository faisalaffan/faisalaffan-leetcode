# 0352 — Data Stream As Disjoint Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() SummaryRanges
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #352: Data Stream as Disjoint Intervals
// https://leetcode.com/problems/data-stream-as-disjoint-intervals/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// SummaryRanges maintains a set of disjoint intervals from a stream of numbers.
type SummaryRanges struct {
	intervals [][]int // sorted by start, each is [start, end] inclusive
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (sr *SummaryRanges) AddNum(val int) {
	// Find position using binary search
	idx := sort.Search(len(sr.intervals), func(i int) bool {
		return sr.intervals[i][0] > val
	})

	// Check for merge with left neighbor (intervals[idx-1])
	if idx > 0 && sr.intervals[idx-1][1] >= val-1 {
		// Merge with left
		if val > sr.intervals[idx-1][1] {
			sr.intervals[idx-1][1] = val
		}
		// Check if we also need to merge with right
		if idx < len(sr.intervals) && sr.intervals[idx][0] <= sr.intervals[idx-1][1]+1 {
			if sr.intervals[idx][1] > sr.intervals[idx-1][1] {
				sr.intervals[idx-1][1] = sr.intervals[idx][1]
			}
			sr.intervals = append(sr.intervals[:idx], sr.intervals[idx+1:]...)
		}
	} else if idx < len(sr.intervals) && sr.intervals[idx][0] <= val+1 {
		// Merge with right only
		if val < sr.intervals[idx][0] {
			sr.intervals[idx][0] = val
		}
	} else {
		// New isolated interval - insert at position idx
		newInterval := []int{val, val}
		// Insert at idx
		sr.intervals = append(sr.intervals, nil)
		copy(sr.intervals[idx+1:], sr.intervals[idx:])
		sr.intervals[idx] = newInterval
	}
}

func (sr *SummaryRanges) GetIntervals() [][]int {
	return sr.intervals
}

func main() {
	// Example 1
	sr := Constructor()
	sr.AddNum(1)
	sr.AddNum(3)
	sr.AddNum(7)
	sr.AddNum(2)
	sr.AddNum(6)
	fmt.Println(sr.GetIntervals())
	// [[1 3] [6 7]]

	// Example 2
	sr2 := Constructor()
	sr2.AddNum(1)
	sr2.AddNum(3)
	sr2.AddNum(2)
	fmt.Println(sr2.GetIntervals())
	// [[1 3]]

	// Example 3: empty
	sr3 := Constructor()
	fmt.Println(sr3.GetIntervals())
	// []

	// Example 4: single
	sr4 := Constructor()
	sr4.AddNum(5)
	fmt.Println(sr4.GetIntervals())
	// [[5 5]]
}
```
