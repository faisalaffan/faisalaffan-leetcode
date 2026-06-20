# 0253 — Meeting Rooms Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMeetingRooms(intervals [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #253: Meeting Rooms II
// https://leetcode.com/problems/meeting-rooms-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

  // Alokasi slice integer
	starts := make([]int, len(intervals))
  // Alokasi slice integer
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(starts)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(ends)

	rooms, endIdx := 0, 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(starts); i++ {
		if starts[i] < ends[endIdx] {
			rooms++
		} else {
			endIdx++
		}
	}

	return rooms
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))
	fmt.Println(minMeetingRooms([][]int{{0, 5}, {5, 10}, {10, 15}}))
}
```
