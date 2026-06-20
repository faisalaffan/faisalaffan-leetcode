# 0252 — Meeting Rooms

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CanAttendMeetings(intervals [][]int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #252: Meeting Rooms
// https://leetcode.com/problems/meeting-rooms/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
  // Custom sort dengan comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}))
	fmt.Println(CanAttendMeetings([][]int{{7, 10}, {2, 4}}))
}
```
