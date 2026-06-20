# 0277 — Find The Celebrity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func knows(a, b int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #277: Find the Celebrity
// https://leetcode.com/problems/find-the-celebrity/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

var knowsMatrix [][]int

func knows(a, b int) bool {
	return knowsMatrix[a][b] == 1
}

func findCelebrity(n int) int {
	candidate := 0

	for i := 1; i < n; i++ {
		if knows(candidate, i) {
			candidate = i
		}
	}

	for i := 0; i < n; i++ {
		if i == candidate {
			continue
		}
		if knows(candidate, i) || !knows(i, candidate) {
			return -1
		}
	}

	return candidate
}

func main() {
	knowsMatrix = [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(findCelebrity(3))

	knowsMatrix = [][]int{{1, 1}, {0, 1}}
	fmt.Println(findCelebrity(2))
}
```
