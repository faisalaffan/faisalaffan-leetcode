# 1899 — Merge Triplets To Form Target Triplet

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeTriplets(triplets [][]int, target []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1899: Merge Triplets to Form Target Triplet
// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}))
	fmt.Println(MergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}))
	fmt.Println(MergeTriplets([][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}))
}

// Time: O(n), Space: O(1)
func MergeTriplets(triplets [][]int, target []int) bool {
	found := [3]bool{}
	for _, t := range triplets {
		// Skip any triplet that exceeds target
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}
		if t[0] == target[0] {
			found[0] = true
		}
		if t[1] == target[1] {
			found[1] = true
		}
		if t[2] == target[2] {
			found[2] = true
		}
	}
	return found[0] && found[1] && found[2]
}
```
