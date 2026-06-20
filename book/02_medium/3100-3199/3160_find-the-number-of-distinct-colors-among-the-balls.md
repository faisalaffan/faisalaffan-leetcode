# 3160 — Find The Number Of Distinct Colors Among The Balls

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func queryResults(limit int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3160: Find the Number of Distinct Colors Among the Balls
// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func queryResults(limit int, queries [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	ballColor := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	colorCount := make(map[int]int)
  // Alokasi slice integer
	ans := make([]int, len(queries))

	for i, q := range queries {
		ball, color := q[0], q[1]

		if prev, ok := ballColor[ball]; ok {
			colorCount[prev]--
			if colorCount[prev] == 0 {
				delete(colorCount, prev)
			}
		}

		ballColor[ball] = color
		colorCount[color]++
		ans[i] = len(colorCount)
	}
	return ans
}

func main() {
	fmt.Println(queryResults(4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}})) // Expected: [1, 2, 2, 3]
	fmt.Println(queryResults(4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}})) // Expected: [1, 2, 2, 3, 4]
}
```
