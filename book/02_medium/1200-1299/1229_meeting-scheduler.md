# 1229 — Meeting Scheduler

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + m log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1229: Meeting Scheduler
// https://leetcode.com/problems/meeting-scheduler/
// Difficulty: Medium [Paid]

// Find earliest time slot of given duration that works for both.
// Two pointers approach after sorting slots by start time.

// Time: O(n log n + m log m)
// Space: O(1)

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
  // Custom sort dengan comparator
	sort.Slice(slots1, func(i, j int) bool { return slots1[i][0] < slots1[j][0] })
  // Custom sort dengan comparator
	sort.Slice(slots2, func(i, j int) bool { return slots2[i][0] < slots2[j][0] })

	i, j := 0, 0
	for i < len(slots1) && j < len(slots2) {
		start := max(slots1[i][0], slots2[j][0])
		end := min(slots1[i][1], slots2[j][1])
		if end-start >= duration {
			return []int{start, start + duration}
		}
		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}
	return []int{}
}

func main() {
	fmt.Printf("%v (expected: [60 68])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}, {140, 210}},
			[][]int{{0, 15}, {60, 70}}, 8))

	fmt.Printf("%v (expected: [])\n",
		minAvailableDuration([][]int{{10, 50}, {60, 120}},
			[][]int{{0, 15}, {55, 58}}, 5))
}
```
