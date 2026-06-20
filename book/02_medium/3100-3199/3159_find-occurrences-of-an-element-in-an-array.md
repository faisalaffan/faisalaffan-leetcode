# 3159 — Find Occurrences Of An Element In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func occurrencesOfElement(nums []int, queries []int, x int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3159: Find Occurrences of an Element in an Array
// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
  // Alokasi slice integer
	pos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q-1 < len(pos) {
			ans[i] = pos[q-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(occurrencesOfElement([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1)) // Expected: [0, -1, 2, -1]
	fmt.Println(occurrencesOfElement([]int{1, 2, 3}, []int{10}, 5))             // Expected: [-1]
}
```
