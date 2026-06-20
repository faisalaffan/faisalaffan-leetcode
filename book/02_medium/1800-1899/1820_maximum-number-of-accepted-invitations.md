# 1820 — Maximum Number Of Accepted Invitations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumInvitations(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(m * n^2) using DFS for bipartite matching  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1820: Maximum Number of Accepted Invitations
// https://leetcode.com/problems/maximum-number-of-accepted-invitations/
// Difficulty: Medium [Paid]
// Time: O(m * n^2) using DFS for bipartite matching

import "fmt"

func maximumInvitations(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
  // Alokasi slice integer
	match := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range match {
		match[i] = -1
	}

	var dfs func(u int, seen []bool) bool
	dfs = func(u int, seen []bool) bool {
		for v := 0; v < n; v++ {
			if grid[u][v] == 1 && !seen[v] {
				seen[v] = true
				if match[v] == -1 || dfs(match[v], seen) {
					match[v] = u
					return true
				}
			}
		}
		return false
	}

	result := 0
	for u := 0; u < m; u++ {
		seen := make([]bool, n)
		if dfs(u, seen) {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(maximumInvitations([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 1}})) // Expected: 3
	fmt.Println(maximumInvitations([][]int{{1, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 1, 0}, {1, 1, 1, 0}})) // Expected: 3
}
```
