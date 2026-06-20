# 1726 — Tuple With Same Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func tupleSameProduct(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2), Space: O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1726: Tuple with Same Product
// https://leetcode.com/problems/tuple-with-same-product/
// Difficulty: Medium
// Time: O(n^2), Space: O(n^2)

import "fmt"

func tupleSameProduct(nums []int) int {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	productCount := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product := nums[i] * nums[j]
			productCount[product]++
		}
	}

	result := 0
	for _, count := range productCount {
		if count > 1 {
			// Each pair of pairs = 8 tuples (4! / 3 = 8)
			result += count * (count - 1) / 2 * 8
		}
	}
	return result
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))     // Expected: 8
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10})) // Expected: 16
	fmt.Println(tupleSameProduct([]int{1, 2, 3, 4, 6, 12})) // Expected: 40
}
```
