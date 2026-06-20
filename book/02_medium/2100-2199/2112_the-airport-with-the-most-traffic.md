# 2112 — The Airport With The Most Traffic

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func airportWithMostTraffic(flights [][]string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2112: The Airport With the Most Traffic
// https://leetcode.com/problems/the-airport-with-the-most-traffic/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func airportWithMostTraffic(flights [][]string) string {
  // HashMap: O(1) lookup
	traffic := make(map[string]int)
	for _, f := range flights {
		departure, arrival := f[0], f[1]
		traffic[departure]++
		traffic[arrival]++
	}

	maxTraffic := 0
	airport := ""
	for a, t := range traffic {
		if t > maxTraffic || (t == maxTraffic && (airport == "" || a < airport)) {
			maxTraffic = t
			airport = a
		}
	}
	return airport
}

func main() {
	// Test case 1
	flights1 := [][]string{{"JFK", "LGA"}, {"JFK", "LAX"}, {"LAX", "SFO"}, {"LGA", "ORD"}}
	fmt.Println("Test 1:", airportWithMostTraffic(flights1))
	// Expected: "JFK" (3 flights)

	// Test case 2
	flights2 := [][]string{{"A", "B"}, {"B", "C"}, {"C", "A"}}
	fmt.Println("Test 2:", airportWithMostTraffic(flights2))
	// Expected: "A" (tie, alphabetically first)

	// Test case 3
	fmt.Println("Test 3:", airportWithMostTraffic([][]string{{"X", "Y"}}))
	// Expected: "X" (tie, alphabetically: X < Y)
}
```
