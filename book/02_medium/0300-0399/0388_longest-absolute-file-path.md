# 0388 — Longest Absolute File Path

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthLongestPath(input string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #388: Longest Absolute File Path
// https://leetcode.com/problems/longest-absolute-file-path/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func lengthLongestPath(input string) int {
	lines := strings.Split(input, "\n")
  // Alokasi slice integer
	stack := make([]int, 0) // lengths at each depth
	maxLen := 0

	for _, line := range lines {
		// Count tabs to determine depth
		depth := 0
		for depth < len(line) && line[depth] == '\t' {
			depth++
		}

		// Current name without tabs
		name := line[depth:]

		// Pop stack to correct depth
		for len(stack) > depth {
			stack = stack[:len(stack)-1]
		}

		// Total length of current path
		total := len(name)
		if len(stack) > 0 {
			total += stack[len(stack)-1] + 1 // +1 for '/'
		}

		// Check if it's a file
		if strings.Contains(name, ".") {
			if total > maxLen {
				maxLen = total
			}
		}

		stack = append(stack, total)
	}
	return maxLen
}

func main() {
	// Test case 1
	input1 := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println("Test 1:", lengthLongestPath(input1))
	// Expected: 20 ("dir/subdir2/file.ext")

	// Test case 2
	input2 := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	fmt.Println("Test 2:", lengthLongestPath(input2))
	// Expected: 32 ("dir/subdir2/subsubdir2/file2.ext")

	// Test case 3: No files
	fmt.Println("Test 3:", lengthLongestPath("dir\n\tsubdir"))
	// Expected: 0
}
```
