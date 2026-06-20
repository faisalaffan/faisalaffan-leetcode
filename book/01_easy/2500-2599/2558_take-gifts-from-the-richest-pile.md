# 2558 — Take Gifts From The Richest Pile

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TakeGiftsFromTheRichestPile(gifts []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2558: Take Gifts From the Richest Pile
// https://leetcode.com/problems/take-gifts-from-the-richest-pile/
// Difficulty: Easy
// Time O(k * n) | Space O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TakeGiftsFromTheRichestPile([]int{25, 64, 9, 4, 100}, 4)) // 29
	fmt.Println(TakeGiftsFromTheRichestPile([]int{1, 1, 1, 1}, 4))         // 4
}

func TakeGiftsFromTheRichestPile(gifts []int, k int) int64 {
	for t := 0; t < k; t++ {
		maxIdx := 0
		for i := 1; i < len(gifts); i++ {
			if gifts[i] > gifts[maxIdx] {
				maxIdx = i
			}
		}
		gifts[maxIdx] = int(math.Sqrt(float64(gifts[maxIdx])))
	}
	sum := int64(0)
	for _, g := range gifts {
		sum += int64(g)
	}
	return sum
}
```
