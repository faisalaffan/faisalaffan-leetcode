# 3736 — Minimum Moves To Equal Array Elements Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumMovesToEqualArrayElementsIii(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3736: Minimum Moves to Equal Array Elements III
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{2, 1, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{4, 4, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumMovesToEqualArrayElementsIii(nums []int) int {
	maxVal := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal*len(nums) - sum
}
```
