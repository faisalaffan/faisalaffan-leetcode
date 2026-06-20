# 3302 — Find The Lexicographically Smallest Valid Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func validSequence(word1 string, word2 string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3302: Find the Lexicographically Smallest Valid Sequence
// https://leetcode.com/problems/find-the-lexicographically-smallest-valid-sequence/
// Difficulty: Medium
// Time: O(n + m) Space: O(n)

import "fmt"

func main() {
	fmt.Println(validSequence("abc", "ab"))   // [0 1]
	fmt.Println(validSequence("abc", "ad"))   // [0 2]
	fmt.Println(validSequence("abbc", "abc")) // [0 1 3]
}

func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
  // Alokasi slice integer
	suf := make([]int, n+1)
	suf[n] = m
	j := m - 1
	for i := n - 1; i >= 0; i-- {
		if j >= 0 && word1[i] == word2[j] {
			j--
		}
		suf[i] = j + 1
	}

	ans := []int{}
	changed := false
	j = 0
	for i := 0; i < n && j < m; i++ {
		if word1[i] == word2[j] {
			ans = append(ans, i)
			j++
		} else if !changed && suf[i+1] <= j+1 {
			changed = true
			ans = append(ans, i)
			j++
		}
	}

	if j < m {
		return []int{}
	}
	return ans
}
```
