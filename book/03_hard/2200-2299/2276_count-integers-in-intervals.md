# 2276 — Count Integers In Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor() CountIntervals`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// CountIntervals maintains a set of disjoint intervals inserted via Add,
// and can report the total number of distinct integers covered.
// Intervals are stored in a sorted slice and merged on every add.
type CountIntervals struct {
	intervals [][2]int // sorted by left endpoint, non-overlapping
	total     int
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

// Add inserts the interval [left, right] (inclusive) and merges overlaps.
func (ci *CountIntervals) Add(left, right int) {
	// Find the insertion / merge range via binary search.
	n := len(ci.intervals)
	// first index with end >= left
	lo := sort.Search(n, func(i int) bool { return ci.intervals[i][1] >= left })
	// last index with start <= right
	hi := sort.Search(n, func(i int) bool { return ci.intervals[i][0] > right })

	if lo == hi {
		// No overlap – insert as a fresh interval at position lo.
		ci.intervals = append(ci.intervals, [2]int{0, 0})
		copy(ci.intervals[lo+1:], ci.intervals[lo:])
		ci.intervals[lo] = [2]int{left, right}
		ci.total += right - left + 1
		return
	}

	// Subtract old contributions.
	for _, iv := range ci.intervals[lo:hi] {
		ci.total -= iv[1] - iv[0] + 1
	}

	// Merge. The merged interval spans from min(left, intervals[lo][0]) to
	// max(right, intervals[hi-1][1]).
	if left > ci.intervals[lo][0] {
		left = ci.intervals[lo][0]
	}
	if right < ci.intervals[hi-1][1] {
		right = ci.intervals[hi-1][1]
	}

	merged := [2]int{left, right}
	// Replace intervals[lo:hi] with the single merged interval.
	ci.intervals = append(ci.intervals[:lo], ci.intervals[hi:]...)
	ci.intervals = append(ci.intervals, [2]int{0, 0})
	copy(ci.intervals[lo+1:], ci.intervals[lo:])
	ci.intervals[lo] = merged

	ci.total += right - left + 1
}

// Count returns the total number of distinct integers covered by all intervals.
func (ci *CountIntervals) Count() int {
	return ci.total
}

// ---------------------------------------------------------------------------
//  LeetCode-style wrapper (required by stub convention)
func CountIntegersInIntervals() interface{} {
	ci := Constructor()
	ci.Add(2, 3)
	ci.Add(7, 10)
	_ = ci.Count()
	ci.Add(5, 8)
	return ci.Count()
}

func main() {
	fmt.Println(CountIntegersInIntervals())

	// ---- test cases ----
	testCases := []struct {
		ops    []string
		args   [][2]int // (-1,-1) for Count
		want   int       // for Count
		expect []int     // for Count when there are multiple calls
	}{
		{
			ops:    []string{"add", "add", "count", "add", "count"},
			args:   [][2]int{{2, 3}, {7, 10}, {-1, -1}, {5, 8}, {-1, -1}},
			expect: []int{0, 0, 6, 0, 8},
		},
	}

	for idx, tc := range testCases {
		ci := Constructor()
		for j, op := range tc.ops {
			switch op {
			case "add":
				ci.Add(tc.args[j][0], tc.args[j][1])
			case "count":
				got := ci.Count()
				if got != tc.expect[j] {
					fmt.Printf("FAIL tc %d step %d: got %d, want %d\n", idx, j, got, tc.expect[j])
				}
			}
		}
	}
	fmt.Println("Done testing 2276.")
}
```
