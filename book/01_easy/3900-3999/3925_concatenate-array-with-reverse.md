# 3925 — Concatenate Array With Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConcatenateArrayWithReverse(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3925: Concatenate Array With Reverse
// https://leetcode.com/problems/concatenate-array-with-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenateArrayWithReverse([]int{1, 2, 3}))
	fmt.Println(ConcatenateArrayWithReverse([]int{1}))
}

// Time: O(n)
// Space: O(n)
func ConcatenateArrayWithReverse(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = nums[n-1-i]
	}
	return ans
}
```
