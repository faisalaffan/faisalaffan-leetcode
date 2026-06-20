# 2613 — Beautiful Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulPairs(nums1, nums2 [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Merge Sort

**Kompleksitas Waktu:** O(n log^2 n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Merge Sort** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2613: Beautiful Pairs
// https://leetcode.com/problems/beautiful-pairs/
// Difficulty: Hard [Paid]
//
// Given two arrays of points nums1 and nums2, find a pair (i,j) such that
// |x1_i - x2_j| + |y1_i - y2_j| is minimized. Return [i, j].
// Uses divide-and-conquer closest-pair algorithm on merged point set.
// Time: O(n log^2 n), Space: O(n)

import (
	"fmt"
	"math"
	"sort"
)

type pt struct {
	x, y, idx int
	fromA     bool // true = nums1, false = nums2
}

func main() {
	// Example 1: nums1=[[1,2],[3,4]], nums2=[[1,0],[3,2]]
	fmt.Println(beautifulPairs([][]int{{1, 2}, {3, 4}}, [][]int{{1, 0}, {3, 2}}))
	// Example 2: single point
	fmt.Println(beautifulPairs([][]int{{0, 0}}, [][]int{{1, 1}}))
	// Example 3: simple
	fmt.Println(beautifulPairs([][]int{{0, 0}, {1, 1}}, [][]int{{2, 2}, {3, 3}}))
}

func beautifulPairs(nums1, nums2 [][]int) []int {
	m, n := len(nums1), len(nums2)
	pts := make([]pt, m+n)
	for i := 0; i < m; i++ {
		pts[i] = pt{nums1[i][0], nums1[i][1], i, true}
	}
	for j := 0; j < n; j++ {
		pts[m+j] = pt{nums2[j][0], nums2[j][1], j, false}
	}
  // Custom sort dengan comparator
	sort.Slice(pts, func(i, j int) bool {
		if pts[i].x != pts[j].x {
			return pts[i].x < pts[j].x
		}
		return pts[i].y < pts[j].y
	})

	bestDist := math.MaxInt32
	bestI, bestJ := 0, 0

	var dc func(l, r int)
	dc = func(l, r int) {
		if r-l <= 1 {
			return
		}
		if r-l <= 8 {
			for i := l; i < r; i++ {
				for j := i + 1; j < r; j++ {
					if pts[i].fromA != pts[j].fromA {
						d := abs(pts[i].x-pts[j].x) + abs(pts[i].y-pts[j].y)
						if d < bestDist {
							bestDist = d
							if pts[i].fromA {
								bestI, bestJ = pts[i].idx, pts[j].idx
							} else {
								bestI, bestJ = pts[j].idx, pts[i].idx
							}
						}
					}
				}
			}
			return
		}

		mid := (l + r) / 2
		midX := pts[mid].x
		dc(l, mid)
		dc(mid, r)

		strip := []pt{}
		for i := l; i < r; i++ {
			if abs(pts[i].x-midX) < bestDist {
				strip = append(strip, pts[i])
			}
		}
  // Custom sort dengan comparator
		sort.Slice(strip, func(i, j int) bool {
			return strip[i].y < strip[j].y
		})

  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(strip); i++ {
			for j := i + 1; j < len(strip) && strip[j].y-strip[i].y < bestDist; j++ {
				if strip[i].fromA != strip[j].fromA {
					d := abs(strip[i].x-strip[j].x) + abs(strip[i].y-strip[j].y)
					if d < bestDist {
						bestDist = d
						if strip[i].fromA {
							bestI, bestJ = strip[i].idx, strip[j].idx
						} else {
							bestI, bestJ = strip[j].idx, strip[i].idx
						}
					}
				}
			}
		}
	}

	dc(0, len(pts))
	return []int{bestI, bestJ}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
