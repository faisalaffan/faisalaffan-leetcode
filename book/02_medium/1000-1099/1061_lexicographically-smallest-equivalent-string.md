# 1061 — Lexicographically Smallest Equivalent String

## Deskripsi

**Soal:** [1061. Lexicographically Smallest Equivalent String](https://leetcode.com/problems/lexicographically-smallest-equivalent-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O((m + n) * alpha(26)) where m = len(s1), n = len(baseStr)  
**Kompleksitas Ruang:** O(26) = O(1)

**Algoritma:** —

> **Ide Kunci:** DSU (Union-Find) to group equivalent characters,

## Solusi Go

```go
package main

// LeetCode #1061: Lexicographically Smallest Equivalent String
// https://leetcode.com/problems/lexicographically-smallest-equivalent-string/
// Difficulty: Medium
//
// Approach: DSU (Union-Find) to group equivalent characters,
//           then replace each char in baseStr with its smallest root.
// Time: O((m + n) * alpha(26)) where m = len(s1), n = len(baseStr)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser")) // "makkek"
	fmt.Println(smallestEquivalentString("hello", "world", "hold"))     // "hdld"
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa < pb {
			parent[pb] = pa
		} else {
			parent[pa] = pb
		}
	}

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s1); i++ {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]byte, len(baseStr))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(baseStr); i++ {
		result[i] = byte('a' + find(int(baseStr[i]-'a')))
	}

	return string(result)
}
```
