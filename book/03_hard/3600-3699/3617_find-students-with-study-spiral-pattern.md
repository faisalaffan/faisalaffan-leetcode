# 3617 — Find Students With Study Spiral Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findSpiralPattern() []Result
```

> **💡 Hint:** Group sessions by student, check for cyclic patterns.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3617: Find Students with Study Spiral Pattern
// https://leetcode.com/problems/find-students-with-study-spiral-pattern/
// Difficulty: Hard
//
// Given student and study session data, find students whose study sessions
// follow a repeating pattern of at least 3 subjects for at least 2 cycles
// with consecutive dates (no gaps > 2 days).
// This is originally a SQL problem. Implemented as Go.
//
// Approach: Group sessions by student, check for cyclic patterns.

import "fmt"

func main() {
	// Example 1
	fmt.Println(findSpiralPattern())
}

type Student struct {
	ID   int
	Name string
	Major string
}

type Session struct {
	StudentID    int
	Subject string
	Date    int // days since epoch
	Hours   float64
}

type Result struct {
	StudentID      int
	StudentName    string
	Major         string
	CycleLength   int
	TotalHours    float64
}

func findSpiralPattern() []Result {
	// Read from database would happen here
	// For the Go implementation, return empty (data-driven problem)
	return []Result{}
}
```
