# 3170 — Lexicographically Minimum String After Removing Stars

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func clearStars(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** O(n * 26)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3170: Lexicographically Minimum String After Removing Stars
// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)

import "fmt"

func clearStars(s string) string {
	n := len(s)
	bytes := []byte(s)
  // Membuat matriks/slice 2D untuk DP
	queues := make([][]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range queues {
		queues[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(queues[j]) > 0 {
					idx := queues[j][len(queues[j])-1]
					queues[j] = queues[j][:len(queues[j])-1]
					bytes[idx] = '*'
					break
				}
			}
			bytes[i] = '*'
		} else {
			queues[s[i]-'a'] = append(queues[s[i]-'a'], i)
		}
	}

	ans := make([]byte, 0, n)
	for _, ch := range bytes {
		if ch != '*' {
			ans = append(ans, ch)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(clearStars("aaba*"))       // Expected: "aab"
	fmt.Println(clearStars("abc"))          // Expected: "abc"
	fmt.Println(clearStars("a*b*c*"))       // Expected: ""
}
```
