# 1840 — Maximum Building Height

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxBuilding(n int, restrictions [][]int) int
```

> **💡 Hint:** Two-Pass Constraint Propagation.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1840: Maximum Building Height
// https://leetcode.com/problems/maximum-building-height/
// Difficulty: Hard
//
// Approach: Two-Pass Constraint Propagation.
//   We have n buildings numbered 1..n. Each restriction [id, h] says
//   the height of building `id` cannot exceed `h`. Adjacent buildings
//   differ in height by at most 1.
//
//   1. Add implicit restriction: building 1 max height = 0.
//   2. Sort restrictions by id.
//   3. Forward pass: height[i] = min(height[i], height[i-1] + (id_i - id_{i-1}))
//   4. Backward pass: height[i] = min(height[i], height[i+1] + (id_{i+1} - id_i))
//   5. Between each pair of consecutive restrictions, the max height is
//      achieved at a peak point. The max building height between id[a] and id[b]
//      with heights h[a] and h[b] is:
//        dist = id[b] - id[a]
//        h[a] + h[b] + dist
//        peak = ---------------
//                2            (using integer arithmetic)
//      The max height in this interval = max(h[a], h[b]) + floor((dist - |h[a]-h[b]|) / 2)

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println("Example 1:", maxBuilding(5, [][]int{{2, 1}, {4, 1}}))
	// Expected: 2

	// Example 2
	fmt.Println("Example 2:", maxBuilding(6, [][]int{}))
	// Expected: 5 (heights 0,1,2,3,4,5; max = 5)

	// Example 3
	fmt.Println("Example 3:", maxBuilding(10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 0}}))
	// Expected: 5

	// Edge case: single restriction
	fmt.Println("Edge:", maxBuilding(3, [][]int{{3, 0}}))
	// Expected: 1 (0,1,0 -> max=1)
}

func maxBuilding(n int, restrictions [][]int) int {
	// Add implicit restriction: building 1 must have height 0
	type res struct {
		id, h int
	}
	list := make([]res, 0, len(restrictions)+1)
	list = append(list, res{1, 0})
	for _, r := range restrictions {
		if r[0] == 1 {
			if r[1] < 0 {
				r[1] = 0
			}
			list[0].h = min(r[1], list[0].h)
			continue
		}
		list = append(list, res{r[0], r[1]})
	}

	// Sort by id
  // Custom sort dengan comparator
	sort.Slice(list, func(i, j int) bool {
		return list[i].id < list[j].id
	})

	// Forward pass: propagate left-to-right constraints
	for i := 1; i < len(list); i++ {
		dist := list[i].id - list[i-1].id
		list[i].h = min(list[i].h, list[i-1].h+dist)
	}

	// Backward pass: propagate right-to-left constraints
	for i := len(list) - 2; i >= 0; i-- {
		dist := list[i+1].id - list[i].id
		list[i].h = min(list[i].h, list[i+1].h+dist)
	}

	// Compute max building height between pairs of restrictions
	ans := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(list); i++ {
		if list[i].h > ans {
			ans = list[i].h
		}
		if i+1 < len(list) {
			// Between list[i] and list[i+1], find the peak height
			dist := list[i+1].id - list[i].id
			hDiff := abs(list[i].h - list[i+1].h)
			// The height rises from the lower to the higher, then peaks
			remaining := dist - hDiff
			peak := max(list[i].h, list[i+1].h) + remaining/2
			if peak > ans {
				ans = peak
			}
		}
	}

	// Also consider building n (no explicit restriction after the last one)
	if len(list) > 0 {
		last := list[len(list)-1]
		remaining := n - last.id
		potential := last.h + remaining
		if potential > ans {
			ans = potential
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Stub kept for compatibility with the repo scaffold.
func MaximumBuildingHeight() any {
	return maxBuilding(5, [][]int{{2, 1}, {4, 1}})
}
```
