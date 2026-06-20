# 0990 — Satisfiability Of Equality Equations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func equationsPossible(equations []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Union-Find

**Waktu:** O(n * alpha(N)) where n = len(equations), alpha is inverse Ackermann  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Union-Find** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #990: Satisfiability of Equality Equations
// https://leetcode.com/problems/satisfiability-of-equality-equations/
// Difficulty: Medium
//
// Approach: Union-Find (DSU)
// Time: O(n * alpha(N)) where n = len(equations), alpha is inverse Ackermann
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))                           // false
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))                           // true
	fmt.Println(equationsPossible([]string{"a==b", "b==c", "a==c"}))                   // true
	fmt.Println(equationsPossible([]string{"a==b", "b!=c", "c==a"}))                   // false
}

func equationsPossible(equations []string) bool {
  // Alokasi slice
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
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
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	// Process all equalities first
	for _, eq := range equations {
		if eq[1] == '=' {
			union(int(eq[0]-'a'), int(eq[3]-'a'))
		}
	}

	// Then check all inequalities
	for _, eq := range equations {
		if eq[1] == '!' {
			if find(int(eq[0]-'a')) == find(int(eq[3]-'a')) {
				return false
			}
		}
	}

	return true
}
```
