# 0775 — Global And Local Inversions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isIdealPermutation(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #775: Global and Local Inversions
// https://leetcode.com/problems/global-and-local-inversions/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(isIdealPermutation([]int{1, 0, 2}))
	fmt.Println(isIdealPermutation([]int{1, 2, 0}))
}

func isIdealPermutation(nums []int) bool {
	maxVal := -1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if maxVal > nums[i+2] {
			return false
		}
	}
	return true
}
```
