# 0990 — Satisfiability Of Equality Equations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func equationsPossible(equations []string) bool
```

> **💡 Hint:** Union-Find (DSU)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Union-Find (DSU)

**Kompleksitas Waktu:** O(n * alpha(N)) where n = len(equations), alpha is inverse Ackermann  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Union-Find (DSU)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
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
