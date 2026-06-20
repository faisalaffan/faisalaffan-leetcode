# 3857 — Minimum Cost To Split Into Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostToSplitIntoOnes(n int) int
```

> **💡 Hint:** Minimum cost = n*(n-1)/2. Equivalent to total edges in complete graph.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3857: Minimum Cost to Split into Ones
// https://leetcode.com/problems/minimum-cost-to-split-into-ones/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Minimum cost = n*(n-1)/2. Equivalent to total edges in complete graph.

import "fmt"

func MinimumCostToSplitIntoOnes(n int) int {
	return n * (n - 1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnes(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnes(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnes(10))
}
```
