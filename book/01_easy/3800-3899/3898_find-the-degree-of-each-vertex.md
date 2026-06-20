# 3898 — Find The Degree Of Each Vertex

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheDegreeOfEachVertex(matrix [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3898: Find the Degree of Each Vertex
// https://leetcode.com/problems/find-the-degree-of-each-vertex/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}))
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}))
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0}}))
}

// Time: O(n^2)
// Space: O(n)
func FindTheDegreeOfEachVertex(matrix [][]int) []int {
  // Alokasi slice integer
	ans := make([]int, len(matrix))
	for i, row := range matrix {
		sum := 0
		for _, v := range row {
			sum += v
		}
		ans[i] = sum
	}
	return ans
}
```
