# 0551 — Student Attendance Record I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func StudentAttendanceRecordI(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


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
