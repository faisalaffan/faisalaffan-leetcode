# 3382 — Maximum Area Rectangle With Point Constraints Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewBIT(n int) *BIT
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3382: Maximum Area Rectangle With Point Constraints II
// https://leetcode.com/problems/maximum-area-rectangle-with-point-constraints-ii/
// Difficulty: Hard
//
// Sweep line left-to-right. For each pair of points sharing the same x (vertical
// neighbors), check if same y-pair was seen earlier. Use Fenwick Tree to verify
// that no other points lie inside the candidate rectangle.

import (
	"fmt"
	"sort"
)

func main() {
	// Example: points=[[1,1],[1,3],[3,1],[3,3]] -> 4
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}}))

	// No rectangle
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {2, 2}, {3, 3}}))

	// 4 points forming rect with extra interior points -> no valid rect
	fmt.Println(maxRectangleArea([][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}, {1, 1}}))

	// Multiple rects
	fmt.Println(maxRectangleArea([][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2},
		{0, 4}, {2, 4}, {0, 6}, {2, 6}}))

	// 5 valid rects (from description)
	fmt.Println(maxRectangleArea([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
}

type BIT struct {
	n    int
	tree []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n, tree: make([]int, n+2)}
}

func (b *BIT) add(idx int, val int) {
	idx++
	for idx <= b.n+1 {
		b.tree[idx] += val
		idx += idx & -idx
	}
}

func (b *BIT) sum(idx int) int {
	idx++
	res := 0
	for idx > 0 {
		res += b.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (b *BIT) rangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return b.sum(r) - b.sum(l-1)
}

func maxRectangleArea(points [][]int) int {
	n := len(points)
	if n < 4 {
		return -1
	}

	// Coordinate compression for y
  // Alokasi slice integer
	yVals := make([]int, n)
	for i, p := range points {
		yVals[i] = p[1]
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(yVals)
  // Alokasi slice integer
	uniqY := make([]int, 0)
	for i, y := range yVals {
		if i == 0 || y != yVals[i-1] {
			uniqY = append(uniqY, y)
		}
	}
  // Membuat map (HashMap) — pencarian O(1)
	yComp := make(map[int]int)
	for i, y := range uniqY {
		yComp[y] = i
	}

	// Sort points by x, then y
  // Alokasi slice integer
	sorted := make([][2]int, n)
	for i, p := range points {
		sorted[i] = [2]int{p[0], p[1]}
	}
  // Custom sort dengan comparator
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i][0] != sorted[j][0] {
			return sorted[i][0] < sorted[j][0]
		}
		return sorted[i][1] < sorted[j][1]
	})

	bit := NewBIT(len(uniqY))
	// key = (y1, y2) -> [x_of_left_side, point_count_at_that_x]
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[[2]int][2]int)
	ans := -1

	i := 0
	for i < n {
		x := sorted[i][0]
		// Process all points with this x
		j := i
		for j < n && sorted[j][0] == x {
			y := sorted[j][1]
			bit.add(yComp[y], 1)
			j++
		}

		// Check consecutive pairs at this x
		for k := i; k+1 < j; k++ {
			y1 := sorted[k][1]
			y2 := sorted[k+1][1]
			if y1 >= y2 {
				continue
			}
			// Same y-pair seen before?
			key := [2]int{yComp[y1], yComp[y2]}
			if prev, ok := seen[key]; ok {
				prevX := prev[0]
				prevCnt := prev[1]
				curCnt := bit.rangeSum(yComp[y1], yComp[y2])
				// If exactly 2 new points added (the two right corners)
				if curCnt == prevCnt+2 {
					area := (x - prevX) * (y2 - y1)
					if area > ans {
						ans = area
					}
				}
			}
			seen[key] = [2]int{x, bit.rangeSum(yComp[y1], yComp[y2])}
		}

		i = j
	}

	return ans
}
```
