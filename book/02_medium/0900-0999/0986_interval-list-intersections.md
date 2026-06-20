# 0986 — Interval List Intersections

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int
```

> **💡 Hint:** Two pointers

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m + n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #986: Interval List Intersections
// https://leetcode.com/problems/interval-list-intersections/
// Difficulty: Medium
//
// Approach: Two pointers
// Time: O(m + n)
// Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	// [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	// []
	fmt.Println(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}))
	// [[3,7]]
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	result := [][]int{}

	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if start <= end {
			result = append(result, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	if len(result) == 0 {
		return [][]int{}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
