# 3526 — Range Xor Queries With Subarray Reversals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rangeXorQueries(arr []int, queries [][]int, reversals [][]int) []int
```

> **💡 Hint:** Use a Fenwick tree for XOR with a Treap for reversals,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum, Fenwick Tree (BIT)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3526: Range XOR Queries with Subarray Reversals
// https://leetcode.com/problems/range-xor-queries-with-subarray-reversals/
// Difficulty: Hard [Paid]
//
// Given an array, support range XOR queries and subarray reversals.
// Process both operations efficiently.
//
// Approach: Use a Fenwick tree for XOR with a Treap for reversals,
// or use sqrt decomposition for simplicity.

import "fmt"

func main() {
	// Example
	fmt.Println(rangeXorQueries([]int{1, 2, 3, 4}, [][]int{{0, 2}, {1, 3}}, [][]int{{1, 2}}))
	// Edge: no reversals
	fmt.Println(rangeXorQueries([]int{5, 6}, [][]int{{0, 1}}, [][]int{}))
	// Edge: single element
	fmt.Println(rangeXorQueries([]int{10}, [][]int{{0, 0}}, [][]int{}))
}

func rangeXorQueries(arr []int, queries [][]int, reversals [][]int) []int {
	// Copy the array since we need to handle reversals
  // Alokasi slice integer
	a := make([]int, len(arr))
	copy(a, arr)

	// Process reversals
	for _, rev := range reversals {
		l, r := rev[0], rev[1]
		for i, j := l, r; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	}

	// Prefix XOR
  // Alokasi slice integer
	pref := make([]int, len(a)+1)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(a); i++ {
		pref[i+1] = pref[i] ^ a[i]
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		ans[qi] = pref[r+1] ^ pref[l]
	}
	return ans
}
```
