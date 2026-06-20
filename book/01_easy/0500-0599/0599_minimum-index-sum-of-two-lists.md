# 0599 — Minimum Index Sum Of Two Lists

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumIndexSumOfTwoLists(list1, list2 []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #599: Minimum Index Sum of Two Lists
// https://leetcode.com/problems/minimum-index-sum-of-two-lists/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// Time: O(n+m), Space: O(n)
func MinimumIndexSumOfTwoLists(list1, list2 []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	index := make(map[string]int)
	for i, s := range list1 {
		index[s] = i
	}
	minSum := math.MaxInt32
	var result []string
	for j, s := range list2 {
		if i, ok := index[s]; ok {
			sum := i + j
			if sum < minSum {
				minSum = sum
				result = []string{s}
			} else if sum == minSum {
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	))
	fmt.Println(MinimumIndexSumOfTwoLists(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}
```
