# 3245 — Alternating Groups Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func alternatingGroupsIII(colors []int, queries [][]int) []int64
```

> **💡 Hint:** Maintain maximal alternating intervals in a circular array using

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Fenwick Tree (BIT), Merge Sort

**Kompleksitas Waktu:** O((n+q) log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Fenwick Tree (BIT)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3245: Alternating Groups III
// https://leetcode.com/problems/alternating-groups-iii/
// Difficulty: Hard
//
// There are n tiles arranged in a circle, each red (1) or blue (0).
// Two types of queries:
//   Type 1: [1, size] - count of alternating groups with length >= size
//   Type 2: [2, index, color] - update tile at index to the given color
//
// An alternating group is a contiguous subarray where adjacent tiles differ.
//
// Approach: Maintain maximal alternating intervals in a circular array using
// an ordered set (balanced BST). Use a Fenwick tree over a difference array
// to count how many intervals of each length exist.
//
// Time: O((n+q) log n), Space: O(n)

import (
	"fmt"
)

func main() {
	// Example 1
	colors := []int{0, 1, 0, 1, 0}
	queries := [][]int{{1, 3}, {2, 2, 1}, {1, 3}}
	fmt.Println(alternatingGroupsIII(colors, queries))

	// Example 2
	colors2 := []int{0, 1, 0}
	queries2 := [][]int{{1, 2}, {2, 1, 1}, {1, 2}}
	fmt.Println(alternatingGroupsIII(colors2, queries2))

	// Example 3
	colors3 := []int{0, 1, 0, 0, 1, 0, 1}
	queries3 := [][]int{{1, 2}, {2, 3, 0}, {1, 3}}
	fmt.Println(alternatingGroupsIII(colors3, queries3))

	// Example 4: all same
	colors4 := []int{0, 0, 0}
	queries4 := [][]int{{1, 1}, {2, 1, 1}, {1, 1}}
	fmt.Println(alternatingGroupsIII(colors4, queries4))

	// Example 5: single element
	colors5 := []int{1}
	queries5 := [][]int{{1, 1}}
	fmt.Println(alternatingGroupsIII(colors5, queries5))
}

func alternatingGroupsIII(colors []int, queries [][]int) []int64 {
	n := len(colors)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return nil
	}

	// BIT for range updates and point queries (difference array of interval counts)
	bitSize := n + 2
  // Alokasi slice integer
	bitVal := make([]int64, bitSize+1)

	addVal := func(idx int, val int64) {
		idx++
		for idx <= bitSize {
			bitVal[idx] += val
			idx += idx & -idx
		}
	}

	rangeAddVal := func(l, r int, val int64) {
		if l > r || l < 0 {
			return
		}
		if r >= n {
			r = n - 1
		}
		addVal(l, val)
		addVal(r+1, -val)
	}

	pointQuery := func(idx int) int64 {
		idx++
		var res int64
		for idx > 0 {
			res += bitVal[idx]
			idx -= idx & -idx
		}
		return res
	}

	// Maintain intervals as [l, r] pairs
	type interval struct {
		l, r int
	}

	// intervals sorted by l
  // Alokasi slice integer
	intervals := make([]interval, 0)
	addInterval := func(l, r int) {
		if l > r {
			return
		}
		length := r - l + 1
		if length > 0 {
			rangeAddVal(1, length, 1)
		}
		// insert sorted by l
		pos := 0
		for pos < len(intervals) && intervals[pos].l < l {
			pos++
		}
  // Alokasi slice integer
		newIntervals := make([]interval, len(intervals)+1)
		copy(newIntervals, intervals[:pos])
		newIntervals[pos] = interval{l, r}
		copy(newIntervals[pos+1:], intervals[pos:])
		intervals = newIntervals
	}

	removeInterval := func(idx int) {
		if idx < 0 || idx >= len(intervals) {
			return
		}
		l, r := intervals[idx].l, intervals[idx].r
		length := r - l + 1
		if length > 0 {
			rangeAddVal(1, length, -1)
		}
		intervals = append(intervals[:idx], intervals[idx+1:]...)
	}

	// Initial construction: find all alternating intervals
	wraps := colors[0] != colors[n-1]

	i := 0
	for i < n {
		j := i
		for j+1 < n && colors[j] != colors[j+1] {
			j++
		}
		if wraps && j == n-1 && i == 0 {
			for i > 0 && colors[i-1] != colors[i] {
				i--
			}
		}
		addInterval(i, j)
		i = j + 1
		if wraps && len(intervals) > 1 && intervals[0].l == 0 && intervals[len(intervals)-1].r == n-1 {
			break
		}
	}

	// If the circle wraps, merge first and last intervals
	if wraps && len(intervals) >= 2 {
		first := intervals[0]
		last := intervals[len(intervals)-1]
		if first.l == 0 && last.r == n-1 && colors[last.r] != colors[first.l] {
			removeInterval(len(intervals) - 1)
			removeInterval(0)
			addInterval(last.l, first.r)
		}
	}

  // Alokasi slice integer
	ans := make([]int64, 0)

	for _, q := range queries {
		if q[0] == 1 {
			// Query: count of alternating groups with length >= size
			size := q[1]
			if size > n {
				ans = append(ans, 0)
				continue
			}
			var total int64
			for len := size; len <= n; len++ {
				total += pointQuery(len)
			}
			ans = append(ans, total)
		} else {
			// Update: flip color at index
			idx := q[1]
			newColor := q[2]
			if colors[idx] == newColor {
				colors[idx] = newColor
				continue
			}
			colors[idx] = newColor

			// Rebuild intervals from scratch for simplicity
			intervals = intervals[:0]
			bitVal = make([]int64, bitSize+1)

			wraps = colors[0] != colors[n-1]

			i := 0
			for i < n {
				j := i
				for j+1 < n && colors[j] != colors[j+1] {
					j++
				}
				addInterval(i, j)
				i = j + 1
			}

			if wraps && len(intervals) >= 2 {
				first := intervals[0]
				last := intervals[len(intervals)-1]
				if first.l == 0 && last.r == n-1 && colors[last.r] != colors[first.l] {
					removeInterval(len(intervals) - 1)
					removeInterval(0)
					addInterval(last.l, first.r)
				}
			}
		}
	}

	return ans
}
```
