# 0646 — Maximum Length Of Pair Chain

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLongestChain(pairs [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #646: Maximum Length of Pair Chain
// https://leetcode.com/problems/maximum-length-of-pair-chain/
// Difficulty: Medium
// Time: O(n log n) for sorting
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}))
	fmt.Println(findLongestChain([][]int{{1, 2}, {7, 8}, {4, 5}}))
}

func findLongestChain(pairs [][]int) int {
  // Custom sort dengan comparator
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	count := 0
	curEnd := -1 << 31

	for _, pair := range pairs {
		if pair[0] > curEnd {
			curEnd = pair[1]
			count++
		}
	}

	return count
}
```
