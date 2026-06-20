# 2460 — Apply Operations To An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ApplyOperationsToAnArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2460: Apply Operations to an Array
// https://leetcode.com/problems/apply-operations-to-an-array/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(ApplyOperationsToAnArray([]int{1, 2, 2, 1, 1, 0})) // [1,4,2,0,0,0]
	fmt.Println(ApplyOperationsToAnArray([]int{0, 1}))              // [1,0]
}

func ApplyOperationsToAnArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

  // Alokasi slice integer
	res := make([]int, n)
	idx := 0
	for _, v := range nums {
		if v != 0 {
			res[idx] = v
			idx++
		}
	}
	return res
}
```
