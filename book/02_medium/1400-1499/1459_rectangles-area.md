# 1459 — Rectangles Area

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func rectanglesArea(points []struct {
	id   int
	x, y int
}) []rectResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2) for finding all vertical pairs  
**Kompleksitas Ruang:** O(n^2) for map

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1459: Rectangles Area
// https://leetcode.com/problems/rectangles-area/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := rectanglesArea([]struct {
		id   int
		x, y int
	}{
		{1, 0, 0},
		{2, 2, 0},
		{3, 0, 2},
		{4, 2, 2},
		{5, 1, 1},
		{6, 3, 1},
		{7, 3, 3},
		{8, 1, 3},
	})
	for _, r := range result {
		fmt.Printf("%d,%d,%d,%d %d\n", r.p1, r.p2, r.p3, r.p4, r.area)
	}
}

type rectResult struct {
	p1, p2, p3, p4, area int
}

// Time: O(n^2) for finding all vertical pairs
// Space: O(n^2) for map
func rectanglesArea(points []struct {
	id   int
	x, y int
}) []rectResult {
	// Map from (y1,y2) -> list of x values that have both these y coords
	type yPair struct{ y1, y2 int }

	// Group points by x
  // Membuat map (HashMap) — pencarian O(1)
	pointsByX := make(map[int][]int) // x -> [ids]
  // Membuat map (HashMap) — pencarian O(1)
	coords := make(map[int]struct{ x, y int })

	for _, p := range points {
		pointsByX[p.x] = append(pointsByX[p.x], p.id)
		coords[p.id] = struct{ x, y int }{p.x, p.y}
	}

  // Membuat map (HashMap) — pencarian O(1)
	verticals := make(map[yPair][]int)

	for _, ids := range pointsByX {
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(ids); i++ {
			for j := i + 1; j < len(ids); j++ {
				yi := coords[ids[i]].y
				yj := coords[ids[j]].y
				var y1, y2 int
				if yi < yj {
					y1, y2 = yi, yj
				} else {
					y1, y2 = yj, yi
				}
				x := coords[ids[i]].x
				key := yPair{y1, y2}
				verticals[key] = append(verticals[key], x)
			}
		}
	}

	var results []rectResult
	for yp, xs := range verticals {
		// Sort xs
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(xs); i++ {
			for j := i + 1; j < len(xs); j++ {
				if xs[j] < xs[i] {
					xs[i], xs[j] = xs[j], xs[i]
				}
			}
		}
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(xs); i++ {
			for j := i + 1; j < len(xs); j++ {
				x1, x2 := xs[i], xs[j]
				area := (x2 - x1) * (yp.y2 - yp.y1)
				if area > 0 {
					var p1, p2, p3, p4 int
					for _, pt := range points {
						if pt.x == x1 && pt.y == yp.y1 {
							p1 = pt.id
						} else if pt.x == x2 && pt.y == yp.y1 {
							p2 = pt.id
						} else if pt.x == x1 && pt.y == yp.y2 {
							p3 = pt.id
						} else if pt.x == x2 && pt.y == yp.y2 {
							p4 = pt.id
						}
					}
					results = append(results, rectResult{p1, p2, p3, p4, area})
				}
			}
		}
	}

	return results
}
```
