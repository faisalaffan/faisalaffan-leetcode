# 1196 — How Many Apples Can You Put Into The Basket

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxNumberOfApples(weight []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1196: How Many Apples Can You Put into the Basket
// https://leetcode.com/problems/how-many-apples-can-you-put-into-the-basket/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumberOfApples([]int{100, 200, 150, 1000}))          // 4
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800})) // 5
}

// LeetCode submission: maxNumberOfApples
func maxNumberOfApples(weight []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(weight)
	sum := 0
	for i, w := range weight {
		sum += w
		if sum > 5000 {
			return i
		}
	}
	return len(weight)
}
```
