# 3661 — Maximum Walls Destroyed By Robots

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxWalls(robots []int, distance []int, walls []int) int
```

> **💡 Hint:** Sort robots and walls. For each robot, count walls within

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3661: Maximum Walls Destroyed by Robots
// https://leetcode.com/problems/maximum-walls-destroyed-by-robots/
// Difficulty: Hard
//
// Given robots at positions with bullet ranges and walls at positions,
// find maximum number of unique walls a single robot can destroy.
//
// Approach: Sort robots and walls. For each robot, count walls within
// its reach (left and right). Track max.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(maxWalls([]int{1, 5, 10}, []int{2, 3, 4}, []int{2, 4, 6, 8, 11}))
	// Example 2
	fmt.Println(maxWalls([]int{3, 7}, []int{2, 2}, []int{1, 4, 6, 9}))
	// Edge: single robot
	fmt.Println(maxWalls([]int{5}, []int{3}, []int{1, 2, 6, 7, 8}))
	// Edge: no walls
	fmt.Println(maxWalls([]int{1}, []int{5}, []int{}))
}

func maxWalls(robots []int, distance []int, walls []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(walls)
	result := 0

	for i, r := range robots {
		d := distance[i]
		left := r - d
		right := r + d
		// Count walls in [left, right]
		l := sort.SearchInts(walls, left)
		rr := sort.SearchInts(walls, right+1)
		cnt := rr - l
		if cnt > result {
			result = cnt
		}
	}
	return result
}
```
