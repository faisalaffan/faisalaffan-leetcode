# 2855 — Minimum Right Shifts To Sort The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumRightShiftsToSortTheArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2855: Minimum Right Shifts to Sort the Array
// https://leetcode.com/problems/minimum-right-shifts-to-sort-the-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{1, 3, 5}))
}

func MinimumRightShiftsToSortTheArray(nums []int) int {
	n := len(nums)
	descentIdx := -1
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			if descentIdx != -1 {
				return -1
			}
			descentIdx = i
		}
	}
	if descentIdx == -1 {
		return 0
	}
	if nums[n-1] > nums[0] {
		return -1
	}
	return n - descentIdx - 1
}
```
