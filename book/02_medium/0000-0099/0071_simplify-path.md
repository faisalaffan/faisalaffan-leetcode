# 0071 — Simplify Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func simplifyPath(path string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #71: Simplify Path
// https://leetcode.com/problems/simplify-path/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}

	return "/" + strings.Join(stack, "/")
}

func main() {
	// Test case 1
	fmt.Println(simplifyPath("/home/")) // "/home"

	// Test case 2
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"

	// Test case 3
	fmt.Println(simplifyPath("/../")) // "/"

	// Test case 4
	fmt.Println(simplifyPath("/a/./b/../../c/")) // "/c"
}

// Time: O(n) | Space: O(n)
```
