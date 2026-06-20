# 1496 — Path Crossing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func isPathCrossing(path string) bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1496: Path Crossing
// https://leetcode.com/problems/path-crossing/
// Difficulty: Easy
//
// LeetCode submission: func isPathCrossing(path string) bool

import "fmt"

func main() {
	fmt.Println(PathCrossing("NES"))   // false
	fmt.Println(PathCrossing("NESWW")) // true
}

// Time: O(n), Space: O(n)
func PathCrossing(path string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	visited := make(map[[2]int]bool)
	x, y := 0, 0
	visited[[2]int{0, 0}] = true
	for _, ch := range path {
		switch ch {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if visited[[2]int{x, y}] {
			return true
		}
		visited[[2]int{x, y}] = true
	}
	return false
}
```
