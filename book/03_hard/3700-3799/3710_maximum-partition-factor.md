# 3710 — Maximum Partition Factor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxPartitionFactor(points [][]int) int
```

> **💡 Hint:** Sort all pairwise Manhattan distances. Binary search + bipartite

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3710: Maximum Partition Factor
// https://leetcode.com/problems/maximum-partition-factor/
// Difficulty: Hard
//
// Split n points into two non-empty groups. Partition factor = minimum
// Manhattan distance between any two points in the same group.
// Maximize the partition factor.
//
// Approach: Sort all pairwise Manhattan distances. Binary search + bipartite
// check (color graph where distance < threshold, must be bipartite).

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(maxPartitionFactor([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}}))
	// Example 2
	fmt.Println(maxPartitionFactor([][]int{{0, 0}, {1, 1}, {2, 2}}))
	// Edge: two points
	fmt.Println(maxPartitionFactor([][]int{{0, 0}, {1, 1}}))
}

func maxPartitionFactor(points [][]int) int {
	n := len(points)
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 0 // each group has 1 element, no intra-group pairs
	}

	// Collect all distances
	type edge struct {
		u, v, d int
	}
	distances := make([]edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			distances = append(distances, edge{i, j, d})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].d < distances[j].d
	})

	// Check if we can achieve partition factor >= x
	// Two points with distance < x must be in different groups (bipartite)
	check := func(x int) bool {
  // Membuat matriks/slice 2D untuk DP
		adj := make([][]int, n)
		for _, e := range distances {
			if e.d >= x {
				break
			}
			adj[e.u] = append(adj[e.u], e.v)
			adj[e.v] = append(adj[e.v], e.u)
		}
  // Alokasi slice integer
		color := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range color {
			color[i] = -1
		}
		for i := 0; i < n; i++ {
			if color[i] == -1 {
				color[i] = 0
				queue := []int{i}
				for len(queue) > 0 {
					u := queue[0]
					queue = queue[1:]
					for _, v := range adj[u] {
						if color[v] == -1 {
							color[v] = color[u] ^ 1
							queue = append(queue, v)
						} else if color[v] == color[u] {
							return false
						}
					}
				}
			}
		}
		// Check both groups are non-empty
		group0, group1 := 0, 0
		for i := 0; i < n; i++ {
			if color[i] == 0 {
				group0++
			} else {
				group1++
			}
		}
		return group0 > 0 && group1 > 0
	}

	// Binary search on distance value
	left, right := 0, distances[len(distances)-1].d
	result := 0
	for left <= right {
		mid := (left + right) / 2
		if check(mid) {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
