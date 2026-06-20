# 2898 — Maximum Linear Stock Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumLinearStockScore(prices []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2898: Maximum Linear Stock Score
// https://leetcode.com/problems/maximum-linear-stock-score/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLinearStockScore(prices []int) int64 {
	// For each stock, score = sum of prices where prices[i] - i is same
  // Membuat map (HashMap) — pencarian O(1)
	score := make(map[int]int64)
	var best int64

	for i, p := range prices {
		key := p - i
		score[key] += int64(p)
		if score[key] > best {
			best = score[key]
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLinearStockScore([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLinearStockScore([]int{2, 1, 3}))
}
```
