# 2121 — Intervals Between Identical Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func getDistances(arr []int) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2121: Intervals Between Identical Elements
// https://leetcode.com/problems/intervals-between-identical-elements/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func getDistances(arr []int) []int64 {
	n := len(arr)
	// Group indices by value
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[int][]int)
	for i, v := range arr {
		groups[v] = append(groups[v], i)
	}

  // Alokasi slice integer
	result := make([]int64, n)
	for _, indices := range groups {
		m := len(indices)
		if m <= 1 {
			continue
		}
		// Prefix sum of distances
  // Alokasi slice integer
		prefix := make([]int64, m+1)
		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(indices[i])
		}
		for i, pos := range indices {
			// Sum of distances to all other same-value elements
			// Left side: pos * i - prefix[i]
			// Right side: (prefix[m] - prefix[i+1]) - pos * (m-1-i)
			left := int64(pos)*int64(i) - prefix[i]
			right := (prefix[m] - prefix[i+1]) - int64(pos)*int64(m-1-i)
			result[pos] = left + right
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getDistances([]int{2, 1, 3, 1, 2, 3, 3}))
	// Expected: [4, 2, 7, 2, 4, 4, 5]

	// Test case 2
	fmt.Println("Test 2:", getDistances([]int{10, 5, 10, 10}))
	// Expected: [5, 0, 3, 4]

	// Test case 3
	fmt.Println("Test 3:", getDistances([]int{1, 2, 3}))
	// Expected: [0, 0, 0]
}
```
