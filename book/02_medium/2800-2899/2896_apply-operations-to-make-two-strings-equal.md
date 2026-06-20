# 2896 — Apply Operations To Make Two Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ApplyOperationsToMakeTwoStringsEqual(s1 string, s2 string, x int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2896: Apply Operations to Make Two Strings Equal
// https://leetcode.com/problems/apply-operations-to-make-two-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeTwoStringsEqual(s1 string, s2 string, x int) int {
	// Find positions where characters differ
  // Alokasi slice integer
	diff := make([]int, 0)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff)%2 != 0 {
		return -1
	}
	if len(diff) == 0 {
		return 0
	}

	m := len(diff)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l > r {
			return 0
		}
		if dp[l][r] != -1 {
			return dp[l][r]
		}

		// Option 1: flip s[l] and s[l+1] (cost = diff[l+1] - diff[l])
		best := solve(l+2, r) + diff[l+1] - diff[l]

		// Option 2: use operation with cost x
		cost := solve(l+1, r-1) + x
		if cost < best {
			best = cost
		}

		dp[l][r] = best
		return best
	}

	return solve(0, m-1)
}

func main() {
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("1100011000", "0101001010", 2))
	fmt.Println(ApplyOperationsToMakeTwoStringsEqual("10110", "00011", 4))
}
```
