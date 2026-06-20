# 1198 — Find Smallest Common Element In All Rows

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestCommonElement(mat [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(max value) = O(10000) since values are 1..10000

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1198: Find Smallest Common Element in All Rows
// https://leetcode.com/problems/find-smallest-common-element-in-all-rows/
// Difficulty: Medium [Paid]

// Find smallest integer that appears in every row.

// Time: O(m * n)
// Space: O(max value) = O(10000) since values are 1..10000

func smallestCommonElement(mat [][]int) int {
	if len(mat) == 0 {
		return -1
	}

  // Alokasi slice integer
	count := make([]int, 10001)
	for _, v := range mat[0] {
		count[v] = 1
	}

	for i := 1; i < len(mat); i++ {
		for _, v := range mat[i] {
			if count[v] == i {
				count[v]++
			}
		}
	}

	for v := 1; v <= 10000; v++ {
		if count[v] == len(mat) {
			return v
		}
	}
	return -1
}

func main() {
	fmt.Printf("%d (expected: 5)\n",
		smallestCommonElement([][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}))

	fmt.Printf("%d (expected: -1)\n",
		smallestCommonElement([][]int{{1, 2, 3}, {4, 5, 6}}))

	fmt.Printf("%d (expected: 2)\n",
		smallestCommonElement([][]int{{2, 3}, {2, 5}}))
}
```
