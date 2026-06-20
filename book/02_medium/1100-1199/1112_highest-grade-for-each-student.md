# 1112 — Highest Grade For Each Student

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func highestGradeForEachStudent(enrollments [][]int) [][]int
```

> **💡 Hint:** Track best grade (highest, then earliest course_id) per student

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(m) where m = unique students

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1112: Highest Grade For Each Student
// https://leetcode.com/problems/highest-grade-for-each-student/
// Difficulty: Medium
//
// Approach: Track best grade (highest, then earliest course_id) per student
// Time: O(n)
// Space: O(m) where m = unique students

import "fmt"

func main() {
	// (student_id, course_id, grade)
	enrollments := [][]int{{1, 1, 90}, {1, 2, 95}, {2, 1, 85}, {2, 2, 85}, {3, 1, 70}}
	fmt.Println(highestGradeForEachStudent(enrollments))
}

func highestGradeForEachStudent(enrollments [][]int) [][]int {
	type best struct {
		grade    int
		courseID int
	}
  // Membuat map (HashMap) — pencarian O(1)
	bestMap := make(map[int]best)

	for _, e := range enrollments {
		sid, cid, grade := e[0], e[1], e[2]
		if existing, ok := bestMap[sid]; !ok || grade > existing.grade || (grade == existing.grade && cid < existing.courseID) {
			bestMap[sid] = best{grade, cid}
		}
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0, len(bestMap))
	for sid, b := range bestMap {
		result = append(result, []int{sid, b.courseID, b.grade})
	}

	return result
}
```
