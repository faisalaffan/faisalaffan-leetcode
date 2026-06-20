# 2767 — Partition String Into Minimum Beautiful Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func PartitionStringIntoMinimumBeautifulSubstrings(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2767: Partition String Into Minimum Beautiful Substrings
// https://leetcode.com/problems/partition-string-into-minimum-beautiful-substrings/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n)

import (
	"fmt"
	"math"
)

func PartitionStringIntoMinimumBeautifulSubstrings(s string) int {
	n := len(s)
  // Membuat map (HashMap) — pencarian O(1)
	powers := make(map[string]bool)
	for i := 0; i <= 10; i++ {
		p := int(math.Pow(5, float64(i)))
		b := fmt.Sprintf("%b", p)
		if len(b) <= 15 {
			powers[b] = true
		}
	}

  // Alokasi slice integer
	memo := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(start int) int {
		if start == n {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		best := math.MaxInt32
		for end := start + 1; end <= n; end++ {
			sub := s[start:end]
			if powers[sub] {
				subResult := dfs(end)
				if subResult != -1 && subResult+1 < best {
					best = subResult + 1
				}
			}
		}
		if best == math.MaxInt32 {
			memo[start] = -1
		} else {
			memo[start] = best
		}
		return memo[start]
	}

	result := dfs(0)
	if result == math.MaxInt32 || result <= 0 {
		return -1
	}
	return result
}

func main() {
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("1011"))
	fmt.Println(PartitionStringIntoMinimumBeautifulSubstrings("111"))
}
```
