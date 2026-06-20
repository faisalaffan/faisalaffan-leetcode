# 2891 — Method Chaining

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MethodChaining(animals [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2891: Method Chaining
// https://leetcode.com/problems/method-chaining/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we filter animals with weight > 100, sort by weight,
// and rename the weight column.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: methodChaining
	// Input: [name, species, age, weight]
	animals := [][]int{{1, 1, 5, 50}, {2, 1, 3, 120}, {3, 2, 4, 150}, {4, 1, 2, 80}}
	fmt.Println(MethodChaining(animals))
	// [[2 1 3 120] [3 2 4 150]]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: methodChaining
func MethodChaining(animals [][]int) [][]int {
	// Filter: weight > 100 (column index 3)
	filtered := [][]int{}
	for _, a := range animals {
		if a[3] > 100 {
			filtered = append(filtered, a)
		}
	}
	// Sort by weight ascending
  // Custom sort dengan comparator
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][3] < filtered[j][3]
	})
	return filtered
}
```
