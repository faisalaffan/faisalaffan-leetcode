# 0899 — Orderly Queue

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func orderlyQueue(s string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #899: Orderly Queue
// https://leetcode.com/problems/orderly-queue/
// Difficulty: Hard
//
// If k > 1: we can reorder arbitrarily (bubble-sort style swaps), so return sorted string.
// If k == 1: only rotation is possible. Find the lexicographically smallest rotation.

import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k > 1 {
		b := []byte(s)
  // Custom sort dengan comparator
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		return string(b)
	}
	// k == 1: try all rotations.
	result := s
	n := len(s)
	for i := 1; i < n; i++ {
		rot := s[i:] + s[:i]
		if rot < result {
			result = rot
		}
	}
	return result
}

func main() {
	// Example 1: "cba", k=1 -> "acb" (rotate once: "bac", "acb", "cba")
	fmt.Println("Test 1:", orderlyQueue("cba", 1)) // "acb"

	// Example 2: "baaca", k=3 -> "aaabc" (sort)
	fmt.Println("Test 2:", orderlyQueue("baaca", 3)) // "aaabc"

	// Edge: k > len(s), k=5 -> sort
	fmt.Println("Test 3:", orderlyQueue("zxy", 5)) // "xyz"

	// Single character
	fmt.Println("Test 4:", orderlyQueue("a", 1)) // "a"

	// Already smallest rotation
	fmt.Println("Test 5:", orderlyQueue("abc", 1)) // "abc"
}
```
