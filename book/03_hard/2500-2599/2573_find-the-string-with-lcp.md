# 2573 — Find The String With Lcp

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func findTheString(lcp [][]int) string
```

> **💡 Hint:** // 1. Validate matrix (diagonal must be n-i, symmetric, values within bounds).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU), Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2573: Find the String with LCP
// https://leetcode.com/problems/find-the-string-with-lcp/
// Difficulty: Hard

import "fmt"

// findTheString reconstructs the lexicographically smallest string that matches
// the given LCP matrix. lcp[i][j] = longest common prefix of suffixes
// starting at i and j.
//
// Approach:
// 1. Validate matrix (diagonal must be n-i, symmetric, values within bounds).
// 2. If lcp[i][j] > 0, then s[i] must equal s[j]. Union via DSU.
// 3. Assign smallest possible character ('a', 'b', ...) to each equivalence class.
// 4. Verify by recomputing LCP from the constructed string.
//
// Complexity: O(n^2) time, O(n^2) space

func findTheString(lcp [][]int) string {
	n := len(lcp)

	// Validate matrix shape
	for i := 0; i < n; i++ {
		if len(lcp[i]) != n {
			return ""
		}
	}

	// Validate diagonal: lcp[i][i] must equal n-i
	for i := 0; i < n; i++ {
		if lcp[i][i] != n-i {
			return ""
		}
	}

	// Validate symmetry: lcp[i][j] == lcp[j][i]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lcp[i][j] != lcp[j][i] {
				return ""
			}
		}
	}

	// Validate bounds: lcp[i][j] <= n - max(i, j)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if lcp[i][j] > n-i || lcp[i][j] > n-j {
				return ""
			}
		}
	}

	// DSU to union positions that must have the same character
  // Alokasi slice integer
	parent := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			parent[y] = x
		}
	}

	// Union positions where lcp[i][j] > 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if lcp[i][j] > 0 {
				union(i, j)
			}
		}
	}

	// Assign characters: each equivalence class gets the smallest unused char
	ans := make([]byte, n)
	nextChar := byte('a')

	for i := 0; i < n; i++ {
		if ans[i] != 0 {
			continue
		}
		if nextChar > 'z' {
			return "" // too many distinct characters needed
		}
		// Assign this char to the entire equivalence class
		for j := i; j < n; j++ {
			if find(j) == find(i) {
				ans[j] = nextChar
			}
		}
		nextChar++
	}

	// Verify by recomputing LCP from the constructed string
  // Membuat matriks/slice 2D untuk DP
	computed := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range computed {
		computed[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if ans[i] != ans[j] {
				computed[i][j] = 0
			} else if i+1 < n && j+1 < n {
				computed[i][j] = computed[i+1][j+1] + 1
			} else {
				computed[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if computed[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(ans)
}

func main() {
	// Test cases
	lcp1 := [][]int{{4, 0, 2, 0}, {0, 3, 0, 1}, {2, 0, 2, 0}, {0, 1, 0, 1}}
	fmt.Println("Test 1: ->", findTheString(lcp1))

	lcp2 := [][]int{{4, 3, 2, 1}, {3, 3, 2, 1}, {2, 2, 2, 1}, {1, 1, 1, 1}}
	fmt.Println("Test 2: ->", findTheString(lcp2)) // "aaaa"

	lcp3 := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Test 3: ->", findTheString(lcp3)) // "ab"

	lcp4 := [][]int{{1, 1}, {1, 1}}
	fmt.Println("Test 4: ->", findTheString(lcp4)) // invalid

	lcp5 := [][]int{{3, 0, 1}, {0, 2, 0}, {1, 0, 1}}
	fmt.Println("Test 5: ->", findTheString(lcp5))

	lcp6 := [][]int{{1}}
	fmt.Println("Test 6: ->", findTheString(lcp6)) // "a"
}
```
