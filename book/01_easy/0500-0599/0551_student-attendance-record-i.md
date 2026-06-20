# 0551 — Student Attendance Record I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func StudentAttendanceRecordI(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #551: Student Attendance Record I
// https://leetcode.com/problems/student-attendance-record-i/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func StudentAttendanceRecordI(s string) bool {
	absences, lateStreak := 0, 0
	for _, c := range s {
		if c == 'A' {
			absences++
			lateStreak = 0
			if absences >= 2 {
				return false
			}
		} else if c == 'L' {
			lateStreak++
			if lateStreak >= 3 {
				return false
			}
		} else {
			lateStreak = 0
		}
	}
	return true
}

func main() {
	fmt.Println(StudentAttendanceRecordI("PPALLP"))
	fmt.Println(StudentAttendanceRecordI("PPALLL"))
}
```
