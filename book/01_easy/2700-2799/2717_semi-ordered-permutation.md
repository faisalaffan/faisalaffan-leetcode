# 2717 — Semi Ordered Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SemiOrderedPermutation(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2717: Semi-Ordered Permutation
// https://leetcode.com/problems/semi-ordered-permutation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SemiOrderedPermutation([]int{2, 1, 4, 3}))
	fmt.Println(SemiOrderedPermutation([]int{2, 4, 1, 3}))
}

func SemiOrderedPermutation(nums []int) int {
	n := len(nums)
	pos1, posN := 0, 0
	for i, v := range nums {
		if v == 1 {
			pos1 = i
		}
		if v == n {
			posN = i
		}
	}

	swaps := pos1 + (n - 1 - posN)
	if pos1 > posN {
		swaps--
	}
	return swaps
}
```
