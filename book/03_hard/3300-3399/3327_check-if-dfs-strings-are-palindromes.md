# 3327 — Check If Dfs Strings Are Palindromes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkIfDfsStringsArePalindromes(parent []int, s string) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3327: Check if DFS Strings Are Palindromes
// https://leetcode.com/problems/check-if-dfs-strings-are-palindromes/
// Difficulty: Hard
//
// Post-order DFS traversal produces a global string. Each subtree maps to a
// contiguous substring. Use rolling hash (forward + reverse) to check if each
// subtree's substring is a palindrome in O(n) total.

import "fmt"

func main() {
	// Example: parent=[-1,0,0,1,1,2], s="abccba" -> [true,true,true,true,true,true]
	fmt.Println(checkIfDfsStringsArePalindromes([]int{-1, 0, 0, 1, 1, 2}, "abccba"))

	// Single node
	fmt.Println(checkIfDfsStringsArePalindromes([]int{-1}, "a"))

	// Chain of 3: parent=[-1,0,1], s="aba" -> [true,true,true]
	fmt.Println(checkIfDfsStringsArePalindromes([]int{-1, 0, 1}, "aba"))

	// Chain of 3: s="abc" -> [true,true,false]
	fmt.Println(checkIfDfsStringsArePalindromes([]int{-1, 0, 1}, "abc"))

	// Binary tree: parent=[-1,0,0], s="aba"
	fmt.Println(checkIfDfsStringsArePalindromes([]int{-1, 0, 0}, "aba"))
}

const MOD = 1000000007
const BASE = 91138233

func checkIfDfsStringsArePalindromes(parent []int, s string) []bool {
	n := len(parent)
  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		children[p] = append(children[p], i)
	}

	// Post-order traversal to build the global string and record [start, end) for each node
	order := make([]byte, 0, n)
  // Alokasi slice integer
	start := make([]int, n)
  // Alokasi slice integer
	end := make([]int, n)

	var dfs func(u int)
	dfs = func(u int) {
		start[u] = len(order)
		for _, v := range children[u] {
			dfs(v)
		}
		order = append(order, s[u])
		end[u] = len(order)
	}
	dfs(0)

	// Rolling hash precomputation
	m := len(order)
  // Alokasi slice integer
	pow := make([]int64, m+1)
  // Alokasi slice integer
	fwd := make([]int64, m+1)
  // Alokasi slice integer
	rev := make([]int64, m+1)
	pow[0] = 1
	for i := 0; i < m; i++ {
		pow[i+1] = pow[i] * BASE % MOD
		v := int64(order[i] - 'a' + 1)
		fwd[i+1] = (fwd[i]*BASE + v) % MOD
	}
	for i := m - 1; i >= 0; i-- {
		v := int64(order[i] - 'a' + 1)
		rev[i] = (rev[i+1]*BASE + v) % MOD
	}

	// Hash functions
	getFwd := func(l, r int) int64 {
		return (fwd[r] - fwd[l]*pow[r-l]%MOD + MOD) % MOD
	}
	getRev := func(l, r int) int64 {
		return (rev[l] - rev[r]*pow[r-l]%MOD + MOD) % MOD
	}

	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		l, r := start[i], end[i]
		ans[i] = getFwd(l, r) == getRev(l, r)
	}
	return ans
}
```
