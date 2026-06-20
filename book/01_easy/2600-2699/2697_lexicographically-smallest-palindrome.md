# 2697 — Lexicographically Smallest Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func LexicographicallySmallestPalindrome(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2697: Lexicographically Smallest Palindrome
// https://leetcode.com/problems/lexicographically-smallest-palindrome/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestPalindrome("egcfe"))
	fmt.Println(LexicographicallySmallestPalindrome("abcd"))
}

func LexicographicallySmallestPalindrome(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] < runes[j] {
			runes[j] = runes[i]
		} else {
			runes[i] = runes[j]
		}
		i++
		j--
	}
	return string(runes)
}
```
