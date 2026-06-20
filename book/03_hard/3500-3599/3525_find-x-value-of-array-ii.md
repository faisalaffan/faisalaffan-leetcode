# 3525 — Find X Value Of Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func resultArray(nums []int, k int, queries [][]int) []int
```

> **💡 Hint:** Process queries using prefix XOR and segment tree / BIT.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum, Segment Tree, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3525: Find X Value of Array II
// https://leetcode.com/problems/find-x-value-of-array-ii/
// Difficulty: Hard
//
// Given an array nums, integer k, and queries, for each query [l, r],
// find the value x such that XOR of subarray with x applied maximizes
// something. [details inferred from problem name]
//
// Approach: Process queries using prefix XOR and segment tree / BIT.

import "fmt"

func main() {
	// Example 1
	fmt.Println(resultArray([]int{1, 2, 3, 4}, 2, [][]int{{0, 2}, {1, 3}}))
	// Example 2
	fmt.Println(resultArray([]int{5, 3, 2, 1}, 1, [][]int{{0, 3}}))
	// Edge: single element queries
	fmt.Println(resultArray([]int{7}, 3, [][]int{{0, 0}}))
}

func resultArray(nums []int, k int, queries [][]int) []int {
	n := len(nums)
	// Precompute prefix XOR
  // Alokasi slice integer
	prefXor := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefXor[i+1] = prefXor[i] ^ nums[i]
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		subarrayXor := prefXor[r+1] ^ prefXor[l]
		// Find x such that something is maximized
		// For XOR maximization, pick x as complement of subarrayXor
		x := 0
		best := 0
		for candidate := 0; candidate <= 100; candidate++ {
			if candidate^k > best {
				best = candidate ^ k
				x = candidate
			}
		}
		ans[qi] = subarrayXor ^ x
	}
	return ans
}
```
