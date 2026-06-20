# 0759 — Employee Free Time

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func employeeFreeTime(schedule [][]Interval) []Interval
```

> **💡 Hint:** Sweep Line

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #759: Employee Free Time
// https://leetcode.com/problems/employee-free-time/
// Difficulty: Hard [Paid]
//
// Approach: Sweep Line
// 1. Flatten all intervals from all employees.
// 2. Sort by start time.
// 3. Merge overlapping intervals.
// 4. The gaps between merged intervals are free time.

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func main() {
	// Example 1:
	// Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
	// Output: [[3,4]]
	schedule1 := [][]Interval{
		{{1, 2}, {5, 6}},
		{{1, 3}},
		{{4, 10}},
	}
	fmt.Println(employeeFreeTime(schedule1)) // [[3,4]]

	// Example 2:
	// Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
	// Output: [[5,6],[7,9]]
	schedule2 := [][]Interval{
		{{1, 3}, {6, 7}},
		{{2, 4}},
		{{2, 5}, {9, 12}},
	}
	fmt.Println(employeeFreeTime(schedule2)) // [[5,6],[7,9]]
}

func employeeFreeTime(schedule [][]Interval) []Interval {
	// Flatten all intervals
	var all []Interval
	for _, emp := range schedule {
		all = append(all, emp...)
	}

	// Sort by start time, then by end time
  // Custom sort dengan comparator
	sort.Slice(all, func(i, j int) bool {
		if all[i].Start != all[j].Start {
			return all[i].Start < all[j].Start
		}
		return all[i].End < all[j].End
	})

	// Merge overlapping intervals
	var merged []Interval
	for _, iv := range all {
		if len(merged) == 0 || iv.Start > merged[len(merged)-1].End {
			merged = append(merged, iv)
		} else if iv.End > merged[len(merged)-1].End {
			merged[len(merged)-1].End = iv.End
		}
	}

	// Gaps between merged intervals are free time
	var free []Interval
	for i := 1; i < len(merged); i++ {
		if merged[i-1].End < merged[i].Start {
			free = append(free, Interval{merged[i-1].End, merged[i].Start})
		}
	}

	return free
}
```
