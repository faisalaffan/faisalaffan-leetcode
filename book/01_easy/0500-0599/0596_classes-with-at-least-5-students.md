# 0596 — Classes With At Least 5 Students

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ClassesWithAtLeastFiveStudents() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #596: Classes With at Least 5 Students
// https://leetcode.com/problems/classes-with-at-least-5-students/
// Difficulty: Easy

import "fmt"

func ClassesWithAtLeastFiveStudents() string {
	return "SELECT class FROM Courses GROUP BY class HAVING COUNT(student) >= 5"
}

func main() {
	fmt.Println(ClassesWithAtLeastFiveStudents())
}
```
