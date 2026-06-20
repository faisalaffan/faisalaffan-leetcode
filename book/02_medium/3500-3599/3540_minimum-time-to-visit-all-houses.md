# 3540 — Minimum Time To Visit All Houses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumTimeToVisitAllHouses(houses [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3540: Minimum Time to Visit All Houses
// https://leetcode.com/problems/minimum-time-to-visit-all-houses/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	houses := [][]int{{0, 0}, {1, 1}, {2, 2}}
	fmt.Println("Test 1:", MinimumTimeToVisitAllHouses(houses))
	// Test case 2
	houses2 := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println("Test 2:", MinimumTimeToVisitAllHouses(houses2))
	// Test case 3
	houses3 := [][]int{{0, 0}}
	fmt.Println("Test 3:", MinimumTimeToVisitAllHouses(houses3))
}

func MinimumTimeToVisitAllHouses(houses [][]int) int {
	if len(houses) == 0 {
		return 0
	}
	time := 0
	for i := 1; i < len(houses); i++ {
		dx := houses[i][0] - houses[i-1][0]
		if dx < 0 {
			dx = -dx
		}
		dy := houses[i][1] - houses[i-1][1]
		if dy < 0 {
			dy = -dy
		}
		// Can move diagonally, so time = max(dx, dy)
		if dx > dy {
			time += dx
		} else {
			time += dy
		}
	}
	return time
}
```
