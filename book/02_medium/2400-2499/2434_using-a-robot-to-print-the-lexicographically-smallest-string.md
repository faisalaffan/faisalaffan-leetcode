# 2434 — Using A Robot To Print The Lexicographically Smallest String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func robotWithString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n * 26)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2434: Using a Robot to Print the Lexicographically Smallest String
// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)
// Track suffix minimum char. Push to stack if remaining suffix has smaller char.
// Otherwise pop from stack.

import "fmt"

func main() {
	fmt.Println(robotWithString("bac"))   // "abc"
	fmt.Println(robotWithString("bdda"))  // "addb"
}

func robotWithString(s string) string {
	n := len(s)
	suffixMin := make([]byte, n+1)
	suffixMin[n] = 'z' + 1
	for i := n - 1; i >= 0; i-- {
		if s[i] < suffixMin[i+1] {
			suffixMin[i] = s[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	stack := make([]byte, 0, n)
	res := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		for len(stack) > 0 && stack[len(stack)-1] <= suffixMin[i+1] {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(res)
}
```
