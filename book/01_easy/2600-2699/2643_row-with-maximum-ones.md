# 2643 — Row With Maximum Ones

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RowWithMaximumOnes(mat [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2643: Row With Maximum Ones
// https://leetcode.com/problems/row-with-maximum-ones/
// Difficulty: Easy
// Time: O(m * n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(RowWithMaximumOnes([][]int{{0, 1}, {1, 0}}))
	fmt.Println(RowWithMaximumOnes([][]int{{0, 0, 0}, {0, 1, 1}}))
}

func RowWithMaximumOnes(mat [][]int) []int {
	maxRow, maxCount := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(mat); i++ {
		count := 0
		for _, val := range mat[i] {
			if val == 1 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxRow = i
		}
	}
	return []int{maxRow, maxCount}
}
```
