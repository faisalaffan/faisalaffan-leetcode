# 1956 — Minimum Time For K Virus Variants To Spread

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minTimeForKVirusVariantsToSpread(points [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1956: Minimum Time For K Virus Variants to Spread
// https://leetcode.com/problems/minimum-time-for-k-virus-variants-to-spread/
// Difficulty: Hard [Paid]
//
// Given n virus variants at positions (xi, yi), each spreads at speed 1
// (Manhattan distance per unit time). Find the minimum integer time T such
// that there exists a point that can be reached by at least k variants.
//
// A variant at (xi, yi) reaches point (x, y) in time T if
// |x - xi| + |y - yi| <= T.
//
// Transform to (u = x+y, v = x-y) space where the reachable region becomes
// an axis-aligned square of side 2T centered at (ui, vi).
// Check function: sweep line over u, maintain difference array for v.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: two points, k=2, meet at midpoint
	// (0,0) and (2,0) -> |x|+|y| <= T, |x-2|+|y| <= T
	// Meet at (1,0) in time 1
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {2, 0}}, 2))

	// Test case 2: three points forming a triangle, k=2
	// (0,0), (1,0), (0,1) -> any two can meet at their midpoint
	// (0,0) and (1,0) meet at (0.5,0.5) ... wait Manhattan midpoint
	// |x|+|y| <= T, |x-1|+|y| <= T for (0.5, 0.5): |0.5|+|0.5|=1, |0.5-1|+|0.5|=1 => T=1
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {1, 0}, {0, 1}}, 2))

	// Test case 3: single point, k=1 -> T=0
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{5, 5}}, 1))

	// Test case 4: colinear, k=2
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {4, 0}}, 2))

	// Test case 5: need exactly k=3 out of 4
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {10, 10}, {0, 10}, {10, 0}}, 3))

	// Test case 6: large gap
	fmt.Println(minTimeForKVirusVariantsToSpread([][]int{{0, 0}, {100, 0}}, 2))
}

// minTimeForKVirusVariantsToSpread returns the minimum integer time T.
func minTimeForKVirusVariantsToSpread(points [][]int, k int) int {
	n := len(points)
	if n < k {
		return -1
	}
	if k <= 1 {
		return 0
	}

	// Binary search on time
	// Upper bound: max coordinate range
	maxCoord := 0
	for _, p := range points {
		for _, c := range p {
			if c > maxCoord {
				maxCoord = c
			}
		}
	}
	high := 2 * maxCoord // safe upper bound
	if high < 0 {
		high = 2000000
	}

	low := 0
	ans := high
	for low <= high {
		mid := (low + high) / 2
		if canMeet(points, k, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func canMeet(points [][]int, k, T int) bool {
	// Transform: u = x+y, v = x-y
	// Each point covers [ui-T, ui+T] in u, [vi-T, vi+T] in v
	type event struct {
		u        int
		vLow     int
		vHigh    int
		isAdd    bool
	}

	var events []event
  // Membuat map (HashMap) — pencarian O(1)
	vSet := make(map[int]bool)

	for _, p := range points {
		x, y := p[0], p[1]
		u := x + y
		v := x - y
		events = append(events, event{u - T, v - T, v + T, true})
		events = append(events, event{u + T + 1, v - T, v + T, false})
		vSet[v-T] = true
		vSet[v+T] = true
	}

	// Coordinate compress v values
	var vVals []int
	for v := range vSet {
		vVals = append(vVals, v)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(vVals)
  // Membuat map (HashMap) — pencarian O(1)
	vComp := make(map[int]int)
	for i, v := range vVals {
		vComp[v] = i
	}

  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		return events[i].u < events[j].u
	})

	// Difference array over compressed v
  // Alokasi slice integer
	diff := make([]int, len(vVals)+1)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(events); {
		curU := events[i].u
		// Apply all events at this u
		for i < len(events) && events[i].u == curU {
			e := events[i]
			l := vComp[e.vLow]
			r := vComp[e.vHigh]
			if e.isAdd {
				diff[l]++
				diff[r+1]--
			} else {
				diff[l]--
				diff[r+1]++
			}
			i++
		}
		// Check current sweep position
		cur := 0
		for j := 0; j < len(vVals)-1; j++ {
			cur += diff[j]
			if cur >= k {
				return true
			}
		}
	}

	return false
}
```
