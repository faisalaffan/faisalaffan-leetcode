# 1310 — Xor Queries Of A Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func xorQueries(arr []int, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n + m) where n = len(arr), m = len(queries)  
**Kompleksitas Ruang:** O(n) for prefix XOR array

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1310: XOR Queries of a Subarray
// https://leetcode.com/problems/xor-queries-of-a-subarray/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	// [2,7,14,8]

	// Test case 2
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
	// [8,0,4,4]

	// Test case 3
	fmt.Println(xorQueries([]int{2}, [][]int{{0, 0}}))
	// [2]
}

// Time: O(n + m) where n = len(arr), m = len(queries)
// Space: O(n) for prefix XOR array
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

  // Alokasi slice integer
	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = prefix[q[1]+1] ^ prefix[q[0]]
	}
	return result
}
```
