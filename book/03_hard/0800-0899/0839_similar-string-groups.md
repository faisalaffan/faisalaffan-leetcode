# 0839 — Similar String Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func numSimilarGroups(strs []string) int
```

> **💡 Hint:** Union-Find. Two strings are similar if they differ by exactly 0 or 2 characters

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #839: Similar String Groups
// https://leetcode.com/problems/similar-string-groups/
// Difficulty: Hard
// Approach: Union-Find. Two strings are similar if they differ by exactly 0 or 2 characters
// (i.e., swapping two positions makes them equal). Group connected components.

import "fmt"

func numSimilarGroups(strs []string) int {
	n := len(strs)
  // Alokasi slice integer
	parent := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	isSimilar := func(a, b string) bool {
		diff := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
				if diff > 2 {
					return false
				}
			}
		}
		return diff == 0 || diff == 2
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int]bool)
	for i := 0; i < n; i++ {
		groups[find(i)] = true
	}
	return len(groups)
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"})) // Expected: 2
	fmt.Println(numSimilarGroups([]string{"abc", "abc"}))                   // Expected: 1
	fmt.Println(numSimilarGroups([]string{"omv", "ovm"}))                   // Expected: 1
}
```
