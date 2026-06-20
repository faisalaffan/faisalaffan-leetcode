# 3447 — Assign Elements To Groups With Constraints

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func assignElements(groups []int, elements []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(mx log mx + n + m) Space: O(mx)  
**Kompleksitas Ruang:** O(mx)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3447: Assign Elements to Groups with Constraints
// https://leetcode.com/problems/assign-elements-to-groups-with-constraints/
// Difficulty: Medium
// Time: O(mx log mx + n + m) Space: O(mx)

import (
	"fmt"
	"slices"
)

func assignElements(groups []int, elements []int) []int {
	mx := slices.Max(groups)
  // Alokasi slice integer
	target := make([]int, mx+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range target {
		target[i] = -1
	}

	for i, x := range elements {
		if x > mx || target[x] >= 0 {
			continue
		}
		for y := x; y <= mx; y += x {
			if target[y] < 0 {
				target[y] = i
			}
		}
	}

	for i, x := range groups {
		groups[i] = target[x]
	}
	return groups
}

func main() {
	fmt.Println(assignElements([]int{8, 4, 3, 2, 4}, []int{4, 2})) // [0 0 -1 1 0]
	fmt.Println(assignElements([]int{10, 5, 7}, []int{2, 5})) // [1 1 -1]
}
```
