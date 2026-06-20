# 2924 — Find Champion Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findChampionII(n int, edges [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2924: Find Champion II
// https://leetcode.com/problems/find-champion-ii/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(findChampionII(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampionII(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampionII(2, [][]int{{0, 1}}))
}

func findChampionII(n int, edges [][]int) int {
  // Alokasi slice integer
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i, x := range indeg {
		if x == 0 {
			cnt++
			ans = i
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
```
