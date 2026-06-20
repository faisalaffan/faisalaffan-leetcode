# 3809 — Best Reachable Tower

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BestReachableTower(towers [][]int, center []int, radius int) []int
```

> **💡 Hint:** One-pass scan checking Manhattan distance and comparing quality.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3809: Best Reachable Tower
// https://leetcode.com/problems/best-reachable-tower/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: One-pass scan checking Manhattan distance and comparing quality.

import "fmt"

func BestReachableTower(towers [][]int, center []int, radius int) []int {
	cx, cy := center[0], center[1]
	bestQuality := -1
	bestX, bestY := -1, -1

	for _, t := range towers {
		x, y, q := t[0], t[1], t[2]
		dist := abs(x-cx) + abs(y-cy)
		if dist > radius {
			continue
		}
		if q > bestQuality || (q == bestQuality && (x < bestX || (x == bestX && y < bestY))) {
			bestQuality = q
			bestX, bestY = x, y
		}
	}

	if bestQuality == -1 {
		return []int{-1, -1}
	}
	return []int{bestX, bestY}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(BestReachableTower([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, []int{1, 1}, 2)) // Expected: [3, 1]

	// Example 2
	fmt.Println(BestReachableTower([][]int{{1, 3, 4}, {2, 2, 4}, {4, 4, 7}}, []int{0, 0}, 5)) // Expected: [1, 3]

	// Example 3
	fmt.Println(BestReachableTower([][]int{{5, 6, 8}, {0, 3, 5}}, []int{1, 2}, 1)) // Expected: [-1, -1]
}
```
