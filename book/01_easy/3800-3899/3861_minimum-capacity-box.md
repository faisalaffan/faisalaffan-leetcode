# 3861 — Minimum Capacity Box

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCapacityBox(capacity []int, itemSize int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3861: Minimum Capacity Box
// https://leetcode.com/problems/minimum-capacity-box/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumCapacityBox([]int{1, 5, 3, 7}, 3))
	fmt.Println(MinimumCapacityBox([]int{3, 5, 4, 3}, 2))
	fmt.Println(MinimumCapacityBox([]int{4}, 5))
}

// Time: O(n)
// Space: O(1)
func MinimumCapacityBox(capacity []int, itemSize int) int {
	ans := -1
	for i, c := range capacity {
		if c >= itemSize && (ans == -1 || c < capacity[ans]) {
			ans = i
		}
	}
	return ans
}
```
