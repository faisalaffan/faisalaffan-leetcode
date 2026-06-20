# 2475 — Number Of Unequal Triplets In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfUnequalTripletsInArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2475: Number of Unequal Triplets in Array
// https://leetcode.com/problems/number-of-unequal-triplets-in-array/
// Difficulty: Easy
// Time O(n^3) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfUnequalTripletsInArray([]int{4, 4, 2, 4, 3})) // 3
	fmt.Println(NumberOfUnequalTripletsInArray([]int{1, 1, 1, 1, 1}))  // 0
}

func NumberOfUnequalTripletsInArray(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
```
