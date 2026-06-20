# 2015 — Average Height Of Buildings In Each Segment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func AverageHeightOfBuildingsInEachSegment(buildings [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2015: Average Height of Buildings in Each Segment
// https://leetcode.com/problems/average-height-of-buildings-in-each-segment/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 4, 2}, {3, 9, 4}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 3, 2}, {2, 5, 3}, {2, 8, 3}}))
	fmt.Println(AverageHeightOfBuildingsInEachSegment([][]int{{1, 2, 1}, {5, 6, 1}}))
}

// Time: O(n log n), Space: O(n)
func AverageHeightOfBuildingsInEachSegment(buildings [][]int) [][]int {
	type event struct {
		pos    int
		height int
		change int
	}

	events := make([]event, 0)
	for _, b := range buildings {
		start, end, height := b[0], b[1], b[2]
		events = append(events, event{start, height, 1})
		events = append(events, event{end, height, -1})
	}

  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos != events[j].pos {
			return events[i].pos < events[j].pos
		}
		return events[i].height < events[j].height
	})

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)
	totalHeight := 0
	buildingCount := 0
	prevPos := -1

	for _, e := range events {
		if prevPos != -1 && prevPos < e.pos && buildingCount > 0 {
			avg := totalHeight / buildingCount
			n := len(result)
			if n > 0 && result[n-1][2] == avg && result[n-1][1] == prevPos {
				result[n-1][1] = e.pos
			} else {
				result = append(result, []int{prevPos, e.pos, avg})
			}
		}

		totalHeight += e.height * e.change
		buildingCount += e.change
		prevPos = e.pos
	}

	return result
}
```
