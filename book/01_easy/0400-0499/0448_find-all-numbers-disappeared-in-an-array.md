# 0448 — Find All Numbers Disappeared In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindAllNumbersDisappearedInAnArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #448: Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindAllNumbersDisappearedInAnArray(nums []int) []int {
	for _, v := range nums {
		idx := v
		if idx < 0 {
			idx = -idx
		}
		idx--
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	var result []int
	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(FindAllNumbersDisappearedInAnArray([]int{1, 1}))
}
```
