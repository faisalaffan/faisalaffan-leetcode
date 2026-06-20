# 1054 — Distant Barcodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func rearrangeBarcodes(barcodes []int) []int
```

> **💡 Hint:** Count frequencies, place most frequent in even indices, then odd

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1054: Distant Barcodes
// https://leetcode.com/problems/distant-barcodes/
// Difficulty: Medium
//
// Approach: Count frequencies, place most frequent in even indices, then odd
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 2, 2, 2})) // [1,2,1,2,1,2] or similar
	fmt.Println(rearrangeBarcodes([]int{1, 1, 1, 1, 2, 2, 3, 3})) // valid rearrangement
}

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, b := range barcodes {
		freq[b]++
	}

	type pair struct {
		val   int
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for val, count := range freq {
		pairs = append(pairs, pair{val, count})
	}
  // Custom sort dengan comparator
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

  // Alokasi slice integer
	result := make([]int, n)
	idx := 0

	for _, p := range pairs {
		for k := 0; k < p.count; k++ {
			result[idx] = p.val
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return result
}
```
