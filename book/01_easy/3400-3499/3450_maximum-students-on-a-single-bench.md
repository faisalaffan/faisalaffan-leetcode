# 3450 — Maximum Students On A Single Bench

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumStudentsOnASingleBench(students [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3450: Maximum Students on a Single Bench
// https://leetcode.com/problems/maximum-students-on-a-single-bench/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 2}, {2, 3}, {1, 3}, {1, 2}}))
	fmt.Println(MaximumStudentsOnASingleBench([][]int{{1, 1}, {2, 1}, {3, 1}}))
}

// MaximumStudentsOnASingleBench returns the maximum number of different students on any single bench.
// Each entry is [student_id, bench_id].
// Time: O(n). Space: O(n).
func MaximumStudentsOnASingleBench(students [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	benchStudents := make(map[int]map[int]bool)
	for _, s := range students {
		studentID, benchID := s[0], s[1]
		if benchStudents[benchID] == nil {
			benchStudents[benchID] = make(map[int]bool)
		}
		benchStudents[benchID][studentID] = true
	}
	maxCount := 0
	for _, students := range benchStudents {
		if len(students) > maxCount {
			maxCount = len(students)
		}
	}
	return maxCount
}
```
