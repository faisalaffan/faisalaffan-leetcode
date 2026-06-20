# 0497 — Random Point In Non Overlapping Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(rects [][]int) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Prefix Sum

**Kompleksitas Waktu:** O(n) for init, O(log n) per pick  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #497: Random Point in Non-overlapping Rectangles
// https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/
// Difficulty: Medium
// Time: O(n) for init, O(log n) per pick
// Space: O(n)

import (
	"fmt"
	"math/rand"
)

func main() {
	sol := Constructor([][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}})
	for i := 0; i < 3; i++ {
		p := sol.Pick()
		fmt.Println(p)
	}
}

type Solution struct {
	rects      [][]int
	prefixSum  []int
	totalPts   int
}

func Constructor(rects [][]int) Solution {
  // Alokasi slice integer
	prefixSum := make([]int, len(rects))
	total := 0
	for i, r := range rects {
		pts := (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		total += pts
		prefixSum[i] = total
	}
	return Solution{rects: rects, prefixSum: prefixSum, totalPts: total}
}

func (s *Solution) Pick() []int {
	// Pick a random point index
	r := rand.Intn(s.totalPts) + 1
	// Binary search to find which rectangle
	idx := search(s.prefixSum, r)

	rect := s.rects[idx]
	prev := 0
	if idx > 0 {
		prev = s.prefixSum[idx-1]
	}
	offset := r - prev - 1
	width := rect[2] - rect[0] + 1
	x := rect[0] + offset%width
	y := rect[1] + offset/width
	return []int{x, y}
}

func search(prefixSum []int, target int) int {
	lo, hi := 0, len(prefixSum)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if prefixSum[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```
