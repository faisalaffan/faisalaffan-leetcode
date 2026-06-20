# 3406 — Find The Lexicographically Largest String From The Box Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLexicographicallyLargestStringFromTheBoxIi(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3406: Find the Lexicographically Largest String From the Box II
// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-ii/
// Difficulty: Hard [Paid]
//
// Two-pointer front/back comparison with look-ahead on ties.

import "fmt"

func main() {
	fmt.Println(FindTheLexicographicallyLargestStringFromTheBoxIi("abcabc"))
	fmt.Println(FindTheLexicographicallyLargestStringFromTheBoxIi("acbac"))
}

func FindTheLexicographicallyLargestStringFromTheBoxIi(s string) string {
	n := len(s)
	i, j := 0, n-1
	var res []byte
	for i <= j {
		if s[i] > s[j] {
			res = append(res, s[i])
			i++
		} else if s[j] > s[i] {
			res = append(res, s[j])
			j--
		} else {
			// Tie: need to look ahead
			li, rj := i, j
			for li <= rj && s[li] == s[rj] {
				li++
				rj--
			}
			if li > rj || s[li] > s[rj] {
				res = append(res, s[i])
				i++
			} else {
				res = append(res, s[j])
				j--
			}
		}
	}
	return string(res)
}
```
