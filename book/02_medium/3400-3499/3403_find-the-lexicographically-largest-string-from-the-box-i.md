# 3403 — Find The Lexicographically Largest String From The Box I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func answerString(word string, numFriends int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3403: Find the Lexicographically Largest String From the Box I
// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word)
	maxLen := n - numFriends + 1
	ans := word[:maxLen]
	for i := 0; i < n; i++ {
		end := i + maxLen
		if end > n {
			end = n
		}
		sub := word[i:end]
		if sub > ans {
			ans = sub
		}
	}
	return ans
}

func main() {
	fmt.Println(answerString("dbca", 2)) // "dbc"
	fmt.Println(answerString("gggg", 2)) // "ggg"
	fmt.Println(answerString("abc", 3))  // "c"
}
```
