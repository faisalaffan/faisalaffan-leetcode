# 3958 — Minimum Cost To Split Into Ones Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostToSplitIntoOnesIi(n int) int64
```

> **💡 Hint:** Optimal strategy splits off one 1 at a time.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3958: Minimum Cost to Split into Ones II
// https://leetcode.com/problems/minimum-cost-to-split-into-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)
// Approach: Optimal strategy splits off one 1 at a time.
// Total cost = 1 + 2 + ... + (n-1) = n*(n-1)/2.

import "fmt"

func MinimumCostToSplitIntoOnesIi(n int) int64 {
	return int64(n) * int64(n-1) / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToSplitIntoOnesIi(3)) // Expected: 3

	// Example 2
	fmt.Println(MinimumCostToSplitIntoOnesIi(4)) // Expected: 6

	// Example 3
	fmt.Println(MinimumCostToSplitIntoOnesIi(1)) // Expected: 0
}
```
