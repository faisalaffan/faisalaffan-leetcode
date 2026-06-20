# 1920 — Build Array From Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BuildArrayFromPermutation(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1920: Build Array from Permutation
// https://leetcode.com/problems/build-array-from-permutation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(BuildArrayFromPermutation([]int{0, 2, 1, 5, 3, 4}))    // [0,1,2,4,5,3]
	fmt.Println(BuildArrayFromPermutation([]int{5, 0, 1, 2, 3, 4}))    // [4,5,0,1,2,3]
}

// Time: O(n), Space: O(n)
func BuildArrayFromPermutation(nums []int) []int {
  // Alokasi slice integer
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = nums[v]
	}
	return ans
}
```
