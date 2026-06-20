# 2363 — Merge Similar Items

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Merge Sort

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2363: Merge Similar Items
// https://leetcode.com/problems/merge-similar-items/
// Difficulty: Easy
// Time O((n+m) log(n+m)) | Space O(n+m)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}})) // [[1,6],[3,9],[4,5]]
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}})) // [[1,4],[2,4],[3,4]]
}

func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueMap := map[int]int{}
	for _, item := range items1 {
		valueMap[item[0]] += item[1]
	}
	for _, item := range items2 {
		valueMap[item[0]] += item[1]
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0, len(valueMap))
	for v, w := range valueMap {
		result = append(result, []int{v, w})
	}
  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })
	return result
}
```
