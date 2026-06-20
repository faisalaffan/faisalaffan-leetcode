# 2904 — Shortest And Lexicographically Smallest Beautiful String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestBeautifulSubstring(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2904: Shortest and Lexicographically Smallest Beautiful String
// https://leetcode.com/problems/shortest-and-lexicographically-smallest-beautiful-string/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(shortestBeautifulSubstring("100011001", 3))
	fmt.Println(shortestBeautifulSubstring("1011", 2))
	fmt.Println(shortestBeautifulSubstring("000", 1))
}

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		for j := i + k; j <= n; j++ {
			t := s[i:j]
			cnt := 0
			for _, c := range t {
				if c == '1' {
					cnt++
				}
			}
			if cnt == k && (ans == "" || j-i < len(ans) || (j-i == len(ans) && t < ans)) {
				ans = t
			}
		}
	}
	return ans
}
```
