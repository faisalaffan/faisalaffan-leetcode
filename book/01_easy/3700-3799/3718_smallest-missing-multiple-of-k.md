# 3718 — Smallest Missing Multiple Of K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SmallestMissingMultipleOfK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + max_missing/k)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3718: Smallest Missing Multiple of K
// https://leetcode.com/problems/smallest-missing-multiple-of-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SmallestMissingMultipleOfK([]int{8, 2, 3, 4, 6}, 2))
	fmt.Println(SmallestMissingMultipleOfK([]int{1, 4, 7, 10, 15}, 5))
}

// Time: O(n + max_missing/k)
// Space: O(n)
func SmallestMissingMultipleOfK(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	has := make(map[int]bool)
	for _, v := range nums {
		has[v] = true
	}

	for x := k; ; x += k {
		if !has[x] {
			return x
		}
	}
}
```
