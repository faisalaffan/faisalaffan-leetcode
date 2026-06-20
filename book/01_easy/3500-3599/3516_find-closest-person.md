# 3516 — Find Closest Person

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindClosestPerson(x int, y int, z int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3516: Find Closest Person
// https://leetcode.com/problems/find-closest-person/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestPerson(1, 2, 3))
	fmt.Println(FindClosestPerson(1, 3, 2))
	fmt.Println(FindClosestPerson(2, 1, 3))
}

// FindClosestPerson finds who is closer to person z. Returns 0 for tie, 1 for person 1, 2 for person 2.
// Time: O(1). Space: O(1).
func FindClosestPerson(x int, y int, z int) int {
	dx := z - x
	if dx < 0 {
		dx = -dx
	}
	dy := z - y
	if dy < 0 {
		dy = -dy
	}
	if dx < dy {
		return 1
	} else if dy < dx {
		return 2
	}
	return 0
}
```
