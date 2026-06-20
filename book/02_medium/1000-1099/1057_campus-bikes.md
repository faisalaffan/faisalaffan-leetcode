# 1057 — Campus Bikes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func assignBikes(workers [][]int, bikes [][]int) []int
```

> **💡 Hint:** Bucket sort by Manhattan distance. Assign closest pairs.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(W * B) where W = workers, B = bikes  
**Kompleksitas Ruang:** O(W * B)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1057: Campus Bikes
// https://leetcode.com/problems/campus-bikes/
// Difficulty: Medium
//
// Approach: Bucket sort by Manhattan distance. Assign closest pairs.
// Time: O(W * B) where W = workers, B = bikes
// Space: O(W * B)

import "fmt"

func main() {
	fmt.Println(assignBikes([][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}})) // [1,0]
	fmt.Println(assignBikes([][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}})) // [0,2,1]
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	w, b := len(workers), len(bikes)
	// Buckets of distances: max distance is 2000 (0-1000, 0-1000)
	maxDist := 2000
  // Membuat matriks/slice 2D untuk DP
	buckets := make([][][2]int, maxDist+1)

	for i := 0; i < w; i++ {
		for j := 0; j < b; j++ {
			dist := abs(workers[i][0]-bikes[j][0]) + abs(workers[i][1]-bikes[j][1])
			buckets[dist] = append(buckets[dist], [2]int{i, j})
		}
	}

  // Alokasi slice integer
	result := make([]int, w)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = -1
	}
	bikeUsed := make([]bool, b)

	for d := 0; d <= maxDist; d++ {
		for _, pair := range buckets[d] {
			wi, bi := pair[0], pair[1]
			if result[wi] == -1 && !bikeUsed[bi] {
				result[wi] = bi
				bikeUsed[bi] = true
			}
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
