# 0613 — Shortest Distance In A Line

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ShortestDistanceInALine() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #613: Shortest Distance in a Line
// https://leetcode.com/problems/shortest-distance-in-a-line/
// Difficulty: Easy [Paid]

import "fmt"

func ShortestDistanceInALine() string {
	return "SELECT MIN(ABS(p1.x - p2.x)) AS shortest FROM Point p1 JOIN Point p2 ON p1.x != p2.x"
}

func main() {
	fmt.Println(ShortestDistanceInALine())
}
```
