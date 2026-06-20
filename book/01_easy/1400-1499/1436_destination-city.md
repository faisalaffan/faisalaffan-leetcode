# 1436 — Destination City

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func destCity(paths [][]string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1436: Destination City
// https://leetcode.com/problems/destination-city/
// Difficulty: Easy
//
// LeetCode submission: func destCity(paths [][]string) string

import "fmt"

func main() {
	paths1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(DestinationCity(paths1)) // "Sao Paulo"

	paths2 := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(DestinationCity(paths2)) // "A"
}

// Time: O(n), Space: O(n)
func DestinationCity(paths [][]string) string {
  // HashMap: O(1) lookup
	outgoing := make(map[string]bool, len(paths))
	for _, p := range paths {
		outgoing[p[0]] = true
	}
	for _, p := range paths {
		if !outgoing[p[1]] {
			return p[1]
		}
	}
	return ""
}
```
