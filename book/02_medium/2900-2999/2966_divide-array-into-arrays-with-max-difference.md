# 2966 — Divide Array Into Arrays With Max Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func divideArray2966(nums []int, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2966: Divide Array Into Arrays With Max Difference
// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divideArray2966([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray2966([]int{1, 3, 3, 2, 7, 3}, 3))
	fmt.Println(divideArray2966([]int{1, 2, 3}, 0))
}

func divideArray2966(nums []int, k int) [][]int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	ans := [][]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 3 {
  // Alokasi slice integer
		t := make([]int, 3)
		copy(t, nums[i:i+3])
		if t[2]-t[0] > k {
			return [][]int{}
		}
		ans = append(ans, t)
	}
	return ans
}
```
