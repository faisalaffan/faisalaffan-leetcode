# 1331 — Rank Transform Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func arrayRankTransform(arr []int) []int

import (
	"fmt"
	"sort"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1331: Rank Transform of an Array
// https://leetcode.com/problems/rank-transform-of-an-array/
// Difficulty: Easy
//
// LeetCode submission: func arrayRankTransform(arr []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RankTransformOfAnArray([]int{40, 10, 20, 30}))       // [4 1 2 3]
	fmt.Println(RankTransformOfAnArray([]int{100, 100, 100}))        // [1 1 1]
	fmt.Println(RankTransformOfAnArray([]int{37, 12, 28, 9, 100, 56})) // [5 3 4 1 6 2]
}

// Time: O(n log n), Space: O(n)
func RankTransformOfAnArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
  // Alokasi slice integer
	sorted := make([]int, len(arr))
	copy(sorted, arr)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)

  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]int, len(arr))
	cur := 1
	for _, v := range sorted {
		if _, seen := rank[v]; !seen {
			rank[v] = cur
			cur++
		}
	}

  // Alokasi slice integer
	res := make([]int, len(arr))
	for i, v := range arr {
		res[i] = rank[v]
	}
	return res
}
```
