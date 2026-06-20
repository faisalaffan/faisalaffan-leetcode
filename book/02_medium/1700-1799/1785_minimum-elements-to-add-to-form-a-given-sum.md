# 1785 — Minimum Elements To Add To Form A Given Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minElements(nums []int, limit int, goal int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1785: Minimum Elements to Add to Form a Given Sum
// https://leetcode.com/problems/minimum-elements-to-add-to-form-a-given-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	diff := goal - sum
	if diff < 0 {
		diff = -diff
	}

	// Minimum elements = ceil(diff / limit)
	return (diff + limit - 1) / limit
}

func main() {
	fmt.Println(minElements([]int{1, -1, 1}, 3, -4)) // Expected: 2
	fmt.Println(minElements([]int{1, -10, 9, 1}, 100, 0)) // Expected: 1
	fmt.Println(minElements([]int{0}, 1, 1000000)) // Expected: 1000000
}
```
