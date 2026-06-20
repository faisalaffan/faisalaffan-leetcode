# 0732 — My Calendar Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor() MyCalendarThree`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #732: My Calendar III
// https://leetcode.com/problems/my-calendar-iii/
// Difficulty: Hard
//
// Algorithm: Sweep Line / Difference Array
// Use a map to track +1 at start time, -1 at end time.
// Maintain running sum to find the maximum number of concurrent bookings (k-booking).

import (
	"fmt"
	"sort"
)

// MyCalendarThree tracks the maximum k-booking (number of concurrent events)
type MyCalendarThree struct {
	events map[int]int // time -> delta (+1 for start, -1 for end)
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{events: make(map[int]int)}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.events[startTime]++
	this.events[endTime]--

	// Sweep line: collect all times
  // Alokasi slice
	times := make([]int, 0, len(this.events))
	for t := range this.events {
		times = append(times, t)
	}
  // Sort O(n log n)
	sort.Ints(times)

	active := 0
	maxActive := 0
	for _, t := range times {
		active += this.events[t]
		if active > maxActive {
			maxActive = active
		}
	}

	return maxActive
}

func main() {
	// Example from problem
	cal := Constructor()
	fmt.Println("MyCalendarIII bookings:")
	bookings := [][2]int{
		{10, 20}, // 1
		{50, 60}, // 1
		{10, 40}, // 2
		{5, 15},  // 3
		{5, 10},  // 3
		{25, 55}, // 3
	}
	for i, b := range bookings {
		result := cal.Book(b[0], b[1])
		fmt.Printf("  Book(%d, %d) = %d (expected: ", b[0], b[1], result)
		expected := []int{1, 1, 2, 3, 3, 3}
		fmt.Printf("%d)\n", expected[i])
	}
	fmt.Println()

	// Test case 2: all overlapping
	cal2 := Constructor()
	fmt.Println("All overlapping:")
	r1 := cal2.Book(0, 10)
	r2 := cal2.Book(0, 10)
	r3 := cal2.Book(0, 10)
	fmt.Printf("  %d %d %d (expected: 1 2 3)\n", r1, r2, r3)
	fmt.Println()

	// Test case 3: non-overlapping
	cal3 := Constructor()
	fmt.Println("Non-overlapping:")
	r1 = cal3.Book(0, 5)
	r2 = cal3.Book(5, 10)
	r3 = cal3.Book(10, 15)
	fmt.Printf("  %d %d %d (expected: 1 1 1)\n", r1, r2, r3)
	fmt.Println()

	// Test case 4: gradual overlap
	cal4 := Constructor()
	fmt.Println("Gradual overlap:")
	results := []int{
		cal4.Book(1, 5),   // 1
		cal4.Book(2, 6),   // 2
		cal4.Book(3, 7),   // 3
		cal4.Book(4, 8),   // 4
	}
	for i, r := range results {
		fmt.Printf("  Book %d = %d\n", i+1, r)
	}
}
```
